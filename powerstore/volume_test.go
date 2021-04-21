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

func TestGetVolume(t *testing.T) {
	reqType := http.MethodGet
	volumeID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/volume?select=*/%s", host, volumeID)
	mockedResponse := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\",\"name\": \"TestLightFoot1\"}"

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return([]byte(mockedResponse), nil)

	result, err := c.GetVolume(m, host, volumeID)

	log.Println("result ", result)
	log.Println("error", err)
	assert.Equal(t, volumeID, result.ID)
	assert.Equal(t, "TestLightFoot1", result.Name)

}

func TestCreateVolume(t *testing.T) {
	reqType := http.MethodPost
	volumeID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/volume", host)
	vol := VolRequest{Name: "testVol"}
	requestBody := "{\"name\":\"testVol\"}"
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

	result, err := c.CreateVolume(m, host, vol)

	log.Println("result ", result)
	log.Println("error", err)
	assert.Equal(t, volumeID, result)

}

func TestUpdateVolume(t *testing.T) {
	reqType := http.MethodPatch
	volumeID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/volume/%s", host, volumeID)
	vol := VolRequest{Name: "testVol"}
	requestBody := "{\"name\":\"testVol\"}"

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, requestBody).
		Return(nil, nil)

	err := c.UpdateVolume(m, host, vol, volumeID)

	log.Println("error", err)
	assert.Nil(t, err)

}
func TestDeleteVolume(t *testing.T) {
	reqType := http.MethodDelete
	volumeID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/volume/%s", host, volumeID)

	c, _ := NewClient(&host, &username, &password)

	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockClientInterface(ctrl)

	m.
		EXPECT().
		doRequest(reqType, url, "").
		Return(nil, nil)

	err := c.DeleteVolume(m, host, volumeID)

	log.Println("error", err)
	assert.Nil(t, err)

}

func TestUpdateVolumeError(t *testing.T) {
	reqType := http.MethodPatch
	volumeID := "f67a1292-ef89-49f1-826c-fbfde97d12c6"
	vol := VolRequest{Name: "testVol"}
	url := fmt.Sprintf("%s/api/rest/volume/%s", host, volumeID)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE0A05001001A\",\"severity\":\"Error\",\"message_l10n\":\"Modify volume test4 (id: f67a1292-ef89-49f1-826c-fbfde97d12c6) with size 12884901888 smaller than or equal to current size 12884901888 is not supported.\",\"arguments\":[\"test4\",\"f67a1292-ef89-49f1-826c-fbfde97d12c6\",\"12884901888\",\"12884901888\"]}]}"
	requestBody := "{\"name\":\"testVol\"}"
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

	err := c.UpdateVolume(m, host, vol, volumeID)
	assert.NotNil(t, err)
}

func TestCreateVolumeError(t *testing.T) {
	reqType := http.MethodPost
	vol := VolRequest{Name: "test1"}
	requestBody := "{\"name\":\"test1\"}"
	url := fmt.Sprintf("%s/api/rest/volume", host)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE0A080010014\",\"severity\":\"Error\",\"message_l10n\":\"The volume name is already in use: test4\",\"arguments\":[\"test4\"]}]}"
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

	result, err := c.CreateVolume(m, host, vol)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "status: 400"))
	assert.Equal(t, "", result)
}

func TestDeleteVolumeError(t *testing.T) {
	reqType := http.MethodDelete
	volumeID := "a797eb5c-1f1f-404f-88e0-7caf8c105174"
	url := fmt.Sprintf("%s/api/rest/volume/%s", host, volumeID)

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

	err := c.DeleteVolume(m, host, volumeID)
	assert.NotNil(t, err)
}
