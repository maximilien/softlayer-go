package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	"strconv"
	"strings"
)

const (
	PACKAGE_TYPE_APPLICATION_DELIVERY_CONTROLLER = "ADDITIONAL_SERVICES_APPLICATION_DELIVERY_APPLIANCE"
	ORDER_TYPE_APPLICATION_DELIVERY_CONTROLLER   = "SoftLayer_Container_Product_Order_Network_Application_Delivery_Controller"
	PACKAGE_ID_APPLICATION_DELIVERY_CONTROLLER   = 192
	DELIMITER                                    = "_"
)

type softLayer_Network_Application_Delivery_Controller_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Network_Application_Delivery_Controller_Service(client softlayer.Client) *softLayer_Network_Application_Delivery_Controller_Service {
	return &softLayer_Network_Application_Delivery_Controller_Service{
		client: client,
	}
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) GetName() string {
	return "SoftLayer_Network_Application_Delivery_Controller"
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) CreateNetscalerVPX(createOptions *softlayer.CreateOptions) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error) {
	// check required fields
	err := slnadcs.checkCreateVpxRequiredValues(createOptions)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	orderService, err := slnadcs.client.GetSoftLayer_Product_Order_Service()
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	items, err := slnadcs.findCreatePriceItems(createOptions)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	order := datatypes.SoftLayer_Container_Product_Order_Network_Application_Delivery_Controller{
		PackageId:   PACKAGE_ID_APPLICATION_DELIVERY_CONTROLLER,
		ComplexType: ORDER_TYPE_APPLICATION_DELIVERY_CONTROLLER,
		Location:    createOptions.Location,
		Prices:      items,
		Quantity:    1,
	}

	receipt, err := orderService.PlaceContainerOrderApplicationDeliveryController(order)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	vpx, err := slnadcs.findVPXByOrderId(receipt.OrderId)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	// TODO maybe call GetObject here

	return vpx, nil
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) GetObject(id int) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error) {

	objectMask := []string{
		"id",
		"name",
		"typeId",
		"modifyDate",
		"description",
		"managedResourceFlag",
		"managementIpAddress",
		"primaryIpAddress",
		"password",
	}

	response, err := slnadcs.client.DoRawHttpRequestWithObjectMask(fmt.Sprintf("%s/%d/getObject.json", slnadcs.GetName(), id), objectMask, "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	nadc := datatypes.SoftLayer_Network_Application_Delivery_Controller{}
	err = json.Unmarshal(response, &nadc)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	return nadc, nil
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) DeleteObject(id int) (bool, error) {
	response, err := slnadcs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", slnadcs.GetName(), id), "DELETE", new(bytes.Buffer))

	if response_value := string(response[:]); response_value != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete Application Delivery Controller with id '%d', got '%s' as response from the API", id, response_value))
	}

	return true, err
}

// Private methods

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) checkCreateVpxRequiredValues(createOptions *softlayer.CreateOptions) error {
	var err error
	var errorMessages []string
	errorTemplate := "* %s is required and cannot be empty\n"

	if createOptions.Plan == "" {
		errorMessages = append(errorMessages, fmt.Sprintf(errorTemplate, "Vpx Plan"))
	}

	if createOptions.Speed <= 0 {
		errorMessages = append(errorMessages, fmt.Sprintf(errorTemplate, "Network speed"))
	}

	if createOptions.Version == "" {
		errorMessages = append(errorMessages, fmt.Sprintf(errorTemplate, "Vpx version"))
	}

	if createOptions.Location == "" {
		errorMessages = append(errorMessages, fmt.Sprintf(errorTemplate, "Location"))
	}

	if len(errorMessages) > 0 {
		err = errors.New(strings.Join(errorMessages, '\n'))
	}

	return err
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) findVPXByOrderId(orderId int) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error) {
	ObjectFilter := string(`{"iscsiNetworkStorage":{"billingItem":{"orderItem":{"order":{"id":{"operation":` + strconv.Itoa(orderId) + `}}}}}}`)

	accountService, err := slnadcs.client.GetSoftLayer_Account_Service()
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	vpxs, err := accountService.GetApplicationDeliveryControllersWithFilter(ObjectFilter)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller{}, err
	}

	if len(vpxs) == 1 {
		return vpxs[0], nil
	}

	return datatypes.SoftLayer_Network_Application_Delivery_Controller{},
		errors.New(fmt.Sprintf("Cannot find Application Delivery Controller with order id %d", orderId))
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) getApplicationDeliveryControllerItems() ([]datatypes.SoftLayer_Product_Item, error) {
	productPackageService, err := slnadcs.client.GetSoftLayer_Product_Package_Service()
	if err != nil {
		return []datatypes.SoftLayer_Product_Item{}, err
	}

	return productPackageService.GetItemsByType(PACKAGE_TYPE_APPLICATION_DELIVERY_CONTROLLER)
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) getVPXPriceItemKeyName(version string, speed int, plan string) string {
	name := "CITRIX_NETSCALER_VPX"
	speedMeasurements := "MBPS"
	versionReplaced := strings.Replace(version, '.', DELIMITER, -1)
	speedString := strconv.Itoa(speed) + speedMeasurements

	return strings.Join([]string{name, versionReplaced, speedString, plan}, DELIMITER)
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) getPublicIpItemKeyName(ipCount int) string {
	name := "STATIC_PUBLIC_IP_ADDRESSES"
	ipCountString := strconv.Itoa(ipCount)

	return strings.Join([]string{name, ipCountString}, DELIMITER)
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) findCreatePriceItems(createOptions *softlayer.CreateOptions) ([]datatypes.SoftLayer_Item_Price, error) {
	items, err := slnadcs.getApplicationDeliveryControllerItems()
	if err != nil {
		return []datatypes.SoftLayer_Item_Price{}, err
	}

	adcKey := slnadcs.getVPXPriceItemKeyName(createOptions.Version, createOptions.Speed, createOptions.Plan)
	ipKey := slnadcs.getPublicIpItemKeyName(createOptions.IpCount)

	var resultList [0]datatypes.SoftLayer_Item_Price

	// TODO test this for cycle
	for _, item := range items {
		itemKey := item.Key
		if itemKey == adcKey || itemKey == ipKey {
			resultList = append(resultList, item.Prices[0])
		}
	}

	return resultList
}
