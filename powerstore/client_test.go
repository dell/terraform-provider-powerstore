package powerstore

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDoRequest(t *testing.T) {
	assert := assert.New(t)
	volumeID := "1234"
	mockedResponse := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\",\"name\": \"TestLightFoot1\"}"
	expected := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\",\"name\": \"TestLightFoot1\"}"

	c := &Client{
		HTTPClient: http.DefaultClient,
	}

	//activate http mock
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	//simulate response
	url := fmt.Sprintf("%s/api/rest/volume?select=*/%s", c.HostURL, volumeID)

	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusOK, mockedResponse)
			return resp, nil
		},
	)
	body, _ := c.doRequest(http.MethodGet, url, "")

	assert.Equal(expected, string(body), "Response body")

}

func TestDoRequestErrorForHttpGet(t *testing.T) {
	assert := assert.New(t)
	volumeID := "1234"
	mockedResponse := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\",\"name\": \"TestLightFoot1\"}"

	c := &Client{
		HTTPClient: http.DefaultClient,
	}

	//activate http mock
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	//simulate response
	url := fmt.Sprintf("%s/api/rest/volume?select=*/%s", c.HostURL, volumeID)

	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusBadRequest, mockedResponse)
			err := fmt.Errorf("status: %d, body: %s", http.StatusBadRequest, mockedResponse)
			return resp, err
		},
	)
	body, err := c.doRequest(http.MethodGet, url, "")

	assert.Nil(body, "Response body")
	assert.NotNil(err)
	assert.True(strings.Contains(err.Error(), "status: 400"))
}

func TestDoRequestWithBody(t *testing.T) {
	reqType := http.MethodPost
	assert := assert.New(t)
	mockedResponse := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\"}"
	expected := "{\"id\": \"a797eb5c-1f1f-404f-88e0-7caf8c105174\"}"
	requestBody := "{\"name\":\"testVol\"}"
	c := &Client{
		HTTPClient: http.DefaultClient,
	}

	//activate http mock
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	//simulate response
	url := fmt.Sprintf("%s/api/rest/volume", c.HostURL)

	httpmock.RegisterResponder(reqType, url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusCreated, mockedResponse)
			return resp, nil
		},
	)
	body, _ := c.doRequest(reqType, url, requestBody)

	assert.Equal(expected, string(body), "Response body")
}

func TestDoRequestErrorForHttpPatch(t *testing.T) {
	reqType := http.MethodPatch
	assert := assert.New(t)
	volumeID := "1234"
	mockedResponse := "{\"messages\":[{\"code\":\"0xE0A05001001A\",\"severity\":\"Error\",\"message_l10n\":\"Modify volume test4 (id: f67a1292-ef89-49f1-826c-fbfde97d12c6) with size 12884901888 smaller than or equal to current size 12884901888 is not supported.\",\"arguments\":[\"test4\",\"f67a1292-ef89-49f1-826c-fbfde97d12c6\",\"12884901888\",\"12884901888\"]}]}"
	requestBody := "{\"name\":\"testVol\"}"

	c := &Client{
		HTTPClient: http.DefaultClient,
	}

	//activate http mock
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	//simulate response
	url := fmt.Sprintf("%s/api/rest/volume/%s", c.HostURL, volumeID)

	httpmock.RegisterResponder(reqType, url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusUnprocessableEntity, mockedResponse)
			err := fmt.Errorf("status: %d, body: %s", http.StatusUnprocessableEntity, mockedResponse)
			return resp, err
		},
	)
	body, err := c.doRequest(reqType, url, requestBody)

	assert.Nil(body, "Response body")
	assert.NotNil(err)
	assert.True(strings.Contains(err.Error(), "status: 422"))
}

func TestDoRequestErrorForHttpPost(t *testing.T) {
	reqType := http.MethodPost
	assert := assert.New(t)
	mockedResponse := "{\"messages\":[{\"code\":\"0xE0A05001001A\",\"severity\":\"Error\",\"message_l10n\":\"Modify volume test4 (id: f67a1292-ef89-49f1-826c-fbfde97d12c6) with size 12884901888 smaller than or equal to current size 12884901888 is not supported.\",\"arguments\":[\"test4\",\"f67a1292-ef89-49f1-826c-fbfde97d12c6\",\"12884901888\",\"12884901888\"]}]}"
	requestBody := "{\"name\":\"testVol\"}"

	c := &Client{
		HTTPClient: http.DefaultClient,
	}

	//activate http mock
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	//simulate response
	url := fmt.Sprintf("%s/api/rest/volume", c.HostURL)

	httpmock.RegisterResponder(reqType, url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusBadRequest, mockedResponse)
			err := fmt.Errorf("status: %d, body: %s", http.StatusBadRequest, mockedResponse)
			return resp, err
		},
	)
	body, err := c.doRequest(reqType, url, requestBody)

	assert.Nil(body, "Response body")
	assert.NotNil(err)
	assert.True(strings.Contains(err.Error(), "status: 400"))
}
