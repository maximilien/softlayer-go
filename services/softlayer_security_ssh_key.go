package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Ssh_Key_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Ssh_Key_Service(client softlayer.Client) *softLayer_Ssh_Key_Service {
	return &softLayer_Ssh_Key_Service{
		client: client,
	}
}

func (slsks *softLayer_Ssh_Key_Service) GetName() string {
	return "SoftLayer_Security_Ssh_Key"
}

func (slsks *softLayer_Ssh_Key_Service) CreateObject(template datatypes.SoftLayer_Security_Ssh_Key) (datatypes.SoftLayer_Security_Ssh_Key, error) {
	parameters := datatypes.SoftLayer_Shh_Key_Parameters{
		Parameters: []datatypes.SoftLayer_Security_Ssh_Key{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	data, err := slsks.client.DoRawHttpRequest(fmt.Sprintf("%s/createObject", slsks.GetName()), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	err = slsks.client.CheckForHttpResponseErrors(data)
	if err != nil {
		return datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	softLayer_Ssh_Key := datatypes.SoftLayer_Security_Ssh_Key{}
	err = json.Unmarshal(data, &softLayer_Ssh_Key)
	if err != nil {
		return datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	return softLayer_Ssh_Key, nil
}

func (slsks *softLayer_Ssh_Key_Service) DeleteObject(sshKeyId int) (bool, error) {
	response, err := slsks.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", slsks.GetName(), sshKeyId), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to destroy ssh key with id '%d', got '%s' as response from the API.", sshKeyId, res))
	}

	return true, err
}
