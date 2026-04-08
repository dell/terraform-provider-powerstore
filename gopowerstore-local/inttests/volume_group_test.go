/*
 *
 * Copyright © 2024 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package inttests

import (
	"context"
	"net/http"
	"strings"
	"testing"

	g "github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	VGPrefix string = "test_vg_"
)

type VolumeGroupTestSuite struct {
	suite.Suite

	client  g.Client
	request g.VolumeGroup

	// Assurance that a volume group exists for the tests.
	assurance g.CreateResponse
}

func TestVolumeGroupSuite(t *testing.T) {
	suite.Run(t, new(VolumeGroupTestSuite))
}

func (s *VolumeGroupTestSuite) SetupSuite() {
	s.client = GetNewClient()

	var err error
	// Make sure a volume group exists on which we can run tests against.
	s.assurance, err = s.client.CreateVolumeGroup(context.Background(), &g.VolumeGroupCreate{
		Name: VGPrefix + randString(8),
	})
	assert.NoError(s.T(), err)
}

func (s *VolumeGroupTestSuite) TearDownSuite() {
	// Delete the volume group we created for the test.
	_, err := s.client.DeleteVolumeGroup(context.Background(), s.assurance.ID)
	assert.NoError(s.T(), err)
}

func (s *VolumeGroupTestSuite) SetupTest() {
}

func (s *VolumeGroupTestSuite) TearDownTest() {
}

// Returns true if one of the volume groups in vgs has an ID matching
// the ID provided by id and false if none of the volume groups have
// a matching ID.
func containsVolumeGroupID(vgs []g.VolumeGroup, id string) bool {
	for _, vg := range vgs {
		if vg.ID == id {
			return true
		}
	}
	return false
}

// Happy path test.
func (s *VolumeGroupTestSuite) TestGetVolumeGroups() {
	resp, err := s.client.GetVolumeGroups(context.Background())

	if assert.NoError(s.T(), err) {
		assert.True(s.T(), containsVolumeGroupID(resp, s.assurance.ID))
	}
}

/*
/	////////////////////////////
/	/ METRO VOLUME GROUP TESTS /
/ 	////////////////////////////
*/

// Test struct for metro volume configure and end metro volume suites.
type MetroVolumeGroupTest struct {
	client g.Client

	vg struct {
		this      g.VolumeGroup
		volumeIDs []string
	}

	metro struct {
		config  g.MetroConfig
		endOpts g.EndMetroVolumeGroupOptions
	}
}

type MetroVolumeGroupTestSuite struct {
	suite.Suite

	MetroVolumeGroupTest
}

func TestMetroVolumeGroupSuite(t *testing.T) {
	suite.Run(t, new(MetroVolumeGroupTestSuite))
}

func (s *MetroVolumeGroupTestSuite) SetupSuite() {
	s.client = GetNewClient()

	// Get a remote system configured for metro replication.
	remoteSystem := GetRemoteSystemForMetro(s.client, s.T())
	if remoteSystem.ID == "" {
		s.T().Skip("Could not get a remote system configured for metro. Skipping test suite...")
	}

	s.metro.config = g.MetroConfig{RemoteSystemID: remoteSystem.ID}

	// Set end-metro configuration to delete remote VG.
	s.metro.endOpts = g.EndMetroVolumeGroupOptions{
		DeleteRemoteVolumeGroup: true,
	}
}

// Create a volume group with one or more volumes in the group for testing purposes.
// Save the volume IDs and the volume group ID for use in subsequent tests.
func (s *MetroVolumeGroupTestSuite) SetupTest() {
	// Create a volume to add to the vg to make it a valid vg we can test with.
	volID, _ := CreateVol(s.T())
	s.vg.volumeIDs = append(s.vg.volumeIDs, volID)

	// Create a unique vg name for each test run.
	s.vg.this.Name = VGPrefix + randString(8)

	isWriteOrderConsistent := true
	// Create a volume group to run tests against.
	resp, err := s.client.CreateVolumeGroup(context.Background(), &g.VolumeGroupCreate{
		Name:                   s.vg.this.Name,
		VolumeIDs:              s.vg.volumeIDs,
		IsWriteOrderConsistent: &isWriteOrderConsistent,
	})
	assert.NoError(s.T(), err)

	s.vg.this.ID = resp.ID
}

// End the metro session on the volume group, remove the volumes from the
// volume group and delete them, delete the volume group, and sanitize
// test variables for the next test run.
func (s *MetroVolumeGroupTestSuite) TearDownTest() {
	// End metro vg replication session created during testing.
	s.client.EndMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.endOpts)

	// Delete all the volumes in the volume group.
	err := deleteAllVolumesInVG(s.client, s.vg.this.ID, s.vg.volumeIDs)
	if err != nil {
		s.T().Logf("%s Please delete from PowerStore when tests complete.", err.Error())
	}

	// Delete the volume group from the previous test.
	_, err = s.client.DeleteVolumeGroup(context.Background(), s.vg.this.ID)
	if err != nil {
		// 404 status means it was already deleted.
		// Warn about other errors encountered while deleting.
		if err.(g.APIError).StatusCode != http.StatusNotFound {
			s.T().Logf("Unable to delete test volume group %s. Please delete from PowerStore when tests complete. err: %s", s.vg.this.Name, err.Error())
		}
	}

	// Sanitize for next test.
	s.vg.this.Name = ""
	s.vg.this.ID = ""
	s.vg.volumeIDs = []string{}
}

func deleteAllVolumesInVG(c g.Client, vgID string, volumeIDs []string) error {
	// Must remove volumes from the volume group before deleting.
	_, err := c.RemoveMembersFromVolumeGroup(context.Background(), &g.VolumeGroupMembers{VolumeIDs: volumeIDs}, vgID)
	if err != nil {
		if !strings.Contains(err.Error(), "One or more volumes to be removed are not part of the volume group") &&
			err.(g.APIError).StatusCode != http.StatusNotFound {
			return err
		}
	}

	for _, volID := range volumeIDs {
		_, err = c.DeleteVolume(context.Background(), nil, volID)
		if err != nil {
			// 404 status means it was already deleted.
			// Warn about other errors encountered while deleting.
			if err.(g.APIError).StatusCode != http.StatusNotFound {
				return err
			}
		}
	}
	return nil
}

// Should configure a metro volume group without errors.
func (s *MetroVolumeGroupTestSuite) TestConfigureMetroVolumeGroup() {
	resp, err := s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.config)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), resp)
}

// Make sure GetVolumeGroup returns the metro_replication_session_id
func (s *MetroVolumeGroupTestSuite) TestGetMetroVGSessionFromVG() {
	resp, err := s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.config)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), resp)

	// Get the metro_replication_session_id from the Volume Group
	vg, err := s.client.GetVolumeGroup(context.Background(), s.vg.this.ID)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), resp.ID, vg.MetroReplicationSessionID)
}

// Try to configure metro on a volume group without any volumes in it.
func (s *MetroVolumeGroupTestSuite) TestConfigMetroVGOnEmptyVG() {
	// Delete all the volumes from the volume group.
	err := deleteAllVolumesInVG(s.client, s.vg.this.ID, s.vg.volumeIDs)
	assert.NoError(s.T(), err)

	// Attempt to configure metro on an empty volume group.
	_, err = s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.config)

	if assert.Error(s.T(), err) {
		assert.Equal(s.T(), http.StatusUnprocessableEntity, err.(g.APIError).StatusCode)
		assert.Contains(s.T(), err.(g.APIError).Message, "Replication session creation failed as Volume Group")
	}
}

// Try to configure metro on a non-existent volume group.
func (s *MetroVolumeGroupTestSuite) TestMetroVGNonExistantVG() {
	// Delete all the volumes from the volume group.
	err := deleteAllVolumesInVG(s.client, s.vg.this.ID, s.vg.volumeIDs)
	assert.NoError(s.T(), err)

	// Delete that volume group, retaining the volume group ID.
	_, err = s.client.DeleteVolumeGroup(context.Background(), s.vg.this.ID)
	assert.NoError(s.T(), err)

	// Try to configure metro volume group using the deleted vg ID.
	_, err = s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.config)

	if assert.Error(s.T(), err) {
		assert.Equal(s.T(), http.StatusNotFound, err.(g.APIError).StatusCode)
		assert.Contains(s.T(), err.(g.APIError).Message, "Unable to find volume group")
	}
}

// Execute ConfigureMetroVolume with a bad request body.
func (s *MetroVolumeGroupTestSuite) TestMetroVGBadRequest() {
	// Pass an emtpy configuration body with the request.
	_, err := s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.this.ID, nil)

	if assert.Error(s.T(), err) {
		assert.Equal(s.T(), http.StatusBadRequest, err.(g.APIError).StatusCode)
	}
}

/*
/	////////////////////////////////
/	/ END METRO VOLUME GROUP TESTS /
/	////////////////////////////////
*/
type EndMetroVolumeGroupTestSuite struct {
	suite.Suite

	MetroVolumeGroupTest
}

func TestEndMetroVolumeGroupSuite(t *testing.T) {
	suite.Run(t, new(EndMetroVolumeGroupTestSuite))
}

func (s *EndMetroVolumeGroupTestSuite) SetupSuite() {
	// Get a new client.
	s.client = GetNewClient()

	// Get the remote PowerStore array.
	remoteSystem := GetRemoteSystemForMetro(s.client, s.T())
	if remoteSystem.ID == "" {
		s.T().Skip("Could not get a remote system configured for metro. Skipping test suite...")
	}

	s.metro.config = g.MetroConfig{RemoteSystemID: remoteSystem.ID}

	// Make sure remote VGs are always deleted.
	s.metro.endOpts = g.EndMetroVolumeGroupOptions{
		DeleteRemoteVolumeGroup: true,
	}
}

func (s *EndMetroVolumeGroupTestSuite) SetupTest() {
	// Create a volume to add to the vg to make it a valid vg we can test with.
	volID, _ := CreateVol(s.T())
	s.vg.volumeIDs = append(s.vg.volumeIDs, volID)

	// Create a unique vg name for each test run.
	s.vg.this.Name = VGPrefix + randString(8)

	isWriteOrderConsistent := true
	// Create a volume group to run tests against.
	resp, err := s.client.CreateVolumeGroup(context.Background(), &g.VolumeGroupCreate{
		Name:                   s.vg.this.Name,
		VolumeIDs:              s.vg.volumeIDs,
		IsWriteOrderConsistent: &isWriteOrderConsistent,
	})
	assert.NoError(s.T(), err)

	s.vg.this.ID = resp.ID

	// Create a metro vg session.
	_, err = s.client.ConfigureMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.config)
	if err != nil {
		s.T().Skipf("Could not create metro volume group session. Skipping test... Err: %s", err)
	}
}

func (s *EndMetroVolumeGroupTestSuite) TearDownTest() {
	// End the metro session if one still exists.
	s.client.EndMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.endOpts)

	// Delete all the volumes in the volume group.
	err := deleteAllVolumesInVG(s.client, s.vg.this.ID, s.vg.volumeIDs)
	if err != nil {
		s.T().Logf("%s Please delete from PowerStore when tests complete.", err.Error())
	}

	// Delete the volume group from the previous test.
	_, err = s.client.DeleteVolumeGroup(context.Background(), s.vg.this.ID)
	if err != nil {
		// 404 status means it was already deleted.
		// Warn about other errors encountered while deleting.
		if err.(g.APIError).StatusCode != http.StatusNotFound {
			s.T().Logf("Unable to delete test volume group %s. Please delete from PowerStore when tests complete. err: %s", s.vg.this.Name, err.Error())
		}
	}

	// Sanitize for next test.
	s.vg.this.Name = ""
	s.vg.this.ID = ""
	s.vg.volumeIDs = []string{}
}

// End a valid metro volume group session. Should end without error.
func (s *EndMetroVolumeGroupTestSuite) TestEndMetroVolumeGroup() {
	// End the metro session for the volume group.
	_, err := s.client.EndMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.endOpts)

	assert.NoError(s.T(), err)
}

// Try to end a metro volume group session with an invalid session ID.
func (s *EndMetroVolumeGroupTestSuite) TestEndMetroVGInvalidSessionID() {
	// Invalid session ID.
	invalidSessionID := "invalid-id"

	// End metro vg session with invalid vg ID.
	_, err := s.client.EndMetroVolumeGroup(context.Background(), invalidSessionID, &s.metro.endOpts)

	if assert.Error(s.T(), err) {
		assert.Equal(s.T(), http.StatusNotFound, err.(g.APIError).StatusCode)
		assert.Contains(s.T(), err.(g.APIError).Message, "Unable to find volume group")
	}
}

// Try to end a metro VG session for a VG that is not currently part of a metro session.
func (s *EndMetroVolumeGroupTestSuite) TestEndMetroOnUnreplicatedVG() {
	// Setup test scenario by ending the metro session.
	_, err := s.client.EndMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.endOpts)
	if err != nil {
		s.T().Skipf("Could not end metro volume group session. Skipping test... Err: %s", err)
	}

	// Try to end the metro vg session again.
	_, err = s.client.EndMetroVolumeGroup(context.Background(), s.vg.this.ID, &s.metro.endOpts)

	if assert.Error(s.T(), err) {
		assert.Equal(s.T(), http.StatusBadRequest, err.(g.APIError).StatusCode)
		assert.Contains(s.T(), err.(g.APIError).Message, "the volume group is not in a metro replication session")
	}
}
