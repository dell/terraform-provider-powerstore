/*
 *
 * Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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

	"github.com/dell/gopowerstore"
	"github.com/stretchr/testify/assert"
)

const TestNFSPrefix = "test_nfs_"

func createNFSServer(t *testing.T, nasID string) (string, string) {
	nfsName := TestNFSPrefix + randString(8)
	createParams := gopowerstore.NFSServerCreate{
		NasServerID:     nasID,
		HostName:        nfsName,
		IsNFSv3Enabled:  true,
		IsNFSv4Enabled:  true,
		IsSecureEnabled: false,
	}
	createResp, err := C.CreateNFSServer(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, nfsName
}

func createNFSExport(t *testing.T, fsID string, fsName string) (string, string) {
	nfsName := TestNFSPrefix + randString(8)
	createParams := gopowerstore.NFSExportCreate{
		Name:         nfsName,
		FileSystemID: fsID,
		Path:         "/" + fsName,
	}
	createResp, err := C.CreateNFSExport(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, nfsName
}

func deleteNFSExport(t *testing.T, id string) {
	_, err := C.DeleteNFSExport(context.Background(), id)
	checkAPIErr(t, err)
}

func TestGetNotExistingNFSExport(t *testing.T) {
	_, err := C.GetNFSExportByName(context.Background(), "some-random-name")
	assert.Error(t, err)
}

func TestGetNFSExportByFileSystemID(t *testing.T) {
	nasID, _ := createNAS(t)
	defer deleteNAS(t, nasID)

	fsID, fsName := createFS(t, nasID)
	defer deleteFS(t, fsID)

	_, _ = createNFSServer(t, nasID)

	nfsID, nfsName := createNFSExport(t, fsID, fsName)
	defer deleteNFSExport(t, nfsID)

	nfs, err := C.GetNFSExportByFileSystemID(context.Background(), fsID)
	assert.NoError(t, err)
	assert.NotEmpty(t, nfs.Name)
	assert.Equal(t, nfsName, nfs.Name)
}

func TestModifyNFSExport(t *testing.T) {
	nasID, _ := createNAS(t)

	fsID, fsName := createFS(t, nasID)

	_, _ = createNFSServer(t, nasID)

	nfsID, nfsName := createNFSExport(t, fsID, fsName)

	nfs, err := C.GetNFSExportByName(context.Background(), nfsName)
	checkAPIErr(t, err)
	assert.NotEmpty(t, nfs.Name)
	assert.Equal(t, nfsName, nfs.Name)

	t.Run("add host 1", func(t *testing.T) {
		modifyParams := gopowerstore.NFSExportModify{
			AddRWRootHosts: []string{"192.168.100.10"},
		}

		_, err = C.ModifyNFSExport(context.Background(), &modifyParams, nfsID)
		checkAPIErr(t, err)
	})

	t.Run("add host 2 and 3", func(t *testing.T) {
		modifyParams := gopowerstore.NFSExportModify{
			AddRWRootHosts: []string{"192.168.100.11", "192.168.100.12"},
		}

		_, err = C.ModifyNFSExport(context.Background(), &modifyParams, nfsID)
		checkAPIErr(t, err)
	})

	t.Run("delete host 2", func(t *testing.T) {
		modifyParams := gopowerstore.NFSExportModify{
			RemoveRWRootHosts: []string{"192.168.100.11/255.255.255.255"},
		}

		_, err = C.ModifyNFSExport(context.Background(), &modifyParams, nfsID)
		checkAPIErr(t, err)
	})

	deleteNFSExport(t, nfsID)
	deleteFS(t, fsID)
	deleteNAS(t, nasID)
}
