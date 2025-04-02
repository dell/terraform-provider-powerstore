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
	replicationRuleURL = "replication_rule"
)

// GetReplicationRules returns a list of replication rules
// TODO: Move to using auto-generated client library
func (c *Client) GetReplicationRules(ctx context.Context, args map[string]string) ([]gopowerstore.ReplicationRule, error) {
	var result []gopowerstore.ReplicationRule
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.ReplicationRule
		policy := gopowerstore.ReplicationRule{}
		qp := c.PStoreClient.APIClient().QueryParamsWithFields(&policy)
		for k, v := range args {
			qp = qp.RawArg(k, v)
		}
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    replicationRuleURL,
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
