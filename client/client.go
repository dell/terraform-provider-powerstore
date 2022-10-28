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

// NewClient returns the gopowermax client
func NewClient(endpoint string, username string, password string, insecure bool) (*Client, error) {

	clientOptions := pstore.NewClientOptions()
	clientOptions.SetInsecure(insecure)

	pstoreClient, err := newClientWithArgs(endpoint, username, password, clientOptions)
	if err != nil {
		return nil, err
	}

	var client = Client{
		PStoreClient: pstoreClient.(*pstore.ClientIMPL),
	}
	return &client, nil
}

