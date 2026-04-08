/*
 *
 * Copyright © 2020-2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	hostMockURL        = APIMockURL + hostURL
	hostMappingMockURL = APIMockURL + hostMappingURL
)

var (
	hostID  = "6b930711-46bc-4a4b-9d6a-22c77a7838c4"
	hostID2 = "3765da74-28a7-49db-a693-10cec1de91f8"
)

func TestClientIMPL_GetHosts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, hostID, hostID2)
	httpmock.RegisterResponder("GET", hostMockURL,
		httpmock.NewStringResponder(200, respData))
	vols, err := C.GetHosts(context.Background())
	assert.Nil(t, err)
	assert.Len(t, vols, 2)
	assert.Equal(t, hostID, vols[0].ID)
}

func TestClientIMPL_GetHost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, hostID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", hostMockURL, hostID),
		httpmock.NewStringResponder(200, respData))
	host, err := C.GetHost(context.Background(), hostID)
	assert.Nil(t, err)
	assert.Equal(t, hostID, host.ID)
}

func TestClientIMPL_GetHostByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", hostMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, hostID)
	setResponder(respData)
	host, err := C.GetHostByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, hostID, host.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetHostByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.HostIsNotExist())
}

func TestClientIMPL_CreateHost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, hostID)
	httpmock.RegisterResponder("POST", hostMockURL,
		httpmock.NewStringResponder(201, respData))
	name := "test_HOst"
	portName := "foobar"
	initiators := []InitiatorCreateModify{{PortName: &portName}}
	createReq := HostCreate{}
	createReq.Name = &name
	createReq.Initiators = &initiators

	resp, err := C.CreateHost(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, hostID, resp.ID)
}

func TestClientIMPL_DeleteHost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", hostMockURL, hostID),
		httpmock.NewStringResponder(204, ""))
	resp, err := C.DeleteHost(context.Background(), nil, hostID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_ModifyHost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, hostID)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", hostMockURL, hostID),
		httpmock.NewStringResponder(201, respData))
	resp, err := C.ModifyHost(context.Background(), nil, hostID)
	assert.Nil(t, err)
	assert.Equal(t, hostID, resp.ID)
}

func TestClientIMPL_GetHostVolumeMappings(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, hostID)
	httpmock.RegisterResponder("GET", hostMappingMockURL,
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetHostVolumeMappings(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, hostID, resp[0].ID)
}

func TestClientIMPL_GetHostVolumeMapping(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, hostID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", hostMappingMockURL, hostID),
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetHostVolumeMapping(context.Background(), hostID)
	assert.Nil(t, err)
	assert.Equal(t, hostID, resp.ID)
}

func TestClientIMPL_GetHostVolumeMappingByVolumeID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, hostID)
	httpmock.RegisterResponder("GET", hostMappingMockURL,
		httpmock.NewStringResponder(200, respData))
	resp, err := C.GetHostVolumeMappingByVolumeID(context.Background(), volID)
	assert.Nil(t, err)
	assert.Len(t, resp, 1)
	assert.Equal(t, hostID, resp[0].ID)
}

func TestClientIMPL_AttachVolumeToHost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/attach", hostMockURL, hostID),
		httpmock.NewStringResponder(204, ""))
	attach := HostVolumeAttach{}
	id := "06c16b46-b015-41a6-9d21-0c44863e395b"
	attach.VolumeID = &id
	_, err := C.AttachVolumeToHost(context.Background(), hostID, &attach)
	assert.Nil(t, err)
}

func TestClientIMPL_DetachVolumeFromHost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/detach", hostMockURL, hostID),
		httpmock.NewStringResponder(204, ""))
	detach := HostVolumeDetach{}
	id := "06c16b46-b015-41a6-9d21-0c44863e395b"
	detach.VolumeID = &id
	_, err := C.DetachVolumeFromHost(context.Background(), hostID, &detach)
	assert.Nil(t, err)
}
