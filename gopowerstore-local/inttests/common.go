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
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/dell/gopowerstore"
)

const (
	TestVolumePrefix       = "test_vol_"
	DefaultVolSize   int64 = 1048576
	DefaultChunkSize int64 = 1048576
	letters                = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func checkAPIErr(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func skipTestOnError(t *testing.T, err error) {
	if err != nil {
		t.Skip("Skipping test..")
	}
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		if len(letters) > 0 {
			n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
			if err == nil {
				b[i] = letters[n.Int64()]
			}
		}
	}
	return string(b)
}

func CreateVol(t *testing.T) (volID, volName string) {
	volName = TestVolumePrefix + randString(8)
	createParams := gopowerstore.VolumeCreate{}
	createParams.Name = &volName
	size := DefaultVolSize
	createParams.Size = &size
	createResp, err := C.CreateVolume(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, volName
}

func DeleteVol(t *testing.T, id string) {
	_, err := C.DeleteVolume(context.Background(), nil, id)
	checkAPIErr(t, err)
}

// check for the target string in the list of strings and return true if found.
func Includes(list *[]string, target string) bool {
	for _, s := range *list {
		if s == target {
			return true
		}
	}
	return false
}

// GetRemoteSystemForMetro queries the source PowerStore array for configured remote systems
// to find a remote system capable of metro replication and returns the system if one is found
// and an empty RemoteSystem struct otherwise.
func GetRemoteSystemForMetro(client gopowerstore.Client, t *testing.T) gopowerstore.RemoteSystem {
	systems, err := client.GetAllRemoteSystems(context.Background())
	if err != nil {
		t.Skip("Could not get remote systems. Skipping test...")
	}

	// try to find a valid remote system with Metro from the list of all available remote systems
	for _, sys := range systems {
		// check remote capabilities for metro and create MetroConfig if found
		if Includes(&sys.Capabilities, string(gopowerstore.BlockMetro)) {
			// make sure the connection is in a good state
			if sys.DataConnectionState == string(gopowerstore.ConnStateOK) {
				return sys
			}
		}
	}

	return gopowerstore.RemoteSystem{}
}
