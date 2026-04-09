/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"

	pstore "github.com/dell/gopowerstore"
)

// isSecureResp is used to decode only the is_secure field from API responses.
// For volumes and volume groups, is_secure is nested inside protection_data.
// For filesystems and snapshot rules, is_secure is a top-level field.
type isSecureResp struct {
	IsSecure       bool `json:"is_secure"`
	ProtectionData struct {
		IsSecure bool `json:"is_secure"`
	} `json:"protection_data"`
}

// FetchIsSecure queries the PowerStore API for the is_secure field of a resource.
// This is needed because the upstream gopowerstore SDK does not yet include
// IsSecure fields on ProtectionData, FileSystem, or SnapshotRule structs.
// Returns false if the resource does not support is_secure or on error.
func (c *Client) FetchIsSecure(ctx context.Context, endpoint, id string) bool {
	var resp isSecureResp
	qp := c.PStoreClient.APIClient().QueryParams()
	qp.Select("id", "is_secure", "protection_data")
	_, err := c.PStoreClient.APIClient().Query(
		ctx,
		pstore.RequestConfig{
			Method:      "GET",
			Endpoint:    endpoint,
			ID:          id,
			QueryParams: qp,
		},
		&resp)
	if err != nil {
		return false
	}
	return resp.IsSecure || resp.ProtectionData.IsSecure
}
