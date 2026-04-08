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
	"github.com/stretchr/testify/suite"
)

const (
	TestFSPrefix  = "test_fs_"
	TestNASPrefix = "test_nas_"
)

const DefaultFSSize int64 = 3221225472

type FsTestSuite struct {
	suite.Suite
	nasID   string
	nasName string
}

func (suite *FsTestSuite) SetupSuite() {
	suite.nasID, suite.nasName = createNAS(suite.T())
}

func (suite *FsTestSuite) TearDownSuite() {
	deleteNAS(suite.T(), suite.nasID)
}

func createNAS(t *testing.T) (string, string) {
	nasName := TestNASPrefix + randString(8)
	createParams := gopowerstore.NASCreate{
		Name: nasName,
	}
	createResp, err := C.CreateNAS(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, nasName
}

func deleteNAS(t *testing.T, id string) {
	_, err := C.DeleteNAS(context.Background(), id)
	checkAPIErr(t, err)
}

func createFS(t *testing.T, nasID string) (string, string) {
	nasName := TestFSPrefix + randString(8)
	createParams := gopowerstore.FsCreate{
		Name:        nasName,
		NASServerID: nasID,
		Size:        DefaultFSSize,
	}
	createResp, err := C.CreateFS(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, nasName
}

func deleteFS(t *testing.T, id string) {
	_, err := C.DeleteFS(context.Background(), id)
	checkAPIErr(t, err)
}

func createFsSnap(t *testing.T, fsID string, fsName string) gopowerstore.CreateResponse {
	snapName := fsName + "_snapshot"
	snapDesc := "just a description"
	snap, snapCreateErr := C.CreateFsSnapshot(context.Background(), &gopowerstore.SnapshotFSCreate{
		Name:        snapName,
		Description: snapDesc,
	}, fsID)
	checkAPIErr(t, snapCreateErr)
	return snap
}

func (suite *FsTestSuite) TestCreateFSSnapshot() {
	t := suite.T()
	fsID, fsName := createFS(t, suite.nasID)
	defer deleteFS(t, fsID)

	snap := createFsSnap(t, fsID, fsName)
	defer C.DeleteFsSnapshot(context.Background(), snap.ID)
}

func (suite *FsTestSuite) TestCreateFsFromSnapshot() {
	t := suite.T()
	fsID, fsName := createFS(t, suite.nasID)
	defer deleteFS(t, fsID)

	snap := createFsSnap(t, fsID, fsName)
	defer C.DeleteFsSnapshot(context.Background(), snap.ID)

	name := "new_fs_from_snap" + randString(8)
	createParams := gopowerstore.FsClone{}
	createParams.Name = &name
	snapVol, err := C.CreateFsFromSnapshot(context.Background(), &createParams, snap.ID)
	checkAPIErr(t, err)
	assert.NotEmpty(t, snapVol.ID)
	deleteFS(t, snapVol.ID)
}

func (suite *FsTestSuite) TestGetFsSnapshot() {
	t := suite.T()
	fsID, fsName := createFS(t, suite.nasID)
	defer deleteFS(t, fsID)

	snap := createFsSnap(t, fsID, fsName)
	assert.NotEmpty(t, snap.ID)
	defer C.DeleteFsSnapshot(context.Background(), snap.ID)

	got, err := C.GetFsSnapshot(context.Background(), snap.ID)
	checkAPIErr(t, err)

	assert.Equal(t, snap.ID, got.ID)
}

func (suite *FsTestSuite) TestGetFsSnapshots() {
	t := suite.T()
	_, err := C.GetFsSnapshots(context.Background())
	checkAPIErr(t, err)
}

func (suite *FsTestSuite) TestGetNonExistingFsSnapshot() {
	t := suite.T()
	fsID, fsName := createFS(t, suite.nasID)
	defer deleteFS(t, fsID)

	snap := createFsSnap(t, fsID, fsName)
	assert.NotEmpty(t, snap.ID)

	_, err := C.DeleteFsSnapshot(context.Background(), snap.ID)
	assert.NotEmpty(t, snap.ID)

	got, err := C.GetSnapshot(context.Background(), snap.ID)
	assert.Error(t, err)
	assert.Empty(t, got)
}

func (suite *FsTestSuite) TestGetNASByName() {
	t := suite.T()
	nas, err := C.GetNASByName(context.Background(), suite.nasName)
	checkAPIErr(t, err)
	assert.NotEmpty(t, nas.Name)
	assert.Equal(t, suite.nasName, nas.Name)
}

func (suite *FsTestSuite) TestGetNASServers() {
	t := suite.T()
	_, err := C.GetNASServers(context.Background())
	checkAPIErr(t, err)
}

func (suite *FsTestSuite) TestGetFSByName() {
	t := suite.T()
	fsID, fsName := createFS(t, suite.nasID)
	defer deleteFS(t, fsID)
	fs, err := C.GetFSByName(context.Background(), fsName)
	checkAPIErr(t, err)
	assert.NotEmpty(t, fs.Name)
	assert.Equal(t, fsName, fs.Name)
}

func (suite *FsTestSuite) TestFsSnapshotAlreadyExist() {
	t := suite.T()
	fsID, fsName := createFS(t, suite.nasID)
	defer deleteFS(t, fsID)
	snap := createFsSnap(t, fsID, fsName)
	assert.NotEmpty(t, snap.ID)

	snapName := fsName + "_snapshot"
	snapDesc := "just a description"
	_, err := C.CreateFsSnapshot(context.Background(), &gopowerstore.SnapshotFSCreate{
		Name:        snapName,
		Description: snapDesc,
	}, fsID)
	assert.NotNil(t, err)
	apiError := err.(gopowerstore.APIError)
	assert.True(t, apiError.FSNameIsAlreadyUse())
	_, err = C.DeleteFsSnapshot(context.Background(), snap.ID)
	assert.NotEmpty(t, snap.ID)
}

func (suite *FsTestSuite) TestModifyFS() {
	t := suite.T()
	fsID, _ := createFS(t, suite.nasID)
	defer deleteFS(t, fsID)

	_, err := C.ModifyFS(context.Background(), &gopowerstore.FSModify{
		Size:        3221225472 * 2,
		Description: "New Description",
	}, fsID)
	checkAPIErr(t, err)

	gotVol, err := C.GetFS(context.Background(), fsID)
	assert.Equal(t, int64(3221225472*2), gotVol.SizeTotal)
	assert.Equal(t, "New Description", gotVol.Description)
}

func TestFsTestSuite(t *testing.T) {
	suite.Run(t, new(FsTestSuite))
}
