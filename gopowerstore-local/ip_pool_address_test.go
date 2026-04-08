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
	ipPoolAddressMockURL = APIMockURL + apiPoolAddressURL
)

func TestClientIMPL_GetIPPoolAddress(t *testing.T) {
	id1 := "IP1"
	purpose1 := IPPurposeTypeEnumStorageIscsiTarget

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s", "purpose": "%s"}]`,
		id1, purpose1)
	httpmock.RegisterResponder("GET", ipPoolAddressMockURL,
		httpmock.NewStringResponder(200, respData))

	vols, err := C.GetStorageISCSITargetAddresses(context.Background())
	assert.Nil(t, err)
	assert.Len(t, vols, 1)
	assert.Equal(t, id1, vols[0].ID)
}

func TestClientIMPL_GetIPPoolAddress_NotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[]`)
	httpmock.RegisterResponder("GET", ipPoolAddressMockURL,
		httpmock.NewStringResponder(200, respData))

	_, err := C.GetStorageISCSITargetAddresses(context.Background())
	assert.NotNil(t, err)
}

func TestClientIMPL_GetStorageNVMETCPTargetAddresses(t *testing.T) {
	want := []IPPoolAddress{
		{ID: "ip1"},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", ipPoolAddressMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := `[{"id": "ip1"}]`
	setResponder(respData)
	resp, err := C.GetStorageNVMETCPTargetAddresses(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, resp, want)
	respData = `[]`
	setResponder(respData)
	_, err = C.GetStorageNVMETCPTargetAddresses(context.Background())
	assert.NotNil(t, err)
	httpmock.Reset()
}
