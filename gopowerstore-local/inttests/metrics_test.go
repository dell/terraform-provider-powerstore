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

package inttests

import (
	"context"
	"testing"
	"time"

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
)

func TestGetCapacity(t *testing.T) {
	resp, err := C.GetCapacity(context.Background())
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
}

func Test_PerformanceMetricsByAppliance(t *testing.T) {
	resp, err := C.PerformanceMetricsByAppliance(context.Background(), "A1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "performance_metrics_by_appliance", resp[0].Entity)
}

func Test_PerformanceMetricsByNode(t *testing.T) {
	resp, err := C.PerformanceMetricsByNode(context.Background(), "N1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "performance_metrics_by_node", resp[0].Entity)
}

func Test_PerformanceMetricsByVolume(t *testing.T) {
	volumesResp, err := C.GetVolumes(context.Background())
	checkAPIErr(t, err)
	if len(volumesResp) == 0 {
		t.Skip("no volumes are available to check for metrics")
		return
	}
	resp, err := C.PerformanceMetricsByVolume(context.Background(), volumesResp[0].ID, gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "performance_metrics_by_volume", resp[0].Entity)
}

func Test_VolumeMirrorTransferRate(t *testing.T) {
	volumesResp, err := C.GetVolumes(context.Background())
	checkAPIErr(t, err)
	if len(volumesResp) == 0 {
		t.Skip("no volumes are available to check for metrics")
		return
	}
	resp, err := C.VolumeMirrorTransferRate(context.Background(), volumesResp[0].ID)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
}

func Test_PerformanceMetricsByCluster(t *testing.T) {
	resp, err := C.PerformanceMetricsByCluster(context.Background(), "0", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "performance_metrics_by_cluster", resp[0].Entity)
}

func Test_PerformanceMetricsByVm(t *testing.T) {
	t.Skip("need a way to get a valid Virtual Machine before this test can be written")
}

func Test_PerformanceMetricsByVg(t *testing.T) {
	t.Skip("need a way to get a valid Volume Group before this test can be written")
}

func Test_PerformanceMetricsByFEFCPort(t *testing.T) {
	portsResp, err := C.GetFCPorts(context.Background())
	checkAPIErr(t, err)
	if len(portsResp) == 0 {
		t.Skip("no ports are available to check for metrics")
		return
	}
	resp, err := C.PerformanceMetricsByFeFcPort(context.Background(), portsResp[0].ID, gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "performance_metrics_by_fe_fc_port", resp[0].Entity)
}

func Test_PerformanceMetricsByFeEthPort(t *testing.T) {
	t.Skip("need a way to get a valid front-end ethernet port before this test can be written")
}

func Test_PerformanceMetricsByFeEthNode(t *testing.T) {
	resp, err := C.PerformanceMetricsByFeEthNode(context.Background(), "N1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "performance_metrics_by_fe_eth_node", resp[0].Entity)
}

func Test_PerformanceMetricsByFeFcNode(t *testing.T) {
	resp, err := C.PerformanceMetricsByFeFcNode(context.Background(), "N1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "performance_metrics_by_fe_fc_node", resp[0].Entity)
}

func Test_PerformanceMetricsByFileSystem(t *testing.T) {
	nasID, _ := createNAS(t)
	defer deleteNAS(t, nasID)

	fsID, _ := createFS(t, nasID)
	defer deleteFS(t, fsID)

	// need to wait until metrics are available for the newly created file system
	timeout := time.After(120 * time.Second)
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-timeout:
			t.Fatalf("timeout waiting for valid response")
		case <-tick:
			resp, err := C.PerformanceMetricsByFileSystem(context.Background(), fsID, gopowerstore.TwentySec)
			checkAPIErr(t, err)
			if len(resp) > 0 {
				assert.NotEmpty(t, resp)
				assert.Equal(t, "performance_metrics_by_file_system", resp[0].Entity)
				return
			}
		}
	}
}

func Test_PerformanceMetricsSmbByNode(t *testing.T) {
	t.Skip("need a way to get a valid SMB node before this test can be written")
}

func Test_PerformanceMetricsBySmbBuiltinclientByNode(t *testing.T) {
	t.Skip("need a way to get a valid SMB client before this test can be written")
}

func Test_PerformanceMetricsBySmbBranchCacheByNode(t *testing.T) {
	t.Skip("need a way to get a valid SMB branch cache before this test can be written")
}

func Test_PerformanceMetricsSmb1ByNode(t *testing.T) {
	t.Skip("need a way to get a valid SMB node before this test can be written")
}

func Test_PerformanceMetricsSmb1BuiltinclientByNode(t *testing.T) {
	t.Skip("need a way to get a valid SMB client before this test can be written")
}

func Test_PerformanceMetricsSmb2ByNode(t *testing.T) {
	t.Skip("need a way to get a valid SMB node before this test can be written")
}

func Test_PerformanceMetricsSmb2BuiltinclientByNode(t *testing.T) {
	t.Skip("need a way to get a valid SMB node before this test can be written")
}

func Test_PerformanceMetricsNfsByNode(t *testing.T) {
	resp, err := C.PerformanceMetricsNfsByNode(context.Background(), "N1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	if assert.NotEmpty(t, resp) {
		assert.Equal(t, "performance_metrics_nfs_by_node", resp[0].Entity)
	}
}

func Test_PerformanceMetricsNfsv3ByNode(t *testing.T) {
	resp, err := C.PerformanceMetricsNfsv3ByNode(context.Background(), "N1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	if assert.NotEmpty(t, resp) {
		assert.Equal(t, "performance_metrics_nfsv3_by_node", resp[0].Entity)
	}
}

func Test_PerformanceMetricsNfsv4ByNode(t *testing.T) {
	resp, err := C.PerformanceMetricsNfsv4ByNode(context.Background(), "N1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	if assert.NotEmpty(t, resp) {
		assert.Equal(t, "performance_metrics_nfsv4_by_node", resp[0].Entity)
	}
}

func Test_WearMetricsByDrive(t *testing.T) {
	t.Skip("no API available to get a drive ID")
}

func Test_SpaceMetricsByCluster(t *testing.T) {
	resp, err := C.SpaceMetricsByCluster(context.Background(), "0", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "space_metrics_by_cluster", resp[0].Entity)
}

func Test_SpaceMetricsByAppliance(t *testing.T) {
	resp, err := C.SpaceMetricsByAppliance(context.Background(), "A1", gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "space_metrics_by_appliance", resp[0].Entity)
}

func Test_SpaceMetricsByVolume(t *testing.T) {
	volumesResp, err := C.GetVolumes(context.Background())
	checkAPIErr(t, err)
	if len(volumesResp) == 0 {
		t.Skip("no volumes are available to check for metrics")
		return
	}
	resp, err := C.SpaceMetricsByVolume(context.Background(), volumesResp[0].ID, gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "space_metrics_by_volume", resp[0].Entity)
}

func Test_SpaceMetricsByVolumeFamily(t *testing.T) {
	volumesResp, err := C.GetVolumes(context.Background())
	if len(volumesResp) == 0 {
		t.Skip("no volumes are available to check for metrics")
		return
	}
	checkAPIErr(t, err)
	resp, err := C.SpaceMetricsByVolumeFamily(context.Background(), volumesResp[0].ID, gopowerstore.FiveMins)
	checkAPIErr(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, "space_metrics_by_volume_family", resp[0].Entity)
}

func Test_SpaceMetricsByVM(t *testing.T) {
	t.Skip("no API available to get a VM ID")
}

func Test_SpaceMetricsByStorageContainer(t *testing.T) {
	t.Skip("no API available to get a storage container ID")
}

func Test_SpaceMetricsByVolumeGroup(t *testing.T) {
	t.Skip("no API available to get a Volume Group ID")
}

func Test_CopyMetricsByAppliance(t *testing.T) {
	resp, err := C.CopyMetricsByAppliance(context.Background(), "A1", gopowerstore.OneDay)
	checkAPIErr(t, err)
	if assert.NotEmpty(t, resp) {
		assert.Equal(t, "copy_metrics_by_appliance", resp[0].Entity)
	}
}

func Test_CopyMetricsByCluster(t *testing.T) {
	resp, err := C.CopyMetricsByCluster(context.Background(), "0", gopowerstore.OneDay)
	checkAPIErr(t, err)
	if assert.NotEmpty(t, resp) {
		assert.Equal(t, "copy_metrics_by_cluster", resp[0].Entity)
	}
}

func Test_CopyMetricsByVolumeGroup(t *testing.T) {
	t.Skip("no API available to get a Volume Group ID")
}

func Test_CopyMetricsByRemoteSystem(t *testing.T) {
	t.Skip("no API available to get a Remote System ID")
}

func Test_CopyMetricsByVolume(t *testing.T) {
	t.Skip("All volumes at the momement does not have Copy Metrics")
}
