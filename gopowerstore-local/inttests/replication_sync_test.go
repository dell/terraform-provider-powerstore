/*
 *
 * Copyright © 2024 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
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
	"os"
	"testing"

	"github.com/dell/gopowerstore"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ReplicationTestSuiteSync struct {
	suite.Suite
	remoteSystem    string
	remoteSystemMIP string
	randomString    string
	remoteClient    gopowerstore.Client
	pp              gopowerstore.CreateResponse
	vg              gopowerstore.CreateResponse
	rr              gopowerstore.CreateResponse
	vol             gopowerstore.CreateResponse
}

func (suite *ReplicationTestSuiteSync) SetupSuite() {
	// instead of completely hardcoded/constant string, let's make it dynamic
	// in case if TearDownSuite doesn't run at the end, we will not be blocked for next round of testing
	suite.randomString = randString(8)
	suite.remoteSystem, suite.remoteSystemMIP = getRemoteSystemSync(suite.T(), suite)
	err := godotenv.Load("GOPOWERSTORE_TEST.env")
	if err != nil {
		return
	}
	user := os.Getenv("GOPOWERSTORE_USERNAME")
	pass := os.Getenv("GOPOWERSTORE_PASSWORD")

	clientOptions := &gopowerstore.ClientOptions{}
	clientOptions.SetInsecure(true)
	client, err := gopowerstore.NewClientWithArgs("https://"+suite.remoteSystemMIP+"/api/rest", user, pass, clientOptions)
	if err != nil {
		return
	}
	suite.remoteClient = client
}

func (suite *ReplicationTestSuiteSync) TearDownSuite() {
	vg, err := C.GetVolumeGroup(context.Background(), suite.vg.ID)
	assert.NoError(suite.T(), err)
	pp, err := C.GetProtectionPolicyByName(context.Background(), "intcsi"+suite.randomString+"-pptst")
	assert.NoError(suite.T(), err)
	rr, err := C.GetReplicationRuleByName(context.Background(), "intcsi"+suite.randomString+"-ruletst")
	assert.NoError(suite.T(), err)
	if len(rr.ProtectionPolicies) != 1 || len(pp.ReplicationRules) != 1 || len(vg.Volumes) != 1 || len(pp.VolumeGroups) != 1 {
		suite.T().Fail()
	}
	C.ModifyVolumeGroup(context.Background(), &gopowerstore.VolumeGroupModify{ProtectionPolicyID: ""}, suite.vg.ID)
	C.RemoveMembersFromVolumeGroup(context.Background(), &gopowerstore.VolumeGroupMembers{VolumeIDs: []string{suite.vol.ID}}, suite.vg.ID)
	C.ModifyVolume(context.Background(), &gopowerstore.VolumeModify{ProtectionPolicyID: ""}, suite.vol.ID)
	C.DeleteProtectionPolicy(context.Background(), suite.pp.ID)
	C.DeleteReplicationRule(context.Background(), suite.rr.ID)
	C.DeleteVolumeGroup(context.Background(), suite.vg.ID)
	vgid, err := suite.remoteClient.GetVolumeGroupByName(context.Background(), "intcsi"+suite.randomString+"-vgtst")
	if err != nil {
		logrus.Info(err)
	}
	suite.remoteClient.DeleteVolumeGroup(context.Background(), vgid.ID)
	C.DeleteVolume(context.Background(), nil, suite.vol.ID)
}

func getRemoteSystemSync(t *testing.T, suite *ReplicationTestSuiteSync) (string, string) {
	resp, err := C.GetAllRemoteSystems(context.Background())
	skipTestOnError(t, err)
	if len(resp) == 0 {
		t.Skip("Skipping test as there are no remote systems configured on array.")
	}
	// try to find the working remote system from the list of all available/configured remoteSystems
	for i := range resp {
		rs, err := C.GetRemoteSystem(context.Background(), resp[i].ID)
		assert.NoError(t, err)
		assert.Equal(t, rs.ID, resp[i].ID)
		// create replicationRule and Protection policy beforeHand to check if remote system is working fine or not
		suite.rr, err = C.CreateReplicationRule(context.Background(), &gopowerstore.ReplicationRuleCreate{
			Name:           "intcsi" + suite.randomString + "-ruletst",
			Rpo:            gopowerstore.RpoZero,
			RemoteSystemID: rs.ID,
		})
		assert.NoError(t, err)

		suite.pp, err = C.CreateProtectionPolicy(context.Background(), &gopowerstore.ProtectionPolicyCreate{
			Name:               "intcsi" + suite.randomString + "-pptst",
			ReplicationRuleIDs: []string{suite.rr.ID},
		})

		if err == nil {
			return resp[i].ID, resp[i].ManagementAddress
		}
		// need to delete replication rule created earlier with the remoteIP not able to create Protection policy
		C.DeleteReplicationRule(context.Background(), suite.rr.ID)
		suite.rr.ID = ""
	}
	t.Skip("Skipping test as there are no working remote systems configured on array.")
	return "", ""
}

func (suite *ReplicationTestSuiteSync) TestReplicationSync() {
	t := suite.T()

	// get the remote powerstore system
	remoteSystem := suite.remoteSystem
	rs, err := C.GetRemoteSystem(context.Background(), remoteSystem)
	assert.NoError(t, err)
	assert.Equal(t, rs.ID, remoteSystem)

	isWriteOrderConsistent := true
	// create a volume group with a protection policy
	// A protection policy with a synchronous replication rule can only be applied to a write-order consistent volume group
	suite.vg, err = C.CreateVolumeGroup(context.Background(), &gopowerstore.VolumeGroupCreate{
		Name:                   "intcsi" + suite.randomString + "-vgtst",
		ProtectionPolicyID:     suite.pp.ID,
		IsWriteOrderConsistent: &isWriteOrderConsistent,
	})
	assert.NoError(t, err)

	// create a volume within the volume group
	volName := "intcsi" + suite.randomString + "-voltst"
	size := int64(OneMB)
	suite.vol, err = C.CreateVolume(context.Background(), &gopowerstore.VolumeCreate{
		Name:          &volName,
		Size:          &size,
		VolumeGroupID: suite.vg.ID,
	})
	assert.NoError(t, err)

	// get the volume group from the volume ID
	volID := suite.vol.ID
	_, err = C.GetVolumeGroupsByVolumeID(context.Background(), volID)
	assert.NoError(t, err)
}

func TestGetRemoteSystemsByFilter(t *testing.T) {
	allSys, err := C.GetAllRemoteSystems(context.Background())
	checkAPIErr(t, err)
	if len(allSys) == 0 {
		t.Skip("Skipping test as there are no remote systems configured on array.")
	}

	// check filter positive
	sysByMgmtAddr, err := C.GetRemoteSystems(context.Background(), map[string]string{
		"management_address": "eq." + allSys[0].ManagementAddress,
	})
	checkAPIErr(t, err)
	assert.Equal(t, len(sysByMgmtAddr), 1)
	assert.EqualExportedValues(t, allSys[0].ID, sysByMgmtAddr[0].ID)

	// check filter negative
	_, err = C.GetRemoteSystems(context.Background(), map[string]string{
		"management_address": "inv." + allSys[0].ManagementAddress,
	})
	assert.Error(t, err)
}

func TestReplicationSuiteSync(t *testing.T) {
	suite.Run(t, new(ReplicationTestSuiteSync))
}
