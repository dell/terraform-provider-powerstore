/*
 *
 * Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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

const TestStorageContainerPrefix = "test_sc_"

func createSC(t *testing.T) (string, string) {
	scName := TestStorageContainerPrefix + randString(8)
	createParams := gopowerstore.StorageContainer{}
	createParams.Name = scName
	createResp, err := C.CreateStorageContainer(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, scName
}

func deleteSC(t *testing.T, id string) {
	_, err := C.DeleteStorageContainer(context.Background(), id)
	checkAPIErr(t, err)
}

func TestGetSC(t *testing.T) {
	scID, scName := createSC(t)
	StorageContainer, err := C.GetStorageContainer(context.Background(), scID)
	checkAPIErr(t, err)
	assert.NotEmpty(t, StorageContainer.Name)
	assert.Equal(t, scName, StorageContainer.Name)
	deleteSC(t, scID)
}

func TestCreateDeleteSC(t *testing.T) {
	scID, _ := createSC(t)
	deleteSC(t, scID)
}

func TestModifySC(t *testing.T) {
	scID, _ := createSC(t)
	defer deleteSC(t, scID)

	newName := TestStorageContainerPrefix + randString(8) + "new"

	_, err := C.ModifyStorageContainer(context.Background(), &gopowerstore.StorageContainer{Name: newName}, scID)
	checkAPIErr(t, err)
	got, err := C.GetStorageContainer(context.Background(), scID)
	checkAPIErr(t, err)
	assert.Equal(t, newName, got.Name)
}
