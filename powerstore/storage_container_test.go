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

func TestGetStorageContainer(t *testing.T) {
	reqType := http.MethodGet
	storageContainerID := "5fb0fcb9-424a-4559-96a1-34896a08e61c"
	url := fmt.Sprintf("%s/api/rest/storage_container?select=*/%s", host, storageContainerID)
	mockedResponse := "{\"id\":\"5fb0fcb9-424a-4559-96a1-34896a08e61c\",\"name\":\"Test-SC\",\"quota\":1099511627776}"

	c, _ := NewClient(&host, &username, &password)
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return([]byte(mockedResponse), nil)

	result, _ := c.GetStorageContainer(m, host, storageContainerID)

	assert.Equal(t, storageContainerID, result.ID)
	assert.Equal(t, "Test-SC", result.Name)
}

func TestGetStorageContainerError(t *testing.T) {
	reqType := http.MethodGet
	storageContainerID := "5fb0fcb9-424a-4559-96a1-34896a08e61c"
	url := fmt.Sprintf("%s/api/rest/storage_container?select=*/%s", host, storageContainerID)
	mockedResponse := "{[{\"id\": \"5fb0fcb9-424a-4559-96a1-34896a08e61c\",\"name\": \"Test-SC\",\"notthere\":\"mockedResponse\"}]}"

	c, _ := NewClient(&host, &username, &password)
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return([]byte(mockedResponse), nil)

	result, err := c.GetStorageContainer(m, host, storageContainerID)

	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCreateStorageContainer(t *testing.T) {
	reqType := http.MethodPost
	storageContainerID := "5fb0fcb9-424a-4559-96a1-34896a08e61c"
	url := fmt.Sprintf("%s/api/rest/storage_container", host)
	storageContainer := StorageContainer{Name: "test-SC"}
	requestBody := "{\"name\":\"test-SC\"}"
	mockedResponse := "{\"id\": \"5fb0fcb9-424a-4559-96a1-34896a08e61c\"}"

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return([]byte(mockedResponse), nil)

	result, err := c.CreateStorageContainer(m, host, storageContainer)

	log.Println("result ", result)
	log.Println("error", err)
	assert.Equal(t, storageContainerID, result)

}

func TestUpdateStorageContainer(t *testing.T) {
	reqType := http.MethodPatch
	storageContainerID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/storage_container/%s", host, storageContainerID)
	storageContainer := StorageContainer{Name: "test_sc"}
	requestBody := "{\"name\":\"test_sc\"}"

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return(nil, nil)

	err := c.UpdateStorageContainer(m, host, storageContainer, storageContainerID)

	log.Println("error", err)
	assert.Nil(t, err)

}

func TestDeleteStorageContainer(t *testing.T) {
	reqType := http.MethodDelete
	storageContainerID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/storage_container/%s", host, storageContainerID)

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return(nil, nil)

	err := c.DeleteStorageContainer(m, host, storageContainerID)

	log.Println("error", err)
	assert.Nil(t, err)

}

func TestUpdateStorageContainerError(t *testing.T) {
	reqType := http.MethodPatch
	storageContainerID := "f67a1292-ef89-49f1-826c-fbfde97d12c6"
	storageContainer := StorageContainer{Name: "test_sc"}
	url := fmt.Sprintf("%s/api/rest/storage_container/%s", host, storageContainerID)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE0A0A0030001\",\"severity\":\"Error\",\"message_l10n\":\"Storage container with ID f67a1292-ef89-49f1-826c-fbfde97d12c6 not found.\",\"arguments\":[\"f67a1292-ef89-49f1-826c-fbfde97d12c6\"]}]}"
	requestBody := "{\"name\":\"test_sc\"}"
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

	err := c.UpdateStorageContainer(m, host, storageContainer, storageContainerID)
	assert.NotNil(t, err)
}

func TestCreateStorageContainerError(t *testing.T) {
	reqType := http.MethodPost
	storageContainer := StorageContainer{Name: "test1"}
	requestBody := "{\"name\":\"test1\"}"
	url := fmt.Sprintf("%s/api/rest/storage_container", host)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE02020010004\",\"severity\":\"Error\",\"message_l10n\":\"The policy name test1 is in use. It needs to be unique regardless of character cases.\",\"arguments\":[\"test1\"]}]"
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

	result, err := c.CreateStorageContainer(m, host, storageContainer)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "status: 400"))
	assert.Equal(t, "", result)
}
func TestDeleteStorageContainerError(t *testing.T) {
	reqType := http.MethodDelete
	storageContainerID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/storage_container/%s", host, storageContainerID)
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

	err := c.DeleteStorageContainer(m, host, storageContainerID)
	assert.NotNil(t, err)
}
