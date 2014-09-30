package services

import (
	"bytes"
	"encoding/json"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Product_Package_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Product_Package_Service(client softlayer.Client) *softLayer_Product_Package_Service {
	return &softLayer_Product_Package_Service{
		client: client,
	}
}

func (slpp *softLayer_Product_Package_Service) GetName() string {
	return "SoftLayer_Product_Package"
}

func (slpp *softLayer_Product_Package_Service) GetItemPrices(packageId int) ([]datatypes.SoftLayer_Item_Price, error) {
	response, err := slpp.client.DoRawHttpRequestWithObjectMask(fmt.Sprintf("%s/%d/getItemPrices.json", slpp.GetName(), packageId), []string{"id", "item.id", "item.description", "item.capacity"}, "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Item_Price{}, err
	}

	itemPrices := []datatypes.SoftLayer_Item_Price{}
	err = json.Unmarshal(response, &itemPrices)
	if err != nil {
		return []datatypes.SoftLayer_Item_Price{}, err
	}

	return itemPrices, nil
}
