package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Dns_Domain_Resource_Record_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Dns_Domain_Resource_Record_Service(client softlayer.Client) *softLayer_Dns_Domain_Resource_Record_Service {
	return &softLayer_Dns_Domain_Resource_Record_Service{
		client: client,
	}
}

func (sldr *softLayer_Dns_Domain_Resource_Record_Service) GetName() string {
	return "SoftLayer_Dns_Domain_ResourceRecord"
}

func (sldr *softLayer_Dns_Domain_Resource_Record_Service) CreateObject(template datatypes.SoftLayer_Dns_Domain_Resource_Record_Template) (datatypes.SoftLayer_Dns_Domain_Resource_Record, error) {
	parameters := datatypes.SoftLayer_Dns_Domain_Resource_Record_Template_Parameters{
		Parameters: []datatypes.SoftLayer_Dns_Domain_Resource_Record_Template{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Resource_Record{}, err
	}

	response, err := sldr.client.DoRawHttpRequest(fmt.Sprintf("%s/createObject", sldr.getNameByType(template.Type)), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Resource_Record{}, err
	}

	err = sldr.client.CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Resource_Record{}, err
	}

	dns_record := datatypes.SoftLayer_Dns_Domain_Resource_Record{}
	err = json.Unmarshal(response, &dns_record)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Resource_Record{}, err
	}

	return dns_record, nil
}

func (sldr *softLayer_Dns_Domain_Resource_Record_Service) GetObject(id int) (datatypes.SoftLayer_Dns_Domain_Resource_Record, error) {
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
		"service",
		"priority",
		"protocol",
		"port",
		"weight",
	}

	response, err := sldr.client.DoRawHttpRequestWithObjectMask(fmt.Sprintf("%s/%d/getObject.json", sldr.GetName(), id), objectMask, "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Resource_Record{}, err
	}

	err = sldr.client.CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Resource_Record{}, err
	}

	dns_record := datatypes.SoftLayer_Dns_Domain_Resource_Record{}
	err = json.Unmarshal(response, &dns_record)
	if err != nil {
		return datatypes.SoftLayer_Dns_Domain_Resource_Record{}, err
	}

	return dns_record, nil
}

func (sldr *softLayer_Dns_Domain_Resource_Record_Service) DeleteObject(recordId int) (bool, error) {
	response, err := sldr.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", sldr.GetName(), recordId), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete DNS Domain Record with id '%d', got '%s' as response from the API.", recordId, res))
	}

	return true, err
}

func (sldr *softLayer_Dns_Domain_Resource_Record_Service) EditObject(recordId int, template datatypes.SoftLayer_Dns_Domain_Resource_Record) (bool, error) {
	parameters := datatypes.SoftLayer_Dns_Domain_Resource_Record_Parameters{
		Parameters: []datatypes.SoftLayer_Dns_Domain_Resource_Record{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return false, err
	}

	response, err := sldr.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/editObject.json", sldr.getNameByType(template.Type), recordId), "POST", bytes.NewBuffer(requestBody))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to edit DNS Domain Record with id: %d, got '%s' as response from the API.", recordId, res))
	}

	return true, err
}

//Private methods

func (sldr *softLayer_Dns_Domain_Resource_Record_Service) getNameByType(dnsType string) string {
	switch dnsType {
	case "srv":
		// Currently only SRV record type requires additional fields for Create and Update, while all other record types
		// use basic default resource type. Therefore there is no need for now to implement each resource type as separate service
		// https://sldn.softlayer.com/reference/datatypes/SoftLayer_Dns_Domain_ResourceRecord_SrvType
		return "SoftLayer_Dns_Domain_ResourceRecord_SrvType"
	default:
		return "SoftLayer_Dns_Domain_ResourceRecord"
	}
}
