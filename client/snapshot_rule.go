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

	"github.com/dell/gopowerstore"
	"github.com/dell/gopowerstore/api"
)

const (
	snapshotRuleURL = "snapshot_rule"
)

func (c *Client) GetSnapshotRuleByFilter(ctx context.Context, filters map[string]string) ([]gopowerstore.SnapshotRule, error) {
	var result []gopowerstore.SnapshotRule
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.SnapshotRule
		snapshotRule := gopowerstore.SnapshotRule{}
		qp := c.PStoreClient.APIClient().QueryParamsWithFields(&snapshotRule)
		for k, v := range filters {
			qp.RawArg(k, v)
		}
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    snapshotRuleURL,
				QueryParams: qp,
			},
			&page)
		err = gopowerstore.WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}
