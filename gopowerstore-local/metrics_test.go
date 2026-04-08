/*
 *
 * Copyright © 2020-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	metricsMockURL        = APIMockURL + metricsURL
	metricsMockVolMirrURL = APIMockURL + mirrorURL
	volumeID              = "4ffcd8e8-2a93-49ed-b9b3-2e68c8ddc5e4"
)

func TestClientIMPL_GetCapacity(t *testing.T) {
	totalSpace0 := 12077448036352
	usedSpace0 := 1905262588
	totalSpace1 := 12077448036352
	usedSpace1 := 5905262588
	totalSpace2 := 12077448036352
	usedSpace2 := 9905262588
	freeSpace := int64(totalSpace2 - usedSpace2)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"physical_total": %d, "physical_used": %d},{"physical_total": %d, "physical_used": %d},{"physical_total": %d, "physical_used": %d}]`,
		totalSpace0, usedSpace0, totalSpace1, usedSpace1, totalSpace2, usedSpace2)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetCapacity(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, freeSpace, resp)
}

func TestClientIMPL_GetCapacity_Zero(t *testing.T) {
	totalSpace0 := 12077448036352
	usedSpace0 := 10077448036352
	totalSpace1 := 12077448036352
	usedSpace1 := 11077448036352
	totalSpace2 := 12077448036352
	usedSpace2 := 12077448036353
	var freeSpace int64

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"physical_total": %d, "physical_used": %d},{"physical_total": %d, "physical_used": %d},{"physical_total": %d, "physical_used": %d}]`,
		totalSpace0, usedSpace0, totalSpace1, usedSpace1, totalSpace2, usedSpace2)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetCapacity(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, freeSpace, resp)
}

func TestClientIMPL_PerformanceMetricsByAppliance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_appliance", "appliance_id": "A1", "avg_read_latency": 0.0123}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_appliance")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByAppliance(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_appliance", resp[0].Entity)
	assert.Equal(t, "A1", resp[0].ApplianceID)
	assert.Equal(t, float32(0.0123), resp[0].AvgReadLatency)
}

func TestClientIMPL_PerformanceMetricsByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_node", "node_id": "Node-1", "current_logins": 123}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, int64(123), *resp[0].CurrentLogins)
}

func TestClientIMPL_PerformanceMetricsByVolume(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_volume", "volume_id": "Volume-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_volume")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByVolume(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_volume", resp[0].Entity)
	assert.Equal(t, "Volume-1", resp[0].VolumeID)
	assert.Equal(t, float64(1000), resp[0].TotalIops)
}

func TestClientIMPL_VolumeMirrorTransferRate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", metricsMockVolMirrURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volumeID)
	setResponder(respData)
	volMirr, err := C.VolumeMirrorTransferRate(context.Background(), volumeID)
	assert.Nil(t, err)
	assert.Equal(t, volumeID, volMirr[0].ID)
}

func TestClientIMPL_PerformanceMetricsByCluster(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_cluster", "cluster_id": "Cluster-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_cluster")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByCluster(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_cluster", resp[0].Entity)
	assert.Equal(t, "Cluster-1", resp[0].ClusterID)
	assert.Equal(t, float64(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricsByVM(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_vm", "vm_id": "Virtual-Machine-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_vm")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByVM(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_vm", resp[0].Entity)
	assert.Equal(t, "Virtual-Machine-1", resp[0].VMID)
	assert.Equal(t, float64(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricsByVg(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_vg", "vg_id": "Volume-Group-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_vg")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByVg(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_vg", resp[0].Entity)
	assert.Equal(t, "Volume-Group-1", resp[0].VgID)
	assert.Equal(t, float32(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricsByFeFcPort(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_fe_fc_port", "fe_port_id": "FE-Port-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_fe_fc_port")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByFeFcPort(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_fe_fc_port", resp[0].Entity)
	assert.Equal(t, "FE-Port-1", resp[0].FePortID)
	assert.Equal(t, float64(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricsByFeEthPort(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_fe_eth_port", "fe_port_id": "FE-Eth-Port-1", "pkt_rx_ps": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_fe_eth_port")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByFeEthPort(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_fe_eth_port", resp[0].Entity)
	assert.Equal(t, "FE-Eth-Port-1", resp[0].FePortID)
	assert.Equal(t, float32(1000), resp[0].PktRxPs)
}

func TestClientIMPL_PerformanceMetricsByFeEthNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_fe_eth_node", "node_id": "Node-1", "pkt_rx_ps": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_fe_eth_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByFeEthNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_fe_eth_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].PktRxPs)
}

func TestClientIMPL_PerformanceMetricsByFeFcNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_fe_fc_node", "node_id": "Node-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_fe_fc_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByFeFcNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_fe_fc_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float64(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricsByFileSystem(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_file_system", "file_system_id": "File-System-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_file_system")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsByFileSystem(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_file_system", resp[0].Entity)
	assert.Equal(t, "File-System-1", resp[0].FileSystemID)
	assert.Equal(t, float32(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricsSmbByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_by_smb_node", "node_id": "Node-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_by_smb_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsSmbByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_by_smb_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricSmbBuiltinclientByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_smb_builtinclient_by_node", "node_id": "Node-1", "total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_smb_builtinclient_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsSmbBuiltinclientByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_smb_builtinclient_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].TotalIops)
}

func TestClientIMPL_PerformanceMetricsSmbBranchCacheByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_smb_branch_cache_by_node", "node_id": "Node-1", "max_used_threads": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_smb_branch_cache_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsSmbBranchCacheByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_smb_branch_cache_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxUsedThreads)
}

func TestClientIMPL_PerformanceMetricsSmb1ByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_smb1_by_node", "node_id": "Node-1", "max_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_smb1_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsSmb1ByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_smb1_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxIops)
}

func TestClientIMPL_PerformanceMetricsSmb1BuiltinclientByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_smb1_builtinclient_by_node", "node_id": "Node-1", "max_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_smb1_builtinclient_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsSmb1BuiltinclientByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_smb1_builtinclient_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxIops)
}

func TestClientIMPL_PerformanceMetricsSmb2ByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_smb2_by_node", "node_id": "Node-1", "max_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_smb2_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsSmb2ByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_smb2_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxIops)
}

func TestClientIMPL_PerformanceMetricsSmb2BuiltinclientByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_smb2_builtinclient_by_node", "node_id": "Node-1", "max_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_smb2_builtinclient_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsSmb2BuiltinclientByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_smb2_builtinclient_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxIops)
}

func TestClientIMPL_PerformanceMetricsNfsByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_nfs_by_node", "node_id": "Node-1", "max_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_nfs_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsNfsByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_nfs_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxIops)
}

func TestClientIMPL_PerformanceMetricsNfsv3ByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_nfsv3_by_node", "node_id": "Node-1", "max_total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_nfsv3_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsNfsv3ByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_nfsv3_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxTotalIops)
}

func TestClientIMPL_PerformanceMetricsNfsv4ByNode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "performance_metrics_nfsv4_by_node", "node_id": "Node-1", "max_total_iops": 1000}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			err := decoder.Decode(&b)
			assert.Nil(t, err)
			assert.Equal(t, b.Entity, "performance_metrics_nfsv4_by_node")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.PerformanceMetricsNfsv4ByNode(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "performance_metrics_nfsv4_by_node", resp[0].Entity)
	assert.Equal(t, "Node-1", resp[0].NodeID)
	assert.Equal(t, float32(1000), resp[0].MaxTotalIops)
}

func TestClientIMPL_WearMetricsByDrive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "wear_metrics_by_drive", "drive_id": "Drive-1", "percent_endurance_remaining": 72.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "wear_metrics_by_drive")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.WearMetricsByDrive(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "Drive-1", resp[0].DriveID)
	assert.Equal(t, "wear_metrics_by_drive", resp[0].Entity)
	assert.Equal(t, float32(72.0), resp[0].PercentEnduranceRemaining)
}

func TestClientIMPL_SpaceMetricsByCluster(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "space_metrics_by_cluster", "cluster_id": "Cluster-1", "snapshot_savings": 72.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "space_metrics_by_cluster")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.SpaceMetricsByCluster(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "Cluster-1", resp[0].ClusterID)
	assert.Equal(t, "space_metrics_by_cluster", resp[0].Entity)
	assert.Equal(t, float32(72.0), resp[0].SnapshotSavings)
}

func TestClientIMPL_SpaceMetricsByAppliance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "space_metrics_by_appliance", "appliance_id": "Application-1", "snapshot_savings": 742.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "space_metrics_by_appliance")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.SpaceMetricsByAppliance(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "Application-1", resp[0].ApplianceID)
	assert.Equal(t, "space_metrics_by_appliance", resp[0].Entity)
	assert.Equal(t, float32(742.0), resp[0].SnapshotSavings)
}

func TestClientIMPL_SpaceMetricsByVolume(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "space_metrics_by_volume", "volume_id": "Volume-1", "max_thin_savings": 72.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "space_metrics_by_volume")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.SpaceMetricsByVolume(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "Volume-1", resp[0].VolumeID)
	assert.Equal(t, "space_metrics_by_volume", resp[0].Entity)
	assert.Equal(t, float32(72.0), resp[0].MaxThinSavings)
}

func TestClientIMPL_SpaceMetricsByVolumeFamily(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "space_metrics_by_volume_family", "family_id": "Family-1", "snapshot_savings": 72.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "space_metrics_by_volume_family")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.SpaceMetricsByVolumeFamily(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "Family-1", resp[0].FamilyID)
	assert.Equal(t, "space_metrics_by_volume_family", resp[0].Entity)
	assert.Equal(t, float32(72.0), resp[0].SnapshotSavings)
}

func TestClientIMPL_SpaceMetricsByVM(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "space_metrics_by_vm", "vm_id": "VM-1", "snapshot_savings": 72.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "space_metrics_by_vm")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.SpaceMetricsByVM(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "VM-1", resp[0].VMID)
	assert.Equal(t, "space_metrics_by_vm", resp[0].Entity)
	assert.Equal(t, float32(72.0), resp[0].SnapshotSavings)
}

func TestClientIMPL_SpaceMetricsByStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "space_metrics_by_storage_container", "storage_container_id": "Storage-1", "snapshot_savings": 72.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "space_metrics_by_storage_container")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.SpaceMetricsByStorageContainer(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "Storage-1", resp[0].StorageContainerID)
	assert.Equal(t, "space_metrics_by_storage_container", resp[0].Entity)
	assert.Equal(t, float32(72.0), resp[0].SnapshotSavings)
}

func TestClientIMPL_SpaceMetricsByVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "space_metrics_by_vg", "vg_id": "VG-1", "snapshot_savings": 562.0}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "space_metrics_by_vg")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.SpaceMetricsByVolumeGroup(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "VG-1", resp[0].VgID)
	assert.Equal(t, "space_metrics_by_vg", resp[0].Entity)
	assert.Equal(t, float32(562.0), resp[0].SnapshotSavings)
}

func TestClientIMPL_CopyMetricsByAppliance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "copy_metrics_by_appliance", "appliance_id": "Application-1", "repeat_count": 987654431}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "copy_metrics_by_appliance")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.CopyMetricsByAppliance(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "Application-1", resp[0].ApplianceID)
	assert.Equal(t, "copy_metrics_by_appliance", resp[0].Entity)
	assert.Equal(t, int32(987654431), *resp[0].RepeatCount)
}

func TestClientIMPL_CopyMetricsByCluster(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "copy_metrics_by_cluster", "repeat_count": 987654431}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "copy_metrics_by_cluster")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.CopyMetricsByCluster(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "copy_metrics_by_cluster", resp[0].Entity)
	assert.Equal(t, int32(987654431), *resp[0].RepeatCount)
}

func TestClientIMPL_CopyMetricsByVolumeGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "copy_metrics_by_vg", "vg_id": "VG-1", "repeat_count": 987654431}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "copy_metrics_by_vg")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.CopyMetricsByVolumeGroup(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "VG-1", resp[0].VgID)
	assert.Equal(t, "copy_metrics_by_vg", resp[0].Entity)
	assert.Equal(t, int32(987654431), *resp[0].RepeatCount)
}

func TestClientIMPL_CopyMetricsByRemoteSystem(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "copy_metrics_by_remote_system", "remote_system_id": "SYSTEM-1", "repeat_count": 987654431}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "copy_metrics_by_remote_system")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.CopyMetricsByRemoteSystem(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "SYSTEM-1", resp[0].RemoteSystemID)
	assert.Equal(t, "copy_metrics_by_remote_system", resp[0].Entity)
	assert.Equal(t, int32(987654431), *resp[0].RepeatCount)
}

func TestClientIMPL_CopyMetricsByVolume(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"entity": "copy_metrics_by_volume", "volume_id": "VOL-1", "repeat_count": 987654431}]`)
	httpmock.RegisterResponder("POST", metricsMockURL+"/generate",
		func(req *http.Request) (*http.Response, error) {
			decoder := json.NewDecoder(req.Body)
			var b MetricsRequest
			_ = decoder.Decode(&b)
			assert.Equal(t, b.Entity, "copy_metrics_by_volume")
			return httpmock.NewStringResponse(200, respData), nil
		},
	)
	resp, err := C.CopyMetricsByVolume(context.Background(), "", TwentySec)
	assert.Nil(t, err)
	assert.Equal(t, "VOL-1", resp[0].VolumeID)
	assert.Equal(t, "copy_metrics_by_volume", resp[0].Entity)
	assert.Equal(t, int32(987654431), *resp[0].RepeatCount)
}
