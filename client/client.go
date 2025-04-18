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
	"github.com/dell/gopowerstore"
	pstore "github.com/dell/gopowerstore"
	"github.com/dell/gopowerstore/api"
)

// Client type is to hold powerstore client
type Client struct {
	PStoreClient *pstore.ClientIMPL
}

var (
	newClientWithArgs = pstore.NewClientWithArgs
)

const (
	paginationDefaultPageSize = 1000
)

// NewClient returns the gopowerstore client
func NewClient(endpoint string, username string, password string, insecure bool, timeout int64) (*Client, error) {

	clientOptions := pstore.NewClientOptions()
	clientOptions.SetInsecure(insecure)
	if timeout == 0 {
		timeout = int64(clientOptions.DefaultTimeout())
	}
	clientOptions.SetDefaultTimeout(int64(timeout))

	pstoreClient, err := newClientWithArgs(endpoint, username, password, clientOptions)
	if err != nil {
		return nil, err
	}

	var client = Client{
		PStoreClient: pstoreClient.(*pstore.ClientIMPL),
	}
	return &client, nil
}

// method allow to read paginated data from backend
func (c *Client) readPaginatedData(f func(int) (api.RespMeta, error)) error {
	var err error
	var meta api.RespMeta
	meta, err = f(0)
	if err != nil {
		return err
	}
	if meta.Pagination.IsPaginate {
		for {
			nextOffset := meta.Pagination.Last + 1
			if nextOffset >= meta.Pagination.Total {
				break
			}
			meta, err = f(nextOffset)
			err = gopowerstore.WrapErr(err)
			if err != nil {
				apiError, ok := err.(*gopowerstore.APIError)
				if !ok {
					return err
				}
				if apiError.BadRange() {
					// could happen if some instances was deleted during pagination
					break
				}
			}
		}
	}
	return nil
}
