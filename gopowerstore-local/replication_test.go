/*
 *
 * Copyright © 2021-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
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
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	policyMockURL             = APIMockURL + policyURL
	replicationRuleMockURL    = APIMockURL + replicationRuleURL
	replicationSessionMockURL = APIMockURL + replicationSessionURL
)

var (
	protectionPolicyID  = "15c03067-c4f2-428b-b637-18b0266979f0"
	protectionPolicyID2 = "3224ff5a-2e83-4a7f-a0c4-009df20e36db"
	replicationRuleID   = "6b930711-46bc-4a4b-9d6a-22c77a7838c4"
	replicationRuleID2  = "2d0780e3-2ce7-4d8b-b2ec-349c5e9e26a9"
)

func TestClientIMPL_CreateProtectionPolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", policyMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := ProtectionPolicyCreate{
		Name:               "pp-test",
		Description:        "pp-test",
		ReplicationRuleIDs: []string{"id"},
	}

	resp, err := C.CreateProtectionPolicy(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_CreateReplicationRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", replicationRuleMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := ReplicationRuleCreate{
		Name:           "rr-test",
		Rpo:            RpoFiveMinutes,
		RemoteSystemID: "XX-0000X",
	}

	resp, err := C.CreateReplicationRule(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_CreateReplicationRuleSync(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", replicationRuleMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := ReplicationRuleCreate{
		Name:           "rr-test",
		Rpo:            RpoZero,
		RemoteSystemID: "XX-0000X",
	}

	resp, err := C.CreateReplicationRule(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_ModifyReplicationRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", replicationRuleMockURL, replicationRuleID),
		httpmock.NewStringResponder(204, ""))

	modifyParams := ReplicationRuleModify{
		Name: "rr-test-modified",
		Rpo:  "One_Day",
	}

	resp, err := C.ModifyReplicationRule(context.Background(), &modifyParams, replicationRuleID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_DeleteProtectionPolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", policyMockURL, volID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteProtectionPolicy(context.Background(), volID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_DeleteReplicationRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", replicationRuleMockURL, volID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteReplicationRule(context.Background(), volID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_GetProtectionPolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, protectionPolicyID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", policyMockURL, protectionPolicyID),
		httpmock.NewStringResponder(200, respData))
	protectionPolicy, err := C.GetProtectionPolicy(context.Background(), protectionPolicyID)
	assert.Nil(t, err)
	assert.Equal(t, protectionPolicyID, protectionPolicy.ID)
}

func TestClientIMPL_GetProtectionPolicyByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", policyMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	vol, err := C.GetProtectionPolicyByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, volID, vol.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetProtectionPolicyByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_ModifyProtectionPolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/%s", policyMockURL, protectionPolicyID),
		httpmock.NewStringResponder(201, ""))

	modifyParams := ProtectionPolicyCreate{
		Description: "Test ModifyProtectionPolicy",
	}

	resp, err := C.ModifyProtectionPolicy(context.Background(), &modifyParams, protectionPolicyID)
	assert.Nil(t, err)
	assert.Equal(t, EmptyResponse(""), resp)
}

func TestClientIMPL_GetReplicationRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, replicationRuleID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", replicationRuleMockURL, replicationRuleID),
		httpmock.NewStringResponder(200, respData))
	replicationRule, err := C.GetReplicationRule(context.Background(), replicationRuleID)
	assert.Nil(t, err)
	assert.Equal(t, replicationRuleID, replicationRule.ID)
}

func TestClientIMPL_GetReplicationRuleByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", replicationRuleMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	replicationRule, err := C.GetReplicationRuleByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, volID, replicationRule.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetReplicationRuleByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetReplicationSessionByLocalResourceID(t *testing.T) {
	// test getting a valid replication session
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{
		"id":"%s",
		"name":"%s"
	}]`, volID, "session")
	httpmock.RegisterResponder("GET", replicationSessionMockURL,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetReplicationSessionByLocalResourceID(context.Background(), volID2)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)

	// test when the replication group does not exist
	httpmock.Reset()

	httpmock.RegisterResponder("GET", replicationSessionMockURL,
		httpmock.NewStringResponder(200, ""))
	_, err = C.GetReplicationSessionByLocalResourceID(context.Background(), volID2)

	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetProtectionPolicies(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, protectionPolicyID, protectionPolicyID2)
	httpmock.RegisterResponder("GET", policyMockURL,
		httpmock.NewStringResponder(200, respData))
	policies, err := C.GetProtectionPolicies(context.Background())
	assert.Nil(t, err)
	assert.Len(t, policies, 2)
	assert.Equal(t, protectionPolicyID, policies[0].ID)
}

func TestClientIMPL_GetReplicationRules(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, replicationRuleID, replicationRuleID2)
	httpmock.RegisterResponder("GET", replicationRuleMockURL,
		httpmock.NewStringResponder(200, respData))
	rules, err := C.GetReplicationRules(context.Background())
	assert.Nil(t, err)
	assert.Len(t, rules, 2)
	assert.Equal(t, replicationRuleID, rules[0].ID)
}
