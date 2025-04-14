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
	"fmt"

	"github.com/dell/gopowerstore"
	"github.com/dell/gopowerstore/api"
)

const (
	volumeURL = "volume"
)

func (c *Client) GetVolumesByFilter(ctx context.Context, filters map[string]string) ([]gopowerstore.Volume, error) {
	var result []gopowerstore.Volume
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.Volume
		volume := gopowerstore.Volume{}
		qp := c.PStoreClient.APIClient().QueryParamsWithFields(&volume)
		for k, v := range filters {
			qp.RawArg(k, v)
		}
		qp.RawArg("type", fmt.Sprintf("not.eq.%s", "Snapshot"))
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    volumeURL,
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

func (c *Client) GetVolumesSnapshotsByFilter(ctx context.Context, filters map[string]string) ([]gopowerstore.Volume, error) {
	var result []gopowerstore.Volume
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.Volume
		volume := gopowerstore.Volume{}
		qp := c.PStoreClient.APIClient().QueryParamsWithFields(&volume)
		for k, v := range filters {
			qp.RawArg(k, v)
		}
		qp.RawArg("type", fmt.Sprintf("eq.%s", "Snapshot"))
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    volumeURL,
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
