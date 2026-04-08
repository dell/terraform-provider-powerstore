/*
 *
 * Copyright © 2022 Dell Inc. or its subsidiaries. All Rights Reserved.
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

const TestSnapshotRulePrefix = "test_snapshotrule_"

func createSnapshotRule(t *testing.T) (string, string) {
	snapshotRuleName := TestSnapshotRulePrefix + randString(8)
	createParams := gopowerstore.SnapshotRuleCreate{
		Name:             snapshotRuleName,
		DesiredRetention: 8,
		Interval:         gopowerstore.SnapshotRuleIntervalEnumFourHours,
	}
	createResp, err := C.CreateSnapshotRule(context.Background(), &createParams)
	checkAPIErr(t, err)
	return createResp.ID, snapshotRuleName
}

func deleteSnapshotRule(t *testing.T, id string) {
	_, err := C.DeleteSnapshotRule(
		context.Background(),
		&gopowerstore.SnapshotRuleDelete{
			DeleteSnaps: true,
		},
		id,
	)
	checkAPIErr(t, err)
}

func TestGetSnapshotRules(t *testing.T) {
	_, err := C.GetSnapshotRules(context.Background())
	checkAPIErr(t, err)
}

func TestGetSnapshotRule(t *testing.T) {
	snapshotRuleID, snapshotRuleName := createSnapshotRule(t)
	defer deleteSnapshotRule(t, snapshotRuleID)

	got, err := C.GetSnapshotRule(context.Background(), snapshotRuleID)
	checkAPIErr(t, err)

	assert.Equal(t, snapshotRuleID, got.ID)
	assert.Equal(t, snapshotRuleName, got.Name)
}

func TestModifySnapshotRule(t *testing.T) {
	snapshotRuleID, _ := createSnapshotRule(t)
	defer deleteSnapshotRule(t, snapshotRuleID)

	_, err := C.ModifySnapshotRule(context.Background(), &gopowerstore.SnapshotRuleCreate{DesiredRetention: 7}, snapshotRuleID)
	checkAPIErr(t, err)
	got, err := C.GetSnapshotRule(context.Background(), snapshotRuleID)
	checkAPIErr(t, err)
	assert.Equal(t, int32(7), got.DesiredRetention)
}
