package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
)

const (
	DATACENTER_TYPE_NAME   = "SoftLayer_Location_Datacenter"
	ROUTING_TYPE_NAME      = "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type"
	ROUTING_METHOD_NAME    = "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method"
	HEALTH_CHECK_TYPE_NAME = "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type"
)

func GetDatacenterByName(client softlayer.Client, name string) (int, error) {
	ObjectFilter := string(`{"name":{"operation":"` + name + `"}}`)
	ObjectMasks := []string{"id", "name"}

	response, errorCode, err := client.GetHttpClient().DoRawHttpRequestWithObjectFilterAndObjectMask(fmt.Sprintf("%s/getDatacenters.json", DATACENTER_TYPE_NAME), ObjectMasks, ObjectFilter, "GET", new(bytes.Buffer))
	if err != nil {
		return -1, err
	}

	if IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not retrieve datacenters, HTTP error code: '%d'", errorCode)
		return -1, errors.New(errorMessage)
	}

	locations := []datatypes.SoftLayer_Location{}
	err = json.Unmarshal(response, &locations)
	if err != nil {
		return -1, err
	}

	for _, location := range locations {
		if location.Name == name {
			return location.Id, nil
		}
	}

	return -1, fmt.Errorf("Datacenter %s not found", name)
}

func GetRoutingTypeByName(client softlayer.Client, name string) (int, error) {
	ObjectFilter := string(`{"keyname":{"operation":"` + name + `"}}`)
	ObjectMasks := []string{"id", "keyname"}

	response, errorCode, err := client.GetHttpClient().DoRawHttpRequestWithObjectFilterAndObjectMask(fmt.Sprintf("%s/getAllObjects.json", ROUTING_TYPE_NAME), ObjectMasks, ObjectFilter, "GET", new(bytes.Buffer))
	if err != nil {
		return -1, err
	}

	if IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not retrieve routing types, HTTP error code: '%d'", errorCode)
		return -1, errors.New(errorMessage)
	}

	routingTypes := []datatypes.SoftLayer_Routing_Type{}
	err = json.Unmarshal(response, &routingTypes)
	if err != nil {
		return -1, err
	}

	for _, routingType := range routingTypes {
		if routingType.KeyName == name {
			return routingType.Id, nil
		}
	}

	return -1, fmt.Errorf("Routing type %s not found", name)
}

func GetRoutingMethodByName(client softlayer.Client, name string) (int, error) {
	ObjectFilter := string(`{"keyname":{"operation":"` + name + `"}}`)
	ObjectMasks := []string{"id", "keyname"}

	response, errorCode, err := client.GetHttpClient().DoRawHttpRequestWithObjectFilterAndObjectMask(fmt.Sprintf("%s/getAllObjects.json", ROUTING_METHOD_NAME), ObjectMasks, ObjectFilter, "GET", new(bytes.Buffer))
	if err != nil {
		return -1, err
	}

	if IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not retrieve routing methods, HTTP error code: '%d'", errorCode)
		return -1, errors.New(errorMessage)
	}

	routingMethods := []datatypes.SoftLayer_Routing_Method{}
	err = json.Unmarshal(response, &routingMethods)
	if err != nil {
		return -1, err
	}

	for _, routingMethod := range routingMethods {
		if routingMethod.KeyName == name {
			return routingMethod.Id, nil
		}
	}

	return -1, fmt.Errorf("Routing method %s not found", name)
}

func GetHealthCheckTypeByName(client softlayer.Client, name string) (int, error) {
	ObjectFilter := string(`{"keyname":{"operation":"` + name + `"}}`)
	ObjectMasks := []string{"id", "keyname"}

	response, errorCode, err := client.GetHttpClient().DoRawHttpRequestWithObjectFilterAndObjectMask(fmt.Sprintf("%s/getAllObjects.json", HEALTH_CHECK_TYPE_NAME), ObjectMasks, ObjectFilter, "GET", new(bytes.Buffer))
	if err != nil {
		return -1, err
	}

	if IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not retrieve health check types, HTTP error code: '%d'", errorCode)
		return -1, errors.New(errorMessage)
	}

	healthCheckTypes := []datatypes.SoftLayer_Health_Check_Type{}
	err = json.Unmarshal(response, &healthCheckTypes)
	if err != nil {
		return -1, err
	}

	for _, healthCheckType := range healthCheckTypes {
		if healthCheckType.KeyName == name {
			return healthCheckType.Id, nil
		}
	}

	return -1, fmt.Errorf("Health check type %s not found", name)
}
