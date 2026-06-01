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
