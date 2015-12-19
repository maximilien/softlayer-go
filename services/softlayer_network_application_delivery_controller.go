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

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) CreateNetscalerVPX(createOptions *softlayer.NetworkApplicationDeliveryControllerCreateOptions) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error) {
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
	// TODO wait here ?

	return vpx, nil
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) CreateVirtualIpAddress(nadcId int, template datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template) (bool, error) {
	// check required fields

	nadc, err := slnadcs.GetObject(nadcId)
	if err != nil {
		return false, err
	}
	if nadc.Id != nadcId {
		err = errors.New(fmt.Sprintf("Network application delivery controller with id %d is not found", nadcId))
		return false, err
	}

	parameters := datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template_Parameters{
		LoadBalancer: template,
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return false, err
	}

	response, err := slnadcs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/%s.json", slnadcs.GetName(), nadcId, "createLiveLoadBalancer"), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}

	if response_value := string(response[:]); response_value != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete Virtual IP Address with id '%s' from network application delivery controller %d. got '%s' as response from the API", 0, nadcId, response_value))
	}

	return true, nil
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) DeleteVirtualIpAddress(nadcId int, name string) (bool, error) {
	nadc, err := slnadcs.GetObject(nadcId)
	if err != nil {
		return false, err
	}
	if nadc.Id != nadcId {
		err = errors.New(fmt.Sprintf("Network application delivery controller with id %d is not found", nadcId))
		return false, err
	}

	parameters := datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template_Parameters{
		LoadBalancer: datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template{
			Name: name,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return false, err
	}

	response, err := slnadcs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/%s.json", slnadcs.GetName(), nadcId, "deleteLiveLoadBalancer"), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}

	if response_value := string(response[:]); response_value != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete Virtual IP Address with name '%s' from network application delivery controller %d. got '%s' as response from the API", name, nadcId, response_value))
	}

	return true, err
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) EditVirtualIpAddress(nadcId int, template datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template) (bool, error) {
	nadc, err := slnadcs.GetObject(nadcId)
	if err != nil {
		return false, err
	}
	if nadc.Id != nadcId {
		err = errors.New(fmt.Sprintf("Network application delivery controller with id %d is not found", nadcId))
		return false, err
	}

	parameters := datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template_Parameters{
		LoadBalancer: template,
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return false, err
	}

	response, err := slnadcs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/%s.json", slnadcs.GetName(), nadcId, "updateLiveLoadBalancer"), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}

	if response_value := string(response[:]); response_value != "true" {
		return false, errors.New(fmt.Sprintf("Failed to update Virtual IP Address with id '%d' from network application delivery controller %d. got '%s' as response from the API", template.Id, nadcId, response_value))
	}

	return true, err
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) GetVirtualIpAddress(nadcId int, vipId int) (datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress, error) {
	nadc, err := slnadcs.GetObject(nadcId)
	if err != nil {
		return datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress{}, err
	}
	if nadc.Id != nadcId {
		err = errors.New(fmt.Sprintf("Network application delivery controller with id %d is not found", nadcId))
		return datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress{}, err
	}

	response, err := slnadcs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", slnadcs.GetName(), nadcId, "getVirtualIpAddresses"), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress{}, err
	}

	addresses := datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Array{}
	err = json.Unmarshal(response, &addresses)
	if err != nil {
		return datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress{}, err
	}

	var result datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress
	for _, address := range addresses {
		if address.Id == vipId {
			result = address
			break;
		}
	}

	return result, err
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) GetObject(id int) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error) {

	objectMask := []string{
		"id",
		"createDate",
		"name",
		"typeId",
		"modifyDate",
		"description",
		"managedResourceFlag",
		"managementIpAddress",
		"primaryIpAddress",
		"password",
		"notes",
		"datacenter",
		"averageDailyPublicBandwidthUsage",
		"licenseExpirationDate",
		"networkVlan",
		"networkVlanCount",
		"networkVlans",
		"subnetCount",
		"subnets",
		"tagReferenceCount",
		"tagReferences",
		"type",
		"virtualIpAddressCount",
		"virtualIpAddresses",
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

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) checkCreateVpxRequiredValues(createOptions *softlayer.NetworkApplicationDeliveryControllerCreateOptions) error {
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
		err = errors.New(strings.Join(errorMessages, "\n"))
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

// create item key for Netscaler VPX, based on the provided version, speed and plan
func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) getVPXPriceItemKeyName(version string, speed int, plan string) string {
	name := "CITRIX_NETSCALER_VPX"
	speedMeasurements := "MBPS"
	versionReplaced := strings.Replace(version, ".", DELIMITER, -1)
	speedString := strconv.Itoa(speed) + speedMeasurements

	return strings.Join([]string{name, versionReplaced, speedString, plan}, DELIMITER)
}

// create item key for Netscaler VPX, based on provided ips count
func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) getPublicIpItemKeyName(ipCount int) string {
	name := "STATIC_PUBLIC_IP_ADDRESSES"
	ipCountString := strconv.Itoa(ipCount)

	return strings.Join([]string{name, ipCountString}, DELIMITER)
}

// use the create options to build keys for price items
// using these keys the desired price items are found
func (slnadcs *softLayer_Network_Application_Delivery_Controller_Service) findCreatePriceItems(createOptions *softlayer.NetworkApplicationDeliveryControllerCreateOptions) ([]*datatypes.SoftLayer_Item_Price, error) {
	items, err := slnadcs.getApplicationDeliveryControllerItems()
	if err != nil {
		return []*datatypes.SoftLayer_Item_Price{}, err
	}

	// build price item keys based on the configuration values
	nadcKey := slnadcs.getVPXPriceItemKeyName(createOptions.Version, createOptions.Speed, createOptions.Plan)
	ipKey := slnadcs.getPublicIpItemKeyName(createOptions.IpCount)

	var nadcItemPrice, ipItemPrice *datatypes.SoftLayer_Item_Price

	// find the price items by keys
	for _, item := range items {
		itemKey := item.Key
		if itemKey == nadcKey {
			nadcItemPrice = &item.Prices[0]
		}
		if itemKey == ipKey {
			ipItemPrice = &item.Prices[0]
		}
	}

	var errorMessages []string

	if nadcItemPrice == nil {
		errorMessages = append(errorMessages, fmt.Sprintf("VPX version, speed or plan have incorrect values"))
	}

	if ipItemPrice == nil {
		errorMessages = append(errorMessages, fmt.Sprintf("Ip quantity value is incorrect"))
	}

	if len(errorMessages) > 0 {
		err = errors.New(strings.Join(errorMessages, "\n"))
		return []*datatypes.SoftLayer_Item_Price{}, err
	}

	return []*datatypes.SoftLayer_Item_Price{nadcItemPrice, ipItemPrice}, nil
}
