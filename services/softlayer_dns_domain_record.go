package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"errors"

	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	softlayer "github.com/TheWeatherCompany/softlayer-go/softlayer"
)


type softLayer_Dns_Domain_Record_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Dns_Domain_Record_Service(client softlayer.Client) *softLayer_Dns_Domain_Record_Service {
	return &softLayer_Dns_Domain_Record_Service{
		client: client,
	}
}

func (sldr *softLayer_Dns_Domain_Record_Service) GetName() string {
	return "SoftLayer_Dns_Domain_ResourceRecord"
}

func (sldr *softLayer_Dns_Domain_Record_Service) CreateObject(template datatypes.SoftLayer_Dns_Domain_Record_Template) (datatypes.SoftLayer_Dns_Domain_Record, error) {
	parameters := datatypes.SoftLayer_Dns_Domain_Record_Template_Parameters{
		Parameters: []datatypes.SoftLayer_Dns_Domain_Record_Template{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Record{}, err
	}

	response, err := sldr.client.DoRawHttpRequest(fmt.Sprintf("%s.json", sldr.GetName()), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Record{}, err
	}

	err = sldr.client.CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Record{}, err
	}

	dns_record := datatypes.SoftLayer_Dns_Domain_Record{}
	err = json.Unmarshal(response, &dns_record)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Record{}, err
	}

	return dns_record, nil
}

func (sldr *softLayer_Dns_Domain_Record_Service) CreateObjects(templates []datatypes.SoftLayer_Dns_Domain_Record_Template) ([]datatypes.SoftLayer_Dns_Domain_Record, error) {
	//TODO to implement
	return nil, nil
}

func (sldr *softLayer_Dns_Domain_Record_Service) GetObject(id int) (datatypes.SoftLayer_Dns_Domain_Record, error) {

	objectMask := []string{
		"data",
		"domainId",
		"expire",
		"host",
		"id",
		"minimum",
		"mxPriority",
		"refresh",
		"responsiblePerson",
		"retry",
		"ttl",
		"type",
	}

	response, err := sldr.client.DoRawHttpRequestWithObjectMask(fmt.Sprintf("%s/%s.json", sldr.GetName(), id), objectMask, "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Record{}, err
	}

	err = sldr.client.CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Record{}, err
	}

	dns_record := datatypes.SoftLayer_Dns_Domain_Record{}
	err = json.Unmarshal(response, &dns_record)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Record{}, err
	}

	return dns_record, nil
}

func (sldr *softLayer_Dns_Domain_Record_Service) DeleteObject(recordId int) (bool, error) {
	response, err := sldr.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", sldr.GetName(), recordId), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete dns domain record with id '%d', got '%s' as response from the API.", recordId, res))
	}

	return true, err
}

func (sldr *softLayer_Dns_Domain_Record_Service) DeleteObjects(recordIds []int) (bool, error) {
//	response, err := sldr.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", sldr.GetName(), recordId), "DELETE", new(bytes.Buffer))
//
//	if res := string(response[:]); res != "true" {
//		return false, errors.New(fmt.Sprintf("Failed to delete dns domain record with id '%d', got '%s' as response from the API.", recordId, res))
//	}
//
	return true, nil
}

func (sldr *softLayer_Dns_Domain_Record_Service) GetDomain() (datatypes.SoftLayer_Dns_Domain_Record, error) {
	//TODO need to implement and change return type to datatypes.SoftLayer_Dns_Domain
	return datatypes.SoftLayer_Dns_Domain_Record{}, nil
}

func (sldr *softLayer_Dns_Domain_Record_Service) UpdateObject(recordId int) (bool, error) {
	//	response, err := sldr.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", sldr.GetName(), recordId), "DELETE", new(bytes.Buffer))
	//
	//	if res := string(response[:]); res != "true" {
	//		return false, errors.New(fmt.Sprintf("Failed to delete dns domain record with id '%d', got '%s' as response from the API.", recordId, res))
	//	}
	//
	return true, nil
}