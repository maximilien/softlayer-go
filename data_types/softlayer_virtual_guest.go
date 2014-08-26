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

type SoftLayer_Virtual_Guest_Template struct {
	//Required
	Hostname                     string     `json:"hostname"`
	Domain                       string     `json:"domain"`
	StartCpus                    int        `json:"startCpus"`
	MaxMemory                    int        `json:"maxMemory"`
	Datacenter                   Datacenter `json:"datacenter"`
	HourlyBillingFlag            bool       `json:"hourlyBillingFlag"`
	LocalDiskFlag                bool       `json:"localDiskFlag"`
	DedicatedAccountHostOnlyFlag bool       `json:"dedicatedAccountHostOnlyFlag"`

	//Conditionally required
	OperatingSystemReferenceCode string                   `json:"operatingSystemReferenceCode"`
	BlockDeviceTemplateGroup     BlockDeviceTemplateGroup `json:"blockDeviceTemplateGroup"`

	//Optional
	NetworkComponents              []NetworkComponents            `json:"networkComponents"`
	PrivateNetworkOnlyFlag         bool                           `json:"privateNetworkOnlyFlag"`
	PrimaryNetworkComponent        PrimaryNetworkComponent        `json:"primaryNetworkComponent"`
	PrimaryBackendNetworkComponent PrimaryBackendNetworkComponent `json:"primaryBackendNetworkComponent"`

	BlockDevices []BlockDevice `json:"blockDevices"`
	UserData     []UserData    `json:"userData"`
	SshKeys      []SshKey      `json:"sshKeys"`

	PostInstallScriptUri string `json:"postInstallScriptUri"`
}

type Datacenter struct {
	//Required
	Name string `json:"name"`
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

type DiskImage struct {
	//Required
	capacity int `json:"capacity"`
}

type BlockDevice struct {
	//Required
	Device    string    `json:"device"`
	DiskImage DiskImage `json:"diskImage"`
}

type UserData struct {
	//Required
	Value string `json:"value"`
}

type SshKey struct {
	//Required
	Id int `json:"id"`
}
