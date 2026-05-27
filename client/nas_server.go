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
	nasURL = "nas_server"
)

// getNASFieldsWithoutHealthDetails returns NAS fields excluding health_details
// which is not returned by the API and causes unmarshal errors
func getNASFieldsWithoutHealthDetails() []string {
	fields := []string{
		"id", "description", "name", "current_node_id", "operational_status",
		"current_preferred_IPv4_interface_id", "current_preferred_IPv6_interface_id",
		"preferred_node_id", "default_unix_user", "default_windows_user",
		"current_unix_directory_service", "is_username_translation_enabled",
		"is_auto_user_mapping_enabled", "production_IPv4_interface_id",
		"production_IPv6_interface_id", "backup_IPv4_interface_id",
		"backup_IPv6_interface_id", "protection_policy_id",
		"file_events_publishing_mode", "is_replication_destination",
		"is_production_mode_enabled", "is_dr_test", "operational_status_l10n",
		"current_unix_directory_service_l10n", "file_events_publishing_mode_l10n",
	}
	return fields
}

// GetNASWithoutHealthDetails fetches a single NAS server by ID without health_details field
func (c *Client) GetNASWithoutHealthDetails(ctx context.Context, id string) (gopowerstore.NAS, error) {
	var result gopowerstore.NAS
	qp := c.PStoreClient.APIClient().QueryParams()
	qp.Select(getNASFieldsWithoutHealthDetails()...)
	_, err := c.PStoreClient.APIClient().Query(
		ctx,
		gopowerstore.RequestConfig{
			Method:      "GET",
			Endpoint:    nasURL + "/" + id,
			QueryParams: qp,
		},
		&result)
	err = gopowerstore.WrapErr(err)
	return result, err
}

// GetNASByNameWithoutHealthDetails fetches a single NAS server by name without health_details field
func (c *Client) GetNASByNameWithoutHealthDetails(ctx context.Context, name string) (gopowerstore.NAS, error) {
	var results []gopowerstore.NAS
	qp := c.PStoreClient.APIClient().QueryParams()
	qp.Select(getNASFieldsWithoutHealthDetails()...)
	qp.RawArg("name", "eq."+name)
	_, err := c.PStoreClient.APIClient().Query(
		ctx,
		gopowerstore.RequestConfig{
			Method:      "GET",
			Endpoint:    nasURL,
			QueryParams: qp,
		},
		&results)
	err = gopowerstore.WrapErr(err)
	if err != nil {
		return gopowerstore.NAS{}, err
	}
	if len(results) == 0 {
		return gopowerstore.NAS{}, gopowerstore.NewNotFoundError()
	}
	return results[0], nil
}

// GetNASServersWithoutHealthDetails fetches all NAS servers without health_details field
func (c *Client) GetNASServersWithoutHealthDetails(ctx context.Context) ([]gopowerstore.NAS, error) {
	var result []gopowerstore.NAS
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.NAS
		qp := c.PStoreClient.APIClient().QueryParams()
		qp.Select(getNASFieldsWithoutHealthDetails()...)
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    nasURL,
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

// GetNaSServersByFilter fetches NAS servers matching the given filters
func (c *Client) GetNaSServersByFilter(ctx context.Context, filters map[string]string) ([]gopowerstore.NAS, error) {
	var result []gopowerstore.NAS
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.NAS
		qp := c.PStoreClient.APIClient().QueryParams()
		qp.Select(getNASFieldsWithoutHealthDetails()...)
		for k, v := range filters {
			qp.RawArg(k, v)
		}
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    nasURL,
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
