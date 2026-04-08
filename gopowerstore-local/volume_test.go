/*
 *
 * Copyright © 2020-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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

package gopowerstore

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	volumeMockURL    = APIMockURL + volumeURL
	applianceMockURL = APIMockURL + applianceURL
)

var (
	volID       = "6b930711-46bc-4a4b-9d6a-22c77a7838c4"
	volID2      = "3765da74-28a7-49db-a693-10cec1de91f8"
	appID       = "A1"
	volSnapID   = "1966782b-60c9-40e2-a1ee-9b2b8f6b98e7"
	volSnapID2  = "34380c29-2203-4490-aeb7-2853b9a85075"
	metroConfig = MetroConfig{
		RemoteSystemID:    "47921973-b0eb-485d-8492-c5d7f6ca216c",
		RemoteApplianceID: appID,
	}
)

type VolumeTestSuite struct {
	suite.Suite
}

func TestVolumeSuite(t *testing.T) {
	suite.Run(t, new(VolumeTestSuite))
}

func (s *VolumeTestSuite) SetupTest() {
	httpmock.Activate()
}

func (s *VolumeTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (s *VolumeTestSuite) TestClientIMPL_GetVolumes() {
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, volID, volID2)
	httpmock.RegisterResponder("GET", volumeMockURL,
		httpmock.NewStringResponder(200, respData))
	vols, err := C.GetVolumes(context.Background())
	assert.Nil(s.T(), err)
	assert.Len(s.T(), vols, 2)
	assert.Equal(s.T(), volID, vols[0].ID)
}

func (s *VolumeTestSuite) TestClientIMPL_GetVolume() {
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(200, respData))
	vol, err := C.GetVolume(context.Background(), volID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), volID, vol.ID)
}

func (s *VolumeTestSuite) TestClientIMPL_GetVolumeByName() {
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", volumeMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	vol, err := C.GetVolumeByName(context.Background(), "test")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), volID, vol.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetVolumeByName(context.Background(), "test")
	assert.NotNil(s.T(), err)
	apiError := err.(APIError)
	assert.True(s.T(), apiError.NotFound())
}

func (s *VolumeTestSuite) TestClientIMPL_GetAppliance() {
	respData := fmt.Sprintf(`{"id": "%s"}`, appID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", applianceMockURL, appID),
		httpmock.NewStringResponder(200, respData))
	app, err := C.GetAppliance(context.Background(), appID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), appID, app.ID)
}

func (s *VolumeTestSuite) TestClientIMPL_GetApplianceByName() {
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", applianceMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, appID)
	setResponder(respData)
	ap, err := C.GetApplianceByName(context.Background(), "test")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), appID, ap.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetApplianceByName(context.Background(), "test")
	assert.NotNil(s.T(), err)
	apiError := err.(APIError)
	assert.True(s.T(), apiError.NotFound())
}

func (s *VolumeTestSuite) TestClientIMPL_GetSnapshotsByVolumeID() {
	respData := fmt.Sprintf(`[{
		"description":"",
		"id":"%s",
		"name":"rpo.VOLUME_esa51_volume_test1.2019-12-06T12:35:21Z 183173616",
		"size":10737418240,
		"state":"Ready",		
		"type":"Snapshot",
		"wwn":null,
		"protection_data":{
			"family_id": "52ebb13c-16a0-4466-9319-a182b58b1c39",
			"parent_id": "52ebb13c-16a0-4466-9319-a182b58b1c39",
			"source_id": "%s", 
			"creator_type": "System", 
			"copy_signature": "e79bbc7f-da33-4f24-8eaa-0e3ef7533153",
			"source_timestamp": "2019-12-06T12:35:21.873907+00:00",
			"creator_type_l10n": "System",
			"is_app_consistent": null,
			"created_by_rule_id": null,
			"created_by_rule_name": null,
			"expiration_timestamp": null
		}
	}]`, volID2, volID)
	httpmock.RegisterResponder("GET", volumeMockURL,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetSnapshotsByVolumeID(context.Background(), volID2)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(resp))
	assert.Equal(s.T(), volID2, resp[0].ID)
}

func (s *VolumeTestSuite) TestClientIMPL_GetSnapshot() {
	respData := fmt.Sprintf(`{"id": "%s"}`, volSnapID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", volumeMockURL, volSnapID),
		httpmock.NewStringResponder(200, respData))
	snapshot, err := C.GetSnapshot(context.Background(), volSnapID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), volSnapID, snapshot.ID)
}

func (s *VolumeTestSuite) TestClientIMPL_GetSnapshots() {
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, volSnapID, volSnapID2)
	httpmock.RegisterResponder("GET", volumeMockURL,
		httpmock.NewStringResponder(200, respData))
	snapshots, err := C.GetSnapshots(context.Background())
	assert.Nil(s.T(), err)
	assert.Len(s.T(), snapshots, 2)
	assert.Equal(s.T(), volSnapID, snapshots[0].ID)
}

func (s *VolumeTestSuite) TestClientIMPL_GetSnapshotByName() {
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", volumeMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volSnapID)
	setResponder(respData)
	snap, err := C.GetSnapshotByName(context.Background(), "test")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), volSnapID, snap.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetSnapshotByName(context.Background(), "test")
	assert.NotNil(s.T(), err)
	apiError := err.(APIError)
	assert.True(s.T(), apiError.NotFound())
}

func (s *VolumeTestSuite) TestClientIMPL_CreateVolume() {
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", volumeMockURL,
		httpmock.NewStringResponder(201, respData))
	name := "test_vol"
	size := int64(11111111)
	createReq := VolumeCreate{}
	createReq.Name = &name
	createReq.Size = &size

	resp, err := C.CreateVolume(context.Background(), &createReq)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), volID, resp.ID)
}

func (s *VolumeTestSuite) TestClientIMPL_CreateSnapshot() {
	respData := fmt.Sprintf(`{"id": "%s"}`, volID2)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/snapshot", volumeMockURL, volID),
		httpmock.NewStringResponder(201, respData))
	name := "test_vol"
	desc := "desc"
	createReq := SnapshotCreate{}
	createReq.Name = &name
	createReq.Description = &desc

	resp, err := C.CreateSnapshot(context.Background(), &createReq, volID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), volID2, resp.ID)
}

func (s *VolumeTestSuite) TestClientIMPL_CreateVolumeFromSnapshot() {
	respData := fmt.Sprintf(`{"id": "%s"}`, volID2)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/clone", volumeMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	name := "new_volume_from_snap"
	createParams := VolumeClone{}
	createParams.Name = &name
	resp, err := C.CreateVolumeFromSnapshot(context.Background(), &createParams, volID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), volID2, resp.ID)
}

func (s *VolumeTestSuite) TestClientIMPL_ComputeDifferences() {
	respData := `{"chunk_bitmap":"Dw==","next_offset":-1}`
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/compute_differences", volumeMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	baseSnapshotID := ""
	offset := int64(0)
	chunkSize := int64(1048576)
	length := int64(4194304)
	computeDiffParams := VolumeComputeDifferences{
		BaseSnapshotID: &baseSnapshotID,
		ChunkSize:      &chunkSize,
		Length:         &length,
		Offset:         &offset,
	}
	resp, err := C.ComputeDifferences(context.Background(), &computeDiffParams, volID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "Dw==", *resp.ChunkBitmap)
	assert.Equal(s.T(), int64(-1), *resp.NextOffset)
}

func (s *VolumeTestSuite) TestClientIMPL_ModifyVolume() {
	respData := ""
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(201, respData))

	modifyParams := VolumeModify{
		Name: "newname",
		Size: 8192 * 99,
	}

	resp, err := C.ModifyVolume(context.Background(), &modifyParams, volID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), EmptyResponse(""), resp)
}

func (s *VolumeTestSuite) TestClientIMPL_DeleteSnapshot() {
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(204, ""))
	force := true
	deleteReq := VolumeDelete{}
	deleteReq.ForceInternal = &force
	resp, err := C.DeleteSnapshot(context.Background(), &deleteReq, volID)
	assert.Nil(s.T(), err)
	assert.Len(s.T(), string(resp), 0)
}

func (s *VolumeTestSuite) TestClientIMPL_DeleteVolume() {
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", volumeMockURL, volID),
		httpmock.NewStringResponder(204, ""))
	force := true
	deleteReq := VolumeDelete{}
	deleteReq.ForceInternal = &force
	resp, err := C.DeleteVolume(context.Background(), &deleteReq, volID)
	assert.Nil(s.T(), err)
	assert.Len(s.T(), string(resp), 0)
}

func (s *VolumeTestSuite) TestClientIMPL_ConfigureMetroVolume() {
	sessionID := "test-id"
	sessionIDJSON := fmt.Sprintf(`{"metro_replication_session_id": "%s"}`, sessionID)

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/%s", volumeMockURL, volID, VolumeActionConfigureMetro),
		httpmock.NewStringResponder(http.StatusOK, sessionIDJSON))

	resp, err := C.ConfigureMetroVolume(context.Background(), volID, &metroConfig)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), sessionID, resp.ID)
}

func (s *VolumeTestSuite) TestClientIMPL_EndMetroVolume() {
	opts := EndMetroVolumeOptions{DeleteRemoteVolume: true}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/%s", volumeMockURL, volID, VolumeActionEndMetro),
		httpmock.NewStringResponder(http.StatusNoContent, ""))

	resp, err := C.EndMetroVolume(context.Background(), volID, &opts)
	assert.Nil(s.T(), err)
	assert.Empty(s.T(), resp)
}
