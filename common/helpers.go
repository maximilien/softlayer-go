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
	DATACENTER_TYPE_NAME = "SoftLayer_Location_Datacenter"
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
