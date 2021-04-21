package powerstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//SnapshotRuleInterface - interface of all client functions
type SnapshotRuleInterface interface {
	GetSnapshotRule(ClientInterface, string, string) (*SnapshotRule, error)
}

// GetSnapshotRule - Returns a specifc snapshot rule
func (c *Client) GetSnapshotRule(clientInterface ClientInterface, hostURL string, snapshotRuleID string) (*SnapshotRule, error) {

	reqType := http.MethodGet

	// URL for snapshotRule resource type
	url := fmt.Sprintf("%s/api/rest/snapshot_rule?select=*/%s", hostURL, snapshotRuleID)

	body, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return nil, err
	}
	log.Println("body ===", string(body))
	snapshotRule := SnapshotRule{}
	err = json.Unmarshal(body, &snapshotRule)
	if err != nil {
		return nil, err
	}
	return &snapshotRule, nil
}

//CreateSnapshotRule - Create a  snapshot rule
func (c *Client) CreateSnapshotRule(clientInterface ClientInterface, hostURL string, snapshotRule SnapshotRule) (string, error) {

	reqType := http.MethodPost

	// URL for snapshotRule resource type
	url := fmt.Sprintf("%s/api/rest/snapshot_rule", hostURL)
	requestBody, err := json.Marshal(snapshotRule)
	if err != nil {
		return "", err
	}
	body, err := clientInterface.doRequest(reqType, url, string(requestBody))
	if err != nil {
		return "", err
	}
	log.Println("body ===", string(body))
	result := SnapshotRule{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.ID, nil
}

//UpdateSnapshotRule - Update a snapshot rule based on snapshot rule id
func (c *Client) UpdateSnapshotRule(clientInterface ClientInterface, hostURL string, snapshotRule SnapshotRule, snapshotID string) error {

	reqType := http.MethodPatch

	// URL for snapshotRule resource type
	url := fmt.Sprintf("%s/api/rest/snapshot_rule/%s", hostURL, snapshotID)
	requestBody, err := json.Marshal(snapshotRule)
	if err != nil {
		return err
	}
	_, err = clientInterface.doRequest(reqType, url, string(requestBody))
	if err != nil {
		return err
	}

	return nil
}

//DeleteSnapshotRule - Delete a snapshot rule based on snapshot rule id
func (c *Client) DeleteSnapshotRule(clientInterface ClientInterface, hostURL string, snapshotID string) error {

	reqType := http.MethodDelete

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/snapshot_rule/%s", hostURL, snapshotID)

	_, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return err
	}
	return nil
}
