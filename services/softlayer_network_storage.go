package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const (
	NETWORK_STORAGE_PACKAGE_ID         = 0
	CREATE_ISCSI_VOLUME_MAX_RETRY_TIME = 3
	CREATE_ISCSI_VOLUME_CHECK_INTERVAL = 20 // seconds
)

type softLayer_Network_Storage_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Network_Storage_Service(client softlayer.Client) *softLayer_Network_Storage_Service {
	return &softLayer_Network_Storage_Service{
		client: client,
	}
}

func (slns *softLayer_Network_Storage_Service) GetName() string {
	return "SoftLayer_Network_Storage"
}

func (slns *softLayer_Network_Storage_Service) CreateIscsiVolume(size int, location string) (datatypes.SoftLayer_Network_Storage, error) {
	order := datatypes.SoftLayer_Product_Order{
		Location:    "138124", //TODO: now using default dal05 data center, need to use the location string passed in
		ComplexType: "SoftLayer_Container_Product_Order",
		Prices: []datatypes.SoftLayer_Item_Price{
			datatypes.SoftLayer_Item_Price{
				Id: 30587, //TODO: now using default 20GB volume, need to use product_package service to query the iSCSI item id based on the disk size
			},
		},
		PackageId: NETWORK_STORAGE_PACKAGE_ID,
	}

	productOrderService, _ := slns.client.GetSoftLayer_Product_Order_Service()
	receipt, err := productOrderService.PlaceOrder(order)
	if err != nil {
		return datatypes.SoftLayer_Network_Storage{}, err
	}

	var iscsiStorage datatypes.SoftLayer_Network_Storage

	for i := 0; i < CREATE_ISCSI_VOLUME_MAX_RETRY_TIME; i++ {
		iscsiStorage, err = slns.findIscsiVolumeIdByOrderId(receipt.OrderId)
		if err == nil {
			break
		} else if i == CREATE_ISCSI_VOLUME_MAX_RETRY_TIME-1 {
			return datatypes.SoftLayer_Network_Storage{}, err
		}

		time.Sleep(CREATE_ISCSI_VOLUME_CHECK_INTERVAL * time.Second)
	}

	return iscsiStorage, nil
}

func (slns *softLayer_Network_Storage_Service) DeleteIscsiVolume(volumeId int, immediateCancellationFlag bool) error {
	accountService, _ := slns.client.GetSoftLayer_Account_Service()
	iscsiStorages, _ := accountService.GetIscsiNetworkStorage()

	var accountId, billingItemId int

	for i := 0; i < len(iscsiStorages); i++ {
		if iscsiStorages[i].Id == volumeId {
			accountId = iscsiStorages[i].AccountId
			billingItemId = iscsiStorages[i].BillingItem.Id
			break
		}
	}

	billingItemCancellationRequest := datatypes.SoftLayer_Billing_Item_Cancellation_Request{
		ComplexType: "SoftLayer_Billing_Item_Cancellation_Request",
		AccountId:   accountId,
		Items: []datatypes.SoftLayer_Billing_Item_Cancellation_Request_Item{
			{
				BillingItemId:             billingItemId,
				ImmediateCancellationFlag: immediateCancellationFlag,
			},
		},
	}

	billingItemCancellationRequestService, _ := slns.client.GetSoftLayer_Billing_Item_Cancellation_Request_Service()
	_, err := billingItemCancellationRequestService.CreateObject(billingItemCancellationRequest)
	if err != nil {
		return err
	}

	return nil
}

func (slns *softLayer_Network_Storage_Service) GetIscsiVolume(volumeId int) (datatypes.SoftLayer_Network_Storage, error) {
	response, err := slns.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getObject.json", slns.GetName(), volumeId), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Network_Storage{}, err
	}

	volume := datatypes.SoftLayer_Network_Storage{}
	err = json.Unmarshal(response, &volume)
	if err != nil {
		return datatypes.SoftLayer_Network_Storage{}, err
	}

	return volume, nil
}

// Private methods

func (slns *softLayer_Network_Storage_Service) findIscsiVolumeIdByOrderId(orderId int) (datatypes.SoftLayer_Network_Storage, error) {
	accountService, _ := slns.client.GetSoftLayer_Account_Service()
	iscsiStorages, _ := accountService.GetIscsiNetworkStorage()

	for i := 0; i < len(iscsiStorages); i++ {
		storage := iscsiStorages[i]

		if storage.BillingItem != nil && storage.BillingItem.OrderItem.Order.Id == orderId {
			return storage, nil
		}
	}

	return datatypes.SoftLayer_Network_Storage{}, errors.New(fmt.Sprintf("Can not find an iSCSI volume with order id %d", orderId))
}
