/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"terraform-provider-powerstore/clientgen"
	"time"

	pstore "github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Client type is to hold powerstore client
type Client struct {
	PStoreClient *pstore.ClientIMPL
	GenClient    *clientgen.APIClient
}

var (
	newClientWithArgs = pstore.NewClientWithArgs
)

// NewClient returns the gopowerstore client
func NewClient(endpoint string, username string, password string, insecure bool, timeout int64) (*Client, error) {

	clientOptions := pstore.NewClientOptions()
	clientOptions.SetInsecure(insecure)
	if timeout == 0 {
		timeout = int64(clientOptions.DefaultTimeout())
	}
	clientOptions.SetDefaultTimeout(int64(timeout))

	genClient, err := newClientGen(context.Background(), endpoint, username, password, insecure, timeout)
	if err != nil {
		return nil, fmt.Errorf("cannot create generated client: %w", err)
	}

	pstoreClient, err := newClientWithArgs(endpoint, username, password, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("cannot create gopowerstore client: %w", err)
	}

	var client = Client{
		PStoreClient: pstoreClient.(*pstore.ClientIMPL),
		GenClient:    genClient,
	}
	return &client, nil
}

// NewClient returns the gopowerstore client
func newClientGen(ctx context.Context, endpoint string, username string, password string, insecure bool, timeout int64) (*clientgen.APIClient, error) {

	// Setup a User-Agent for your API client (replace the provider name for yours):
	userAgent := "terraform-powerstore-provider/1.0.0"
	jar, err := cookiejar.New(nil)
	if err != nil {
		tflog.Error(ctx, "Got error while creating cookie jar")
	}

	httpclient := &http.Client{
		Timeout: (time.Duration(60) * time.Second),
		Jar:     jar,
	}
	if insecure {
		/* #nosec */
		httpclient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true,
			},
		}
	} else {
		// Loading system certs by default if insecure is set to false
		pool, err := x509.SystemCertPool()
		if err != nil {
			errSysCerts := errors.New("unable to initialize cert pool from system")
			return nil, errSysCerts
		}
		httpclient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				RootCAs:            pool,
				InsecureSkipVerify: false,
			},
		}
	}

	url := endpoint
	basicAuthString := basicAuth(username, password)

	cfg := &clientgen.Configuration{
		HTTPClient:    httpclient,
		DefaultHeader: make(map[string]string),
		UserAgent:     userAgent,
		Debug:         true,
		Servers: clientgen.ServerConfigurations{
			{
				URL:         url,
				Description: url,
			},
		},
		OperationServers: map[string]clientgen.ServerConfigurations{},
	}
	cfg.AddDefaultHeader("Authorization", "Basic "+basicAuthString)

	apiClient := clientgen.NewAPIClient(cfg)

	_, resp, err := apiClient.LoginSessionApi.GetAllLoginSessions(ctx).Execute()
	if err != nil {
		return nil, err
	}

	// get the DELL-EMC-TOKEN header from the response
	token := resp.Header.Get("DELL-EMC-TOKEN")
	if len(token) != 0 {
		cfg.AddDefaultHeader("DELL-EMC-TOKEN", token)
		apiClient = clientgen.NewAPIClient(cfg)
	} else {
		return nil, errors.New("no token returned during login")
	}

	return apiClient, nil

}

// Generate the base 64 Authorization string from username / password.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
