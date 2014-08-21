package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const NAME = "SoftLayer_Account"

type softLayer_Account struct {
	client softlayer.Client
}

func NewSoftLayer_Account(client softlayer.Client) *softLayer_Account {
	return &softLayer_Account{
		client: client,
	}
}

func (sla *softLayer_Account) GetName() string {
	return NAME
}

func (sla *softLayer_Account) GetVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error) {
	path := fmt.Sprintf("%s/%s", sla.GetName(), "getVirtualGuests.json")
	responseBytes, err := sla.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getVirtualGuests, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Virtual_Guest{}, errors.New(errorMessage)
	}

	virtualGuests := []datatypes.SoftLayer_Virtual_Guest{}
	err = json.Unmarshal(responseBytes, &virtualGuests)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return nil, err
	}

	return virtualGuests, nil
}
