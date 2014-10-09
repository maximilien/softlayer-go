package services

import (
	"bytes"
	"encoding/json"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Product_Order_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Product_Order_Service(client softlayer.Client) *softLayer_Product_Order_Service {
	return &softLayer_Product_Order_Service{
		client: client,
	}
}

func (slpo *softLayer_Product_Order_Service) GetName() string {
	return "SoftLayer_Product_Order"
}

func (slpo *softLayer_Product_Order_Service) PlaceOrder(order datatypes.SoftLayer_Product_Order) (datatypes.SoftLayer_Product_Order_Receipt, error) {
	parameters := datatypes.SoftLayer_Product_Order_Parameters{
		Parameters: []datatypes.SoftLayer_Product_Order{
			order,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return datatypes.SoftLayer_Product_Order_Receipt{}, err
	}

	responseBytes, err := slpo.client.DoRawHttpRequest(fmt.Sprintf("%s/placeOrder.json", slpo.GetName()), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return datatypes.SoftLayer_Product_Order_Receipt{}, err
	}

	receipt := datatypes.SoftLayer_Product_Order_Receipt{}
	err = json.Unmarshal(responseBytes, &receipt)
	if err != nil {
		return datatypes.SoftLayer_Product_Order_Receipt{}, err
	}

	return receipt, nil
}
