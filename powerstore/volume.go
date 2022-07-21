package powerstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//VolumeInterface - interface of all client functions
type VolumeInterface interface {
	GetVolume(clientInterface ClientInterface, hostURL string, volumeID string) (*Volume, error)
	CreateVolume(clientInterface ClientInterface, hostURL string, vol VolRequest) (string, error)
}

// GetVolume - Returns a specific volume details
func (c *Client) GetVolume(clientInterface ClientInterface, hostURL string, volumeID string) (*Volume, error) {

	reqType := http.MethodGet

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/volume?select=*/%s", hostURL, volumeID)

	body, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return nil, err
	}
	log.Println("body ===", string(body))
	volume := Volume{}
	err = json.Unmarshal(body, &volume)
	if err != nil {
		return nil, err
	}
	return &volume, nil
}

// CreateVolume - Create a volume and returns volume id
func (c *Client) CreateVolume(clientInterface ClientInterface, hostURL string, vol VolRequest) (string, error) {

	reqType := http.MethodPost

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/volume", hostURL)

	requestBody, err := json.Marshal(vol)
	if err != nil {
		return "", err
	}

	body, err := clientInterface.doRequest(reqType, url, string(requestBody))
	if err != nil {
		return "", err
	}
	log.Println("body ===", string(body))
	volume := Volume{}
	err = json.Unmarshal(body, &volume)
	if err != nil {
		return "", err
	}

	return volume.ID, nil
}

// UpdateVolume - Update a volume and returns volume id
func (c *Client) UpdateVolume(clientInterface ClientInterface, hostURL string, vol VolRequest, volumeID string) error {
	reqType := http.MethodPatch

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/volume/%s", hostURL, volumeID)
	log.Println("update url ==", url)
	log.Println("ID = ", volumeID)
	requestBody, err := json.Marshal(vol)
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
func (c *Client) DeleteVolume(clientInterface ClientInterface, hostURL string, volumeID string) error {

	reqType := http.MethodDelete

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/volume/%s", hostURL, volumeID)

	_, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return err
	}
	return nil
}
