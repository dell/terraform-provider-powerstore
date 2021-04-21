package powerstore

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//ClientInterface - interface of all client functions
//go:generate mockgen -source=powerstore/client.go -destination=powerstore/mock_client.go -package=powerstore ClientInterface
type ClientInterface interface {
	doRequest(reqType string, url string, requestBody string) ([]byte, error)
}

// NewClient -
func NewClient(host, username, password *string) (*Client, error) {
	if (host != nil) && (username != nil) && (password != nil) {
		c := Client{
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true, //nolint ...we know its not secure
					},
				},
			},
			HostURL:  *host,
			Username: *username,
			Password: *password,
		}

		return &c, nil
	}
	err := fmt.Errorf("Host, username or password is empty. Making new client failed")

	return nil, err
}

func (c *Client) doRequest(reqType string, url string, requestBody string) ([]byte, error) {
	var reqBody io.Reader
	if requestBody == "" {
		reqBody = nil
	} else {
		reqBody = strings.NewReader(requestBody)
	}

	req, err := http.NewRequest(reqType, url, reqBody)
	log.Println("url==", url)
	//adding authentication header
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	//api request to the appliance
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if (reqType == http.MethodGet && res.StatusCode != http.StatusOK) || (reqType == http.MethodPost && res.StatusCode != http.StatusCreated) || ((reqType == http.MethodPatch || reqType == http.MethodDelete) && res.StatusCode != http.StatusNoContent) {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
