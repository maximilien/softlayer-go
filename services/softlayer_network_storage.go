package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const (
	NETWORK_PERFORMANCE_STORAGE_PACKAGE_ID = 222
	BLOCK_ITEM_PRICE_ID                    = 40678 // file or block item price id
	CREATE_ISCSI_VOLUME_MAX_RETRY_TIME     = 12
	CREATE_ISCSI_VOLUME_CHECK_INTERVAL     = 5 // seconds
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
	if size < 0 {
		return datatypes.SoftLayer_Network_Storage{}, errors.New("Cannot create negative sized volumes")
	}

	sizeItemPriceId, err := slns.getIscsiVolumeItemIdBasedOnSize(size)
	iopsItemPriceId, err := slns.getPerformanceStorageItemPriceIdByIops(size)

	/*if err != nil {
		return datatypes.SoftLayer_Network_Storage{}, err
	}*/

	order := datatypes.SoftLayer_Product_Order{
		Location:    location,
		ComplexType: "SoftLayer_Container_Product_Order_Network_PerformanceStorage_Iscsi",
		OsFormatType: datatypes.OsFormatType{
			Id:      12,
			KeyName: "LINUX",
		},
		Prices: []datatypes.SoftLayer_Item_Price{
			datatypes.SoftLayer_Item_Price{
				Id: sizeItemPriceId,
			},
			datatypes.SoftLayer_Item_Price{
				Id: iopsItemPriceId,
			},
			datatypes.SoftLayer_Item_Price{
				Id: BLOCK_ITEM_PRICE_ID,
			},
		},
		PackageId: NETWORK_PERFORMANCE_STORAGE_PACKAGE_ID,
		Quantity:  1,
	}

	productOrderService, _ := slns.client.GetSoftLayer_Product_Order_Service()
	receipt, err := productOrderService.PlaceOrder(order)
	if err != nil {
		return datatypes.SoftLayer_Network_Storage{}, err
	}

	var iscsiStorage datatypes.SoftLayer_Network_Storage

	for i := 0; i < CREATE_ISCSI_VOLUME_MAX_RETRY_TIME; i++ {
		iscsiStorage, err = slns.findIscsiVolumeId(receipt.OrderId)
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
	ObjectFilter := string(`{"iscsiNetworkStorage":{"id":{"operation":` + strconv.Itoa(volumeId) + `}}}`)
	accountService, _ := slns.client.GetSoftLayer_Account_Service()
	iscsiStorages, _ := accountService.GetIscsiNetworkStorageWithFilter(ObjectFilter)

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

func (slns *softLayer_Network_Storage_Service) findIscsiVolumeId(orderId int) (datatypes.SoftLayer_Network_Storage, error) {
	ObjectFilter := string(`{"iscsiNetworkStorage":{"billingItem":{"orderItem":{"order":{"id":{"operation":` + strconv.Itoa(orderId) + `}}}}}}`)
	accountService, _ := slns.client.GetSoftLayer_Account_Service()

	iscsiStorages, _ := accountService.GetIscsiNetworkStorageWithFilter(ObjectFilter)

	if len(iscsiStorages) == 1 {
		return iscsiStorages[0], nil
	}

	return datatypes.SoftLayer_Network_Storage{}, errors.New(fmt.Sprintf("Can not find an performance storage (iSCSI volume) with order id %d", orderId))
}

func (slns *softLayer_Network_Storage_Service) getIscsiVolumeItemIdBasedOnSize(size int) (int, error) {
	productPackageService, err := slns.client.GetSoftLayer_Product_Package_Service()
	if err != nil {
		return 0, err
	}

	itemPrices, err := productPackageService.GetItemPricesBySize(NETWORK_PERFORMANCE_STORAGE_PACKAGE_ID, size)
	if err != nil {
		return 0, err
	}

	var currentItemId int

	if len(itemPrices) > 0 {
		for _, itemPrice := range itemPrices {
			if itemPrice.LocationGroupId == 0 {
				currentItemId = itemPrice.Id
			}
		}
	}

	if currentItemId == 0 {
		return 0, errors.New(fmt.Sprintf("No proper performance storage (iSCSI volume)for size %d", size))
	}

	return currentItemId, nil
}

func (slns *softLayer_Network_Storage_Service) getPerformanceStorageItemPriceIdByIops(size int) (int, error) {
	switch size {
	case 20:
		return 40838, nil // 500 IOPS
	case 40:
		return 40988, nil // 1000 IOPS
	case 80:
		return 41288, nil // 2000 IOPS
	default:
		return 41788, nil // 3000 IOPS
	}
}
