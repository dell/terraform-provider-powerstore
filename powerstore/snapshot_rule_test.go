package powerstore

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetSnapshotRule(t *testing.T) {
	reqType := http.MethodGet
	snapshotRuleID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/snapshot_rule?select=*/%s", host, snapshotRuleID)
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

	result, _ := c.GetSnapshotRule(m, host, snapshotRuleID)

	assert.Equal(t, snapshotRuleID, result.ID)
	assert.Equal(t, "test", result.Name)
}

func TestGetSnapshotRuleError(t *testing.T) {
	reqType := http.MethodGet
	snapshotRuleID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/snapshot_rule?select=*/%s", host, snapshotRuleID)
	mockedResponse := "{[{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\",\"name\": \"test\",\"notthere\":\"mockedResponse\"}]}"

	c, _ := NewClient(&host, &username, &password)
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return([]byte(mockedResponse), nil)

	result, err := c.GetSnapshotRule(m, host, snapshotRuleID)

	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCreateSnapshotRule(t *testing.T) {
	reqType := http.MethodPost
	url := fmt.Sprintf("%s/api/rest/snapshot_rule", host)
	snapshotID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	mockedResponse := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\"}"
	requestBody := "{\"name\":\"test1\"}"

	snapshotRule := SnapshotRule{
		Name: "test1",
	}

	c, _ := NewClient(&host, &username, &password)
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return([]byte(mockedResponse), nil)

	result, err := c.CreateSnapshotRule(m, host, snapshotRule)

	assert.Nil(t, err)
	assert.Equal(t, snapshotID, result)
}

func TestCreateSnapshotRuleError(t *testing.T) {
	reqType := http.MethodPost
	snapshot := SnapshotRule{Name: "test1"}
	requestBody := "{\"name\":\"test1\"}"
	url := fmt.Sprintf("%s/api/rest/snapshot_rule", host)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE02030010015\",\"severity\":\"Error\",\"message_l10n\":\"The rule name test2 is already used by another rule. It needs to be unique (case insensitive). Please use a different name.\",\"arguments\":[\"test2\"]}]}"
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

	result, err := c.CreateSnapshotRule(m, host, snapshot)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "status: 400"))
	assert.Equal(t, "", result)
}

func TestUpdateSnapshotRule(t *testing.T) {
	reqType := http.MethodPatch
	snapshotID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/snapshot_rule/%s", host, snapshotID)
	requestBody := "{\"name\":\"test1\"}"

	snapshotRule := SnapshotRule{
		Name: "test1",
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

	err := c.UpdateSnapshotRule(m, host, snapshotRule, snapshotID)

	assert.Nil(t, err)
}

func TestUpdateSnapshotRuleError(t *testing.T) {
	reqType := http.MethodPatch
	snapshotID := "f67a1292-ef89-49f1-826c-fbfde97d12c6"
	snapshot := SnapshotRule{Name: "test"}
	url := fmt.Sprintf("%s/api/rest/snapshot_rule/%s", host, snapshotID)
	mockedResponse := "{\"messages\": [{\"code\": \"0xE02030010015\",\"severity\": \"Error\",\"message_l10n\": \"The rule name test is already used by another rule. It needs to be unique (case insensitive). Please use a different name.\",\"arguments\": [\"test\"]}]}"
	requestBody := "{\"name\":\"test\"}"
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

	err := c.UpdateSnapshotRule(m, host, snapshot, snapshotID)
	assert.NotNil(t, err)
}

func TestDeleteSnapshotRule(t *testing.T) {
	reqType := http.MethodDelete
	snapshotRuleID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/snapshot_rule/%s", host, snapshotRuleID)

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return(nil, nil)

	err := c.DeleteSnapshotRule(m, host, snapshotRuleID)

	assert.Nil(t, err)

}

func TestDeleteSnapshotRuleError(t *testing.T) {
	reqType := http.MethodDelete
	snapshotRuleID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/snapshot_rule/%s", host, snapshotRuleID)
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

	err := c.DeleteSnapshotRule(m, host, snapshotRuleID)
	assert.NotNil(t, err)
}
