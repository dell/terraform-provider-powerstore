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

// VolumeTypeEnum Type of volume.
type VolumeTypeEnum string

const (
	volumeGroupURL = "volume_group"

	// VolumeTypeEnumPrimary - A base object.
	VolumeTypeEnumPrimary VolumeTypeEnum = "Primary"
	// VolumeTypeEnumClone - A read-write object that shares storage with the object from which it is sourced.
	VolumeTypeEnumClone VolumeTypeEnum = "Clone"
	// VolumeTypeEnumSnapshot - A read-only object created from a volume or clone.
	VolumeTypeEnumSnapshot VolumeTypeEnum = "Snapshot"
)

// TODO: Move to using auto-generated client library
// GetVolumeGroups returns a list of volume groups
func (c *Client) GetVolumeGroups(ctx context.Context, args map[string]string) ([]gopowerstore.VolumeGroup, error) {
	var result []gopowerstore.VolumeGroup
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.VolumeGroup
		volumegroup := gopowerstore.VolumeGroup{}
		qp := c.PStoreClient.APIClient().QueryParamsWithFields(&volumegroup)
		for k, v := range args {
			qp.RawArg(k, v)
		}
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    volumeGroupURL,
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

// GetVolumeGroupSnapshots returns all volume group snapshots
func (c *Client) GetVolumeGroupSnapshots(ctx context.Context, args map[string]string) ([]gopowerstore.VolumeGroup, error) {
	var result []gopowerstore.VolumeGroup
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.VolumeGroup
		volumegroup := gopowerstore.VolumeGroup{}
		qp := c.PStoreClient.APIClient().QueryParamsWithFields(&volumegroup)
		for k, v := range args {
			qp.RawArg(k, v)
		}
		qp.RawArg("type", fmt.Sprintf("eq.%s", VolumeTypeEnumSnapshot))
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    volumeGroupURL,
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
