package powerstore

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProtectionPolicy(t *testing.T) {
	reqType := http.MethodGet
	protectionPolicyID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/policy/%s?select=name,id,description,type,file_systems(id,name),replication_rules(id,name),snapshot_rules(id,name,interval,time_of_day,days_of_week),virtual_machines(id,name),volume(id,name),volume_group(id,name)&type=eq.Protection", host, protectionPolicyID)
	mockedResponse := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\",\"name\": \"test\"}"
	c, _ := NewClient(&host, &username, &password)
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return([]byte(mockedResponse), nil)

	result, _ := c.GetProtectionPolicy(m, host, protectionPolicyID)

	assert.Equal(t, protectionPolicyID, result.ID)
	assert.Equal(t, "test", result.Name)
}

func TestCreateProtectionPolicy(t *testing.T) {
	reqType := http.MethodPost
	policyID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/policy", host)
	policy := ProtectionPolicyRequest{Name: "testPolicy"}
	requestBody := "{\"name\":\"testPolicy\"}"
	mockedResponse := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\"}"

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return([]byte(mockedResponse), nil)

	result, err := c.CreateProtectionPolicy(m, host, policy)

	log.Println("result ", result)
	log.Println("error", err)
	assert.Equal(t, policyID, result)

}

func TestCreateProtectionPolicyError(t *testing.T) {
	reqType := http.MethodPost
	protectionPolicy := ProtectionPolicyRequest{Name: "test1"}
	requestBody := "{\"name\":\"test1\"}"
	url := fmt.Sprintf("%s/api/rest/policy", host)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE02020010004\",\"severity\":\"Error\",\"message_l10n\":\"The policy name test1 is in use. It needs to be unique regardless of character cases.\",\"arguments\":[\"test1\"]}]}"
	expectedErr := fmt.Errorf("status: %d, body: %s", http.StatusBadRequest, mockedResponse)

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return([]byte(mockedResponse), expectedErr)

	result, err := c.CreateProtectionPolicy(m, host, protectionPolicy)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "status: 400"))
	assert.Equal(t, "", result)
}

func TestGetProtectionPolicyError(t *testing.T) {
	reqType := http.MethodGet
	protectionPolicyID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/policy/%s?select=name,id,description,type,file_systems(id,name),replication_rules(id,name),snapshot_rules(id,name,interval,time_of_day,days_of_week),virtual_machines(id,name),volume(id,name),volume_group(id,name)&type=eq.Protection", host, protectionPolicyID)
	errorResponse := fmt.Errorf("{\"messages\":[{\"code\":\"0xE04040020009\",\"severity\":\"Error\",\"message_l10n\":\"Instance with id 3fdcagb7-361f-4434-8974-02341540c78d was not found.\",\"arguments\":[\"3fdcagb7-361f-4434-8974-02341540c78d\"]}]}")

	c, _ := NewClient(&host, &username, &password)
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return(nil, errorResponse)

	result, err := c.GetProtectionPolicy(m, host, protectionPolicyID)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestUpdateProtectionPolicy(t *testing.T) {
	reqType := http.MethodPatch
	protectionPolicyID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/policy/%s", host, protectionPolicyID)
	requestBody := "{\"description\":\"updated description\"}"
	protectionPolicy := ProtectionPolicyRequest{
		Description: "updated description",
	}

	c, _ := NewClient(&host, &username, &password)
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return(nil, nil)

	err := c.UpdateProtectionPolicy(m, host, protectionPolicy, protectionPolicyID)

	assert.Nil(t, err)
}

func TestUpdateProtectionPolicyError(t *testing.T) {
	reqType := http.MethodPatch
	protectionPolicyID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/policy/%s", host, protectionPolicyID)
	requestBody := "{\"description\":\"updated description\"}"
	protectionPolicy := ProtectionPolicyRequest{
		Description: "updated description",
	}
	mockedResponse := "{\"messages\": [{\"code\": \"0xE0A090010001\", \"severity\": \"Error\", \"message_l10n\": \"Unable to find the policy with ID 63fd0fad-ad68-4bd7-9ad6-dcee3b61b8a2\", \"arguments\": [\"63fd0fad-ad68-4bd7-9ad6-dcee3b61b8a2\"]}]}"
	expectedErr := fmt.Errorf("status: %d, body: %s", http.StatusBadRequest, mockedResponse)

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return([]byte(mockedResponse), expectedErr)

	err := c.UpdateProtectionPolicy(m, host, protectionPolicy, protectionPolicyID)
	assert.NotNil(t, err)
}

func TestDeleteProtectionPolicy(t *testing.T) {
	reqType := http.MethodDelete
	protectionPolicyID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/policy/%s", host, protectionPolicyID)
	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return(nil, nil)

	err := c.DeleteProtectionPolicy(m, host, protectionPolicyID)

	assert.Nil(t, err)

}

func TestDeleteProtectionPolicyError(t *testing.T) {
	reqType := http.MethodDelete
	protectionPolicyID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/policy/%s", host, protectionPolicyID)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE04040020009\",\"severity\":\"Error\",\"message_l10n\":\"Instance with id 8b29e8e9-d6d4-4ada-99d9-0ace219f6a17 was not found.\",\"arguments\":[\"8b29e8e9-d6d4-4ada-99d9-0ace219f6a17\"]}]}"
	expectedErr := fmt.Errorf("status: %d, body: %s", http.StatusNotFound, mockedResponse)

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return([]byte(mockedResponse), expectedErr)

	err := c.DeleteProtectionPolicy(m, host, protectionPolicyID)
	assert.NotNil(t, err)
}
