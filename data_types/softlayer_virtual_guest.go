package data_types

import (
	"time"
)

type SoftLayer_Virtual_Guest struct {
	AccountId                    int       `json:"accountId"`
	CreateDate                   time.Time `json:"createDate"`
	DedicatedAccountHostOnlyFlag bool      `json:"dedicatedAccountHostOnlyFlag"`
	Domain                       string    `json:"domain"`
	FullyQualifiedDomainName     string    `json:"fullyQualifiedDomainName"`
	Hostname                     string    `json:"hostname"`
	Id                           int       `json:"id"`
	LastPowerStateId             int       `json:"lastPowerStateId"`
	//LastVerifiedDate             time.Time 	`json:"lastVerifiedDate"`
	MaxCpu      int    `json:"maxCpu"`
	MaxCpuUnits string `json:"maxCpuUnits"`
	MaxMemory   int    `json:"maxMemory"`
	//MetricPollDate               time.Time  `json:"metricPollDate"`
	ModifyDate             time.Time `json:"modifyDate"`
	Notes                  string    `json:"notes"`
	PostInstallScriptUri   string    `json:"postInstallScriptUri"`
	PrivateNetworkOnlyFlag bool      `json:"privateNetworkOnlyFlag"`
	StartCpus              int       `json:"startCpus"`
	StatusId               int       `json:"statusId"`
	Uuid                   string    `json:"uuid"`
}

type Datacenter struct {
	//Required
	Name string 	`json:"name"`	
}

type BlockDeviceTemplateGroup struct {
	//Required
	GlobalIdentifier string `json:"globalIdentifier"`
}

type NetworkComponents struct {
	//Required, defaults to 10
	MaxSpeed int `json:"maxSpeed"`
}

type NetworkVlan struct {
	//Required
	Id int `json:"id"`
}

type PrimaryNetworkComponent struct {
	//Required
	NetworkVlan NetworkVlan `json:"networkVlan"`
}

type PrimaryBackendNetworkComponent struct {
	//Required
	NetworkVlan NetworkVlan `json:"networkVlan"`
}

type SoftLayer_Virtual_Guest_Template struct {
	//Required
	Hostname string `json:"hostname"`			
	Domain string `json:"domain"`				
	StartCpus int `json:"startCpus"`			
	MaxMemory int `json:"maxMemory"`			
	Datacenter Datacenter `json:"datacenter"`	
	HourlyBillingFlag bool `json:"hourlyBillingFlag"`
	LocalDiskFlag bool `json:"localDiskFlag"`
	DedicatedAccountHostOnlyFlag bool `json:"dedicatedAccountHostOnlyFlag"`

	//Conditionally required
	OperatingSystemReferenceCode string `json:"operatingSystemReferenceCode"`
	BlockDeviceTemplateGroup BlockDeviceTemplateGroup `json:"blockDeviceTemplateGroup"`

	//Optional
	NetworkComponents []NetworkComponents `json:"networkComponents"`
	PrivateNetworkOnlyFlag bool `json:"privateNetworkOnlyFlag"`
	PrimaryNetworkComponent PrimaryNetworkComponent `json:"primaryNetworkComponent"`
	PrimaryBackendNetworkComponent PrimaryBackendNetworkComponent `json:"primaryBackendNetworkComponent"`

	PostInstallScriptUri string `json:"postInstallScriptUri"`
}


// blockDevices
// Block device and disk image settings for the computing instance
// Optional
// Type - array of [[SoftLayer_Virtual_Guest_Block_Device (type)|SoftLayer_Virtual_Guest_Block_Device]
// Default - The smallest available capacity for the primary disk will be used. If an image template is specified the disk capacity will be be provided by the template.
// Description - The blockDevices property is an array of block device structures.
// Each block device must specify the device property along with the diskImage property, which is a disk image structure with the capacity property set.
// The device number '1' is reserved for the SWAP disk attached to the computing instance.
// See getCreateObjectOptions for available options.
// Example
// { 
//     "blockDevices": [ 
//         { 
//             "device": "0", 
//             "diskImage": { 
//                 "capacity": 100 
//             } 
//         } 
//     ], 
//     "localDiskFlag": true 
// }

// userData.value
// Arbitrary data to be made available to the computing instance.
// Optional
// Type - string
// Description - The userData property is an array with a single attribute structure with the value property set to an arbitrary value.
// This value can be retrieved via the getUserMetadata method from a request originating from the computing instance. This is primarily useful for providing data to software that may be on the instance and configured to execute upon first boot.
// Example
// { 
//     "userData": [ 
//         { 
//             "value": "someValue" 
//         } 
//     ] 
// }

// sshKeys
// SSH keys to install on the computing instance upon provisioning.
// Optional
// Type - array of SoftLayer_Security_Ssh_Key
// Description - The sshKeys property is an array of SSH Key structures with the id property set to the value of an existing SSH key.
// To create a new SSH key, call createObject on the SoftLayer_Security_Ssh_Key service.
// To obtain a list of existing SSH keys, call getSshKeys on the SoftLayer_Account service.
// Example
// { 
//     "sshKeys": [ 
//         { 
//             "id": 123 
//         } 
//     ] 
// }