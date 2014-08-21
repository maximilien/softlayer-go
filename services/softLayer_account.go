package services

import (
	"errors"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const NAME = "SoftLayer_Account"

type SoftLayer_Account interface {
	softlayer.Service

	GetVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error)
}

type softLayer_Account struct {
	client softlayer.Client
}

func NewSoftLayer_Account(client softlayer.Client) *softLayer_Account {
	return &softLayer_Account{
		client: client,
	}
}

func GetName() string {
	return NAME
}

func GetVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error) {
	return []datatypes.SoftLayer_Virtual_Guest{}, errors.New("Implement me!")
}
