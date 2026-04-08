/*
 *
 * Copyright © 2022-2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	hostGroupMockURL = APIMockURL + hostGroupURL
)

var (
	hostGroupID  = "6b930711-46bc-4a4b-9d6a-22c77a7838c4"
	hostGroupID2 = "3765da74-28a7-49db-a693-10cec1de91f8"
)

func TestClientIMPL_AttachVolumeToHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/attach", hostGroupMockURL, hostGroupID),
		httpmock.NewStringResponder(204, ""))
	attach := HostVolumeAttach{}
	id := "06c16b46-b015-41a6-9d21-0c44863e395b"
	attach.VolumeID = &id
	_, err := C.AttachVolumeToHostGroup(context.Background(), hostGroupID, &attach)
	assert.Nil(t, err)
}

func TestClientIMPL_DetachVolumeFromHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/detach", hostGroupMockURL, hostGroupID),
		httpmock.NewStringResponder(204, ""))
	detach := HostVolumeDetach{}
	id := "06c16b46-b015-41a6-9d21-0c44863e395b"
	detach.VolumeID = &id
	_, err := C.DetachVolumeFromHostGroup(context.Background(), hostGroupID, &detach)
	assert.Nil(t, err)
}

func TestClientIMPL_GetHostGroupByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", hostGroupMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, hostGroupID)
	setResponder(respData)
	hostGroup, err := C.GetHostGroupByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, hostGroupID, hostGroup.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetHostByName(context.Background(), "test")
	assert.NotNil(t, err)
}

func TestClientIMPL_GetHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, hostGroupID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", hostGroupMockURL, hostGroupID),
		httpmock.NewStringResponder(200, respData))
	hostGroup, err := C.GetHostGroup(context.Background(), hostGroupID)
	assert.Nil(t, err)
	assert.Equal(t, hostGroupID, hostGroup.ID)
}

func TestClientIMPL_GetHostGroups(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, hostGroupID, hostGroupID2)
	httpmock.RegisterResponder("GET", hostGroupMockURL,
		httpmock.NewStringResponder(200, respData))
	hostGroups, err := C.GetHostGroups(context.Background())
	assert.Nil(t, err)
	assert.Len(t, hostGroups, 2)
	assert.Equal(t, hostGroupID, hostGroups[0].ID)
}

func TestClientIMPL_CreateHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, hostGroupID)
	httpmock.RegisterResponder("POST", hostGroupMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := HostGroupCreate{
		Name:        "hg-test",
		Description: "create hg-test",
		HostIDs:     []string{hostID},
	}

	resp, err := C.CreateHostGroup(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, hostGroupID, resp.ID)
}

func TestClientIMPL_DeleteHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", hostGroupMockURL, hostGroupID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteHostGroup(context.Background(), hostGroupID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_ModifyHostGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(``)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", hostGroupMockURL, hostGroupID),
		httpmock.NewStringResponder(201, respData))

	modifyParams := HostGroupModify{
		AddHostIDs: []string{hostID2},
	}

	resp, err := C.ModifyHostGroup(context.Background(), &modifyParams, hostGroupID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}
