package powerstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//StorageContainerInterface - interface of all client functions
type StorageContainerInterface interface {
	GetStorageContainer(ClientInterface, string, string) (*StorageContainer, error)
	CreateStorageContainer(ClientInterface, string, StorageContainer) (string, error)
	UpdateStorageContainer(ClientInterface, string, StorageContainer, string) error
	DeleteStorageContainer(ClientInterface, string, string) error
}

// GetStorageContainer - Returns a specifc storage container
func (c *Client) GetStorageContainer(clientInterface ClientInterface, hostURL string, storageContainerID string) (*StorageContainer, error) {

	reqType := http.MethodGet

	// URL for storage container resource type
	url := fmt.Sprintf("%s/api/rest/storage_container?select=*/%s", hostURL, storageContainerID)

	body, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return nil, err
	}
	log.Println("body ===", string(body))
	storageContainer := StorageContainer{}
	err = json.Unmarshal(body, &storageContainer)
	if err != nil {
		return nil, err
	}
	return &storageContainer, nil
}

// CreateVolume - Create a volume and returns volume id
func (c *Client) CreateStorageContainer(clientInterface ClientInterface, hostURL string, storageContainer StorageContainer) (string, error) {

	reqType := http.MethodPost

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/storage_container", hostURL)

	requestBody, err := json.Marshal(storageContainer)
	if err != nil {
		return "", err
	}

	body, err := clientInterface.doRequest(reqType, url, string(requestBody))
	if err != nil {
		return "", err
	}
	log.Println("body ===", string(body))
	storageContainerResponse := StorageContainer{}
	err = json.Unmarshal(body, &storageContainerResponse)
	if err != nil {
		return "", err
	}

	return storageContainerResponse.ID, nil
}

// UpdateVolume - Update a volume and returns volume id
func (c *Client) UpdateStorageContainer(clientInterface ClientInterface, hostURL string, storageContainer StorageContainer, storageContainerID string) error {
	reqType := http.MethodPatch

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/storage_container/%s", hostURL, storageContainerID)
	log.Println("update url ==", url)
	log.Println("ID = ", storageContainerID)
	requestBody, err := json.Marshal(storageContainer)
	if err != nil {
		log.Println("Unmarshalling Error: ", err.Error())
		return err
	}
	log.Println("request body == ", string(requestBody))
	_, err = clientInterface.doRequest(reqType, url, string(requestBody))
	if err != nil {
		log.Println("response Error: ", err.Error())
		return err
	}
	return nil
}

//DeleteVolume - Delete a volume based on volume id
func (c *Client) DeleteStorageContainer(clientInterface ClientInterface, hostURL string, storageContainerID string) error {

	reqType := http.MethodDelete

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/storage_container/%s", hostURL, storageContainerID)

	_, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return err
	}
	return nil
}

