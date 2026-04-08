/*
 *
 * Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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

package gopowerstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	remoteSystemMockURL = APIMockURL + remoteSystemURL
	clusterMockURL      = APIMockURL + clusterURL
)

func TestClientIMPL_GetAllRemoteSystems(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, volID, volID2)
	httpmock.RegisterResponder("GET", remoteSystemMockURL,
		httpmock.NewStringResponder(200, respData))
	remoteSystems, err := C.GetAllRemoteSystems(context.Background())
	assert.Nil(t, err)
	assert.Len(t, remoteSystems, 2)
	assert.Equal(t, volID, remoteSystems[0].ID)
}

func TestClientIMPL_GetRemoteSystems(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, volID, volID2)
	httpmock.RegisterResponder("GET", remoteSystemMockURL,
		httpmock.NewStringResponder(200, respData))
	remoteSystems, err := C.GetRemoteSystems(context.Background(), nil)
	assert.Nil(t, err)
	assert.Len(t, remoteSystems, 2)
	assert.Equal(t, volID, remoteSystems[0].ID)
}

func TestClientIMPL_GetCluster(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	httpmock.RegisterResponder("GET", clusterMockURL,
		httpmock.NewStringResponder(200, respData))
	cluster, err := C.GetCluster(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, volID, cluster.ID)
}

func TestClientIMPL_GetRemoteSystem(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", remoteSystemMockURL, volID),
		httpmock.NewStringResponder(200, respData))
	remoteSystem, err := C.GetRemoteSystem(context.Background(), volID)
	assert.Nil(t, err)
	assert.Equal(t, volID, remoteSystem.ID)
}

func TestClientIMPL_GetRemoteSystemByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", remoteSystemMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	remoteSystem, err := C.GetRemoteSystemByName(context.Background(), "XX-0000X")
	assert.Nil(t, err)
	assert.Equal(t, volID, remoteSystem.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetRemoteSystemByName(context.Background(), "XX-0000X")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.HostIsNotExist())
}
