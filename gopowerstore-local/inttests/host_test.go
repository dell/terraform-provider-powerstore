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

package inttests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
)

const (
	testHostPrefix = "test_host_"
)

func generateFakeInitiatorName() string {
	return fmt.Sprintf("iqn.1994-05.com.dell:%s", strings.ToLower(randString(12)))
}

func checkFields(t *testing.T, host *gopowerstore.Host) {
	assert.NotEmpty(t, host.ID)
	assert.NotEmpty(t, host.Name)
}

func createHost(t *testing.T) (string, string) {
	hostName := testHostPrefix + randString(8)
	osType := gopowerstore.OSTypeEnumLinux
	portName := generateFakeInitiatorName()
	portType := gopowerstore.InitiatorProtocolTypeEnumISCSI
	initiatorData := []gopowerstore.InitiatorCreateModify{{PortName: &portName, PortType: &portType}}
	createParams := gopowerstore.HostCreate{Name: &hostName, Initiators: &initiatorData, OsType: &osType}
	resp, err := C.CreateHost(context.Background(), &createParams)
	if err != nil {
		t.FailNow()
	}
	return resp.ID, hostName
}

func deleteHost(t *testing.T, id string) {
	_, err := C.DeleteHost(context.Background(), nil, id)
	if err != nil {
		t.FailNow()
	}
}

func TestGetHosts(t *testing.T) {
	_, err := C.GetHosts(context.Background())
	checkAPIErr(t, err)
}

func TestGetHost(t *testing.T) {
	hostID, _ := createHost(t)
	h, err := C.GetHost(context.Background(), hostID)
	checkAPIErr(t, err)
	checkFields(t, &h)
	deleteHost(t, hostID)
}

func TestGetHostByName(t *testing.T) {
	hostID, hostName := createHost(t)
	h, err := C.GetHostByName(context.Background(), hostName)
	checkAPIErr(t, err)
	checkFields(t, &h)
	deleteHost(t, hostID)
}

func TestCreateHost(t *testing.T) {
	hostID, _ := createHost(t)
	deleteHost(t, hostID)
}

func TestModifyHost(t *testing.T) {
	hostID, _ := createHost(t)
	// add initiator
	portName := generateFakeInitiatorName()
	portType := gopowerstore.InitiatorProtocolTypeEnumISCSI

	addInitiators := []gopowerstore.InitiatorCreateModify{{PortName: &portName, PortType: &portType}}
	modifyParams := gopowerstore.HostModify{AddInitiators: &addInitiators}
	_, err := C.ModifyHost(context.Background(), &modifyParams, hostID)
	checkAPIErr(t, err)

	_, err = C.GetHost(context.Background(), hostID)
	checkAPIErr(t, err)
	// remove initiator
	removeInitiator := []string{portName}
	modifyParams = gopowerstore.HostModify{RemoveInitiators: &removeInitiator}
	_, err = C.ModifyHost(context.Background(), &modifyParams, hostID)
	checkAPIErr(t, err)

	// update chap
	user := "admin"
	password := "password"
	updateInitiator := []gopowerstore.UpdateInitiatorInHost{{
		PortName: &portName, ChapMutualPassword: &password,
		ChapSinglePassword: &password, ChapMutualUsername: &user, ChapSingleUsername: &user,
	}}
	modifyParams = gopowerstore.HostModify{ModifyInitiators: &updateInitiator}

	deleteHost(t, hostID)
}

func TestGetHostVolumeMappings(t *testing.T) {
	_, err := C.GetHostVolumeMappings(context.Background())
	assert.Nil(t, err)
}

func TestGetHostVolumeMapping(t *testing.T) {
	mapping, err := C.GetHostVolumeMappings(context.Background())
	assert.Nil(t, err)
	if len(mapping) == 0 {
		t.Skipf("volumes mappings not found")
		return
	}
	_, err = C.GetHostVolumeMapping(context.Background(), mapping[0].ID)
	assert.Nil(t, err)
}

func TestAttachDetachVolume(t *testing.T) {
	volID, _ := CreateVol(t)
	hostID, _ := createHost(t)
	// attach
	attach := gopowerstore.HostVolumeAttach{}
	attach.VolumeID = &volID
	_, err := C.AttachVolumeToHost(context.Background(), hostID, &attach)
	assert.Nil(t, err)

	// read volume mapping
	resp, err := C.GetHostVolumeMappingByVolumeID(context.Background(), volID)
	assert.Nil(t, err)
	assert.Equal(t, resp[0].HostID, hostID)

	// detach
	detach := gopowerstore.HostVolumeDetach{}
	detach.VolumeID = &volID
	_, err = C.DetachVolumeFromHost(context.Background(), hostID, &detach)
	assert.Nil(t, err)
	_, err = C.DetachVolumeFromHost(context.Background(), hostID, &detach)
	assert.NotNil(t, err)
	// try detach second time
	apiError := err.(gopowerstore.APIError)
	assert.True(t, apiError.HostIsNotAttachedToVolume())
	DeleteVol(t, volID)
	deleteHost(t, hostID)
	// try detach not exist host
	_, err = C.DetachVolumeFromHost(context.Background(), hostID, &detach)
	assert.NotNil(t, err)
	apiError = err.(gopowerstore.APIError)
	assert.True(t, apiError.HostIsNotExist())
}

func TestDeleteAttachedVolume(t *testing.T) {
	volID, _ := CreateVol(t)
	hostID, _ := createHost(t)
	// attach
	attach := gopowerstore.HostVolumeAttach{}
	attach.VolumeID = &volID
	_, err := C.AttachVolumeToHost(context.Background(), hostID, &attach)
	assert.Nil(t, err)

	// try to delete the volume
	_, err = C.DeleteVolume(context.Background(), nil, volID)
	assert.NotNil(t, err)
	apiError := err.(gopowerstore.APIError)
	assert.True(t, apiError.VolumeAttachedToHost())

	// detach
	mappings, err := C.GetHostVolumeMappingByVolumeID(context.Background(), volID)
	if err != nil {
		t.FailNow()
	}
	for _, m := range mappings {
		dp := &gopowerstore.HostVolumeDetach{VolumeID: &volID}
		_, err = C.DetachVolumeFromHost(context.Background(), m.HostID, dp)
		if err != nil {
			t.FailNow()
		}
	}
	// should pass
	_, err = C.DeleteVolume(context.Background(), nil, volID)
	assert.Nil(t, err)
}
