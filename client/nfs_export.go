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
	"terraform-provider-powerstore/models/jsonmodel"

	"github.com/dell/gopowerstore"
)

const (
	nfsURL = "nfs_export"
)

// ModifyNFSExport modifies nfs export
// TODO: Move to using auto-generated client library
func (c *Client) ModifyNFSExport(ctx context.Context,
	modifyParams *jsonmodel.NFSExportModify, id string,
) error {
	_, err := c.PStoreClient.APIClient().Query(
		ctx,
		gopowerstore.RequestConfig{
			Method:   "PATCH",
			Endpoint: nfsURL,
			ID:       id,
			Body:     modifyParams,
		},
		&gopowerstore.CreateResponse{})
	err = gopowerstore.WrapErr(err)
	return err
}
