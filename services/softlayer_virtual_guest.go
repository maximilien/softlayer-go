package services

import (
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayerVirtualGuest struct {
	client softlayer.Client
}

func NewSoftLayer_Virtual_Guest(client softlayer.Client) *softLayerVirtualGuest {
	return &softLayerVirtualGuest{
		client: client,
	}
}

func (slvg *softLayerVirtualGuest) GetName() string {
	return "SoftLayer_Virtual_Guest"
}

func (slvg *softLayerVirtualGuest) CreateObject(template datatypes.SoftLayer_Virtual_Guest_Template) (datatypes.SoftLayer_Virtual_Guest, error) {
	err := slvg.checkCreateObjectRequiredValues(template)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	return datatypes.SoftLayer_Virtual_Guest{}, errors.New("Implement me!")
}

func (slvg *softLayerVirtualGuest) DeleteObject(template datatypes.SoftLayer_Virtual_Guest_Template) (bool, error) {
	return false, errors.New("Implement me!")
}

//Private methods

func (slvg *softLayerVirtualGuest) checkCreateObjectRequiredValues(template datatypes.SoftLayer_Virtual_Guest_Template) error {
	var err error
	errorMessage, errorTemplate := "", "* %s is required and cannot be empty\n"

	if template.Hostname == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Hostname for the computing instance")
	}	

	if template.Domain == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Domain for the computing instance")
	}	

	if template.StartCpus <= 0 {
		errorMessage += fmt.Sprintf(errorTemplate, "StartCpus: the number of CPU cores to allocate")
	}

	if template.MaxMemory <= 0 {
		errorMessage += fmt.Sprintf(errorTemplate, "MaxMemory: the amount of memory to allocate in megabytes")
	}

	if template.Datacenter.Name == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "datacenter.name: specifies which datacenter the instance is to be provisioned in")
	}

// hourlyBillingFlag
// Specifies the billing type for the instance.
// Required
// Type - boolean
// When true the computing instance will be billed on hourly usage, otherwise it will be billed on a monthly basis.

// localDiskFlag
// Specifies the disk type for the instance.
// Required
// Type - boolean
// When true the disks for the computing instance will be provisioned on the host which it runs, otherwise SAN disks will be provisioned.

// dedicatedAccountHostOnlyFlag
// Specifies whether or not the instance must only run on hosts with instances from the same account
// Optional
// Type - boolean
// Default - false
// When true this flag specifies that a compute instance is to run on hosts that only have guests from the same account.

// operatingSystemReferenceCode
// An identifier for the operating system to provision the computing instance with.
// Conditionally required - Disallowed when blockDeviceTemplateGroup.globalIdentifier is provided, as the template will specify the operating system.
// Type - string
// Notice - Some operating systems are charged based on the value specified in startCpus. The price which is used can be determined by calling generateOrderTemplate with your desired device specifications.
// See getCreateObjectOptions for available options.

// blockDeviceTemplateGroup.globalIdentifier
// A global identifier for the template to be used to provision the computing instance.
// Conditionally required - Disallowed when operatingSystemReferenceCode is provided, as the template will specify the operating system.
// Type - string
// Notice - Some operating systems are charged based on the value specified in startCpus. The price which is used can be determined by calling generateOrderTemplate with your desired device specifications.

// Both public and non-public images may be specified.
// A list of public images may be obtained via a request to getPublicImages.
// A list of non-public images, images owned by an account or specifically shared with an account, may be obtained via a request to getBlockDeviceTemplateGroups.
// Example
// { 
//     "blockDeviceTemplateGroup": { 
//         "globalIdentifier": "07beadaa-1e11-476e-a188-3f7795feb9fb" 
//     } 
// }

	if errorMessage != "" {
		err = errors.New(errorMessage)
	}

	return err
}
