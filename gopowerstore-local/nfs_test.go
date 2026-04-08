/*
 *
 * Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	nfsMockURL       = APIMockURL + nfsURL
	nfsServerMockURL = APIMockURL + nfsServerURL
	fileMockURL      = APIMockURL + fileInterfaceURL
)

var (
	nfsID       = "3765da74-28a7-49db-a693-10cec1de91f8"
	nfsServerID = "5e8d8e8e-671b-336f-db4e-cee0fbdc981e"
)

func TestClientIMPL_GetNFSExport(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, nfsID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", nfsMockURL, nfsID),
		httpmock.NewStringResponder(200, respData))
	nfs, err := C.GetNFSExport(context.Background(), nfsID)
	assert.Nil(t, err)
	assert.Equal(t, nfsID, nfs.ID)
}

func TestClientIMPL_GetNFSExportByFilter(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, nfsID)
	httpmock.RegisterResponder("GET", nfsMockURL,
		httpmock.NewStringResponder(200, respData))
	nfs, err := C.GetNFSExportByFilter(context.Background(), nil)
	assert.Nil(t, err)
	assert.Equal(t, nfsID, nfs[0].ID)
}

func TestClientIMPL_GetNFSExportByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", nfsMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, nfsID)
	setResponder(respData)
	nfs, err := C.GetNFSExportByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, nfsID, nfs.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetNFSExportByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_CreateNFSExport(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, nfsID)
	httpmock.RegisterResponder("POST", nfsMockURL,
		httpmock.NewStringResponder(201, respData))
	createReq := NFSExportCreate{
		Name:         "new-nfs",
		Description:  "some description",
		FileSystemID: fsID,
		Path:         "/nfs-export",
	}

	nfs, err := C.CreateNFSExport(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, nfsID, nfs.ID)
}

func TestClientIMPL_DeleteNFSExport(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", nfsMockURL, nfsID),
		httpmock.NewStringResponder(204, ""))
	resp, err := C.DeleteNFSExport(context.Background(), nfsID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_ModifyNFSExport(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, nfsID)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", nfsMockURL, nfsID),
		httpmock.NewStringResponder(201, respData))
	modifyReq := NFSExportModify{
		AddRWRootHosts:    []string{"192.168.100.10", "192.168.100.11"},
		RemoveRWRootHosts: []string{"192.168.100.9"},
		AddNoAccessHosts:  []string{"127.0.0.1"},
	}
	resp, err := C.ModifyNFSExport(context.Background(), &modifyReq, nfsID)
	assert.Nil(t, err)
	assert.Equal(t, nfsID, resp.ID)
}

func TestClientIMPL_CreateNFSServer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, nfsServerID)
	httpmock.RegisterResponder("POST", nfsServerMockURL,
		httpmock.NewStringResponder(201, respData))
	createReq := NFSServerCreate{
		NasServerID:     nasID,
		HostName:        "nfs-server",
		IsNFSv3Enabled:  true,
		IsNFSv4Enabled:  true,
		IsSecureEnabled: false,
	}

	nfsServer, err := C.CreateNFSServer(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, nfsServerID, nfsServer.ID)
}

func TestClientIMPL_GetNFSExportByFileSystemID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", nfsMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, nfsID)
	setResponder(respData)
	nfs, err := C.GetNFSExportByFileSystemID(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, nfsID, nfs.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetNFSExportByFileSystemID(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetFileInterface(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", fileMockURL, "test"),
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`{"id": "%s"}`, nfsID)
	setResponder(respData)
	fileInterface, err := C.GetFileInterface(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, nfsID, fileInterface.ID)
}
