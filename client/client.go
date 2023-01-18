package client

import (
	pstore "github.com/dell/gopowerstore"
)

// Client type is to hold powerstore client
type Client struct {
	PStoreClient *pstore.ClientIMPL
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
	clientOptions.SetDefaultTimeout(uint64(timeout))

	pstoreClient, err := newClientWithArgs(endpoint, username, password, clientOptions)
	if err != nil {
		return nil, err
	}

	var client = Client{
		PStoreClient: pstoreClient.(*pstore.ClientIMPL),
	}
	return &client, nil
}
