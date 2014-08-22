package services

import (
	"errors"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayerVirtualGuest struct {
	client softlayer.Client
}

func NewSoftLayer_Virtual_Guest(client softlayer.Client) *softLayerVirtualGuest {
	return &softLayerVirtualGuest{
		client: client,
	}
}

func (slvg *softLayerVirtualGuest) GetName() string {
	return "SoftLayer_Virtual_Guest"
}

func (slvg *softLayerVirtualGuest) CreateObject(template datatypes.SoftLayer_Virtual_Guest) (datatypes.SoftLayer_Virtual_Guest, error) {
	return datatypes.SoftLayer_Virtual_Guest{}, errors.New("Implement me!")
}

func (slvg *softLayerVirtualGuest) DeleteObject(template datatypes.SoftLayer_Virtual_Guest) (bool, error) {
	return false, errors.New("Implement me!")
}
