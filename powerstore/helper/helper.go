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

package helper

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"terraform-provider-powerstore/clientgen"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func GetKnownBoolPointer(in types.Bool) *bool {
	if in.IsUnknown() {
		return nil
	}
	return in.ValueBoolPointer()
}

func GetPointer[T any](in T) *T {
	return &in
}

// DetermineClusterTime fetches the cluster system time from the API
func DetermineClusterTime(ctx context.Context, client *clientgen.APIClient) (time.Time, error) {
	sel := "system_time"
	queries := make(url.Values)
	queries.Set("select", sel)
	clusterResponse, _, err := client.ClusterApi.GetAllClusters(ctx).Queries(queries).Execute()
	if err != nil {
		return time.Time{}, fmt.Errorf("could not fetch cluster time: %w", err)
	}
	if len(clusterResponse) == 0 {
		return time.Time{}, fmt.Errorf("cluster not found")
	}
	if clusterResponse[0].SystemTime == nil {
		return time.Time{}, fmt.Errorf("cluster system time is nil")
	}
	return *clusterResponse[0].SystemTime, nil
}
