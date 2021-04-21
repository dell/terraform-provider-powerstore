package powerstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetProtectionPolicy(clientInterface ClientInterface, hostURL string, protectionPolicyID string) (*ProtectionPolicy, error) {

	reqType := http.MethodGet

	// URL for protectionPolicy resource type
	url := fmt.Sprintf("%s/api/rest/policy/%s?select=name,id,description,type,file_systems(id,name),replication_rules(id,name),snapshot_rules(id,name,interval,time_of_day,days_of_week),virtual_machines(id,name),volume(id,name),volume_group(id,name)&type=eq.Protection", hostURL, protectionPolicyID)

	body, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return nil, err
	}
	log.Println("body ===", string(body))
	protectionPolicy := ProtectionPolicy{}
	err = json.Unmarshal(body, &protectionPolicy)
	if err != nil {
		return nil, err
	}
	return &protectionPolicy, nil
}

func (c *Client) UpdateProtectionPolicy(clientInterface ClientInterface, hostURL string, protectionPolicy ProtectionPolicyRequest, protectionPolicyID string) error {
	reqType := http.MethodPatch

	// URL for protectionPolicy resource type
	url := fmt.Sprintf("%s/api/rest/policy/%s", hostURL, protectionPolicyID)
	requestBody, err := json.Marshal(protectionPolicy)
	if err != nil {
		return err
	}
	_, err = clientInterface.doRequest(reqType, url, string(requestBody))
	if err != nil {
		return err
	}

	return nil
}

// CreateProtectionPolicy - Create a ProtectionPolicy and returns ProtectionPolicy id
func (c *Client) CreateProtectionPolicy(clientInterface ClientInterface, hostURL string, protectionPolicy ProtectionPolicyRequest) (string, error) {

	reqType := http.MethodPost

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/policy", hostURL)

	requestBody, err := json.Marshal(protectionPolicy)
	if err != nil {
		return "", err
	}

	body, err := clientInterface.doRequest(reqType, url, string(requestBody))
	if err != nil {
		return "", err
	}
	log.Println("body ===", string(body))
	policy := ProtectionPolicyRequest{}
	err = json.Unmarshal(body, &policy)
	if err != nil {
		return "", err
	}

	return policy.ID, nil
}

//DeleteProtectionPolicy - Delete a protection policy based on protection policy id
func (c *Client) DeleteProtectionPolicy(clientInterface ClientInterface, hostURL string, protectionPolicyID string) error {

	reqType := http.MethodDelete

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/policy/%s", hostURL, protectionPolicyID)

	_, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return err
	}
	return nil
}
