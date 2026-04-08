/*
 *
 * Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	smbShareID            = "6732e829-29c9-7fed-686a-ee23cab1d298"
	smbShareMockURL       = APIMockURL + smbShareURL
	smbShareSetACLMockURL = APIMockURL + smbShareURL + "/" + smbShareID + smbShareSetACLURL
	smbShareGetACLMockURL = APIMockURL + smbShareURL + "/" + smbShareID + smbShareGetACLURL
)

func TestClientIMPL_CreateSMBShare(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, smbShareID)
	httpmock.RegisterResponder("POST", smbShareMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := SMBShareCreate{
		FileSystemID: "6721f37d-8eda-7508-664c-ee23cab1d298",
		Name:         "test_smb_share",
		Path:         "XX-0000X",
	}

	resp, err := C.CreateSMBShare(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, smbShareID, resp.ID)
}

func TestClientIMPL_ModifySMBShare(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", smbShareMockURL, smbShareID),
		httpmock.NewStringResponder(204, ""))

	desc := "modify smb share"
	flag := true
	modifyParams := SMBShareModify{
		Description:                     &desc,
		IsContinuousAvailabilityEnabled: &flag,
	}

	resp, err := C.ModifySMBShare(context.Background(), smbShareID, &modifyParams)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_DeleteSMBShare(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", smbShareMockURL, smbShareID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteSMBShare(context.Background(), smbShareID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_GetSMBShare(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, smbShareID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", smbShareMockURL, smbShareID),
		httpmock.NewStringResponder(200, respData))
	smbShare, err := C.GetSMBShare(context.Background(), smbShareID)
	assert.Nil(t, err)
	assert.Equal(t, smbShareID, smbShare.ID)
}

func TestClientIMPL_GetSMBShareByFilter(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}]`, smbShareID)
	httpmock.RegisterResponder("GET", smbShareMockURL,
		httpmock.NewStringResponder(200, respData))
	smbShares, err := C.GetSMBShares(context.Background(), map[string]string{"name": "ilike.test_*"})
	assert.Nil(t, err)
	assert.Equal(t, smbShareID, smbShares[0].ID)
}

func TestClientIMPL_SetSMBShareAcl(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", smbShareSetACLMockURL,
		httpmock.NewStringResponder(204, ""))

	createReq := &ModifySMBShareACL{
		AddAces: []SMBShareAce{
			{
				TrusteeType: "WellKnown",
				TrusteeName: "Everyone",
				AccessLevel: "Full",
				AccessType:  "Allow",
			},
		},
	}

	resp, err := C.SetSMBShareACL(context.Background(), smbShareID, createReq)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_GetSMBShareACL(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respData := `{"aces": [{"trustee_type": "WellKnown", "trustee_name": "Everyone", "access_level": "Full", "access_type": "Allow"}]}`

	httpmock.RegisterResponder("POST", smbShareGetACLMockURL,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetSMBShareACL(context.Background(), smbShareID)
	assert.Nil(t, err)
	assert.NotNil(t, resp.Aces)
	assert.NotEqual(t, len(resp.Aces), 0)
}
