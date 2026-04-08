/*
 *
 * Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	storageContainerMockURL = APIMockURL + storageContainerURL
)

var (
	scID  = "435669ba-28f5-4395-b5ca-6a7455726eaa"
	scID2 = "3765da74-28a7-49db-a693-10cec1de91f8"
)

func TestClientIMPL_CreateStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, scID)
	httpmock.RegisterResponder("POST", storageContainerMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := StorageContainer{
		Name: "sc-test",
	}

	resp, err := C.CreateStorageContainer(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, scID, resp.ID)
}

func TestClientIMPL_DeleteStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", storageContainerMockURL, scID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteStorageContainer(context.Background(), scID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_GetStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, scID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", storageContainerMockURL, scID),
		httpmock.NewStringResponder(200, respData))
	storageContainer, err := C.GetStorageContainer(context.Background(), scID)
	assert.Nil(t, err)
	assert.Equal(t, scID, storageContainer.ID)
}

func TestClientIMPL_ModifyStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", storageContainerMockURL, scID),
		httpmock.NewStringResponder(201, ""))

	modifyParams := StorageContainer{
		Quota: 420,
	}

	resp, err := C.ModifyStorageContainer(context.Background(), &modifyParams, scID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}
