package data_types

import (
	"time"
)

type SoftLayer_Virtual_Guest struct {
	AccountId                    int        `json:"accountId"`
	CreateDate                   *time.Time `json:"createDate"`
	DedicatedAccountHostOnlyFlag bool       `json:"dedicatedAccountHostOnlyFlag"`
	Domain                       string     `json:"domain"`
	FullyQualifiedDomainName     string     `json:"fullyQualifiedDomainName"`
	Hostname                     string     `json:"hostname"`
	Id                           int        `json:"id"`
	LastPowerStateId             int        `json:"lastPowerStateId"`
	LastVerifiedDate             *time.Time `json:"lastVerifiedDate"`
	MaxCpu                       int        `json:"maxCpu"`
	MaxCpuUnits                  string     `json:"maxCpuUnits"`
	MaxMemory                    int        `json:"maxMemory"`
	MetricPollDate               *time.Time `json:"metricPollDate"`
	ModifyDate                   *time.Time `json:"modifyDate"`
	Notes                        string     `json:"notes"`
	PostInstallScriptUri         string     `json:"postInstallScriptUri"`
	PrivateNetworkOnlyFlag       bool       `json:"privateNetworkOnlyFlag"`
	StartCpus                    int        `json:"startCpus"`
	StatusId                     int        `json:"statusId"`
	Uuid                         string     `json:"uuid"`
}

type SoftLayer_Virtual_Guest_Template_Paramaters struct {
	Parameters []SoftLayer_Virtual_Guest_Template `json:"parameters"`
}

type SoftLayer_Virtual_Guest_Template struct {
	//Required
	Hostname          string     `json:"hostname"`
	Domain            string     `json:"domain"`
	StartCpus         int        `json:"startCpus"`
	MaxMemory         int        `json:"maxMemory"`
	Datacenter        Datacenter `json:"datacenter"`
	HourlyBillingFlag bool       `json:"hourlyBillingFlag"`
	LocalDiskFlag     bool       `json:"localDiskFlag"`

	//Conditionally required
	OperatingSystemReferenceCode string                    `json:"operatingSystemReferenceCode"`
	BlockDeviceTemplateGroup     *BlockDeviceTemplateGroup `json:"blockDeviceTemplateGroup,omitempty"`

	//Optional
	DedicatedAccountHostOnlyFlag   bool                            `json:"dedicatedAccountHostOnlyFlag,omitempty"`
	NetworkComponents              []NetworkComponents             `json:"networkComponents,omitempty"`
	PrivateNetworkOnlyFlag         bool                            `json:"privateNetworkOnlyFlag,omitempty"`
	PrimaryNetworkComponent        *PrimaryNetworkComponent        `json:"primaryNetworkComponent,omitempty"`
	PrimaryBackendNetworkComponent *PrimaryBackendNetworkComponent `json:"primaryBackendNetworkComponent,omitempty"`

	BlockDevices []BlockDevice `json:"blockDevices,omitempty"`
	UserData     []UserData    `json:"userData,omitempty"`
	SshKeys      []SshKey      `json:"sshKeys,omitempty"`

	PostInstallScriptUri string `json:"postInstallScriptUri,omitempty"`
}

type Datacenter struct {
	//Required
	Name string `json:"name"`
}

type BlockDeviceTemplateGroup struct {
	//Required
	GlobalIdentifier string `json:"globalIdentifier,omitempty"`
}

type NetworkComponents struct {
	//Required, defaults to 10
	MaxSpeed int `json:"maxSpeed,omitempty"`
}

type NetworkVlan struct {
	//Required
	Id int `json:"id,omitempty"`
}

type PrimaryNetworkComponent struct {
	//Required
	NetworkVlan NetworkVlan `json:"networkVlan,omitempty"`
}

type PrimaryBackendNetworkComponent struct {
	//Required
	NetworkVlan NetworkVlan `json:"networkVlan,omitempty"`
}

type DiskImage struct {
	//Required
	capacity int `json:"capacity,omitempty"`
}

type BlockDevice struct {
	//Required
	Device    string    `json:"device,omitempty"`
	DiskImage DiskImage `json:"diskImage,omitempty"`
}

type UserData struct {
	//Required
	Value string `json:"value,omitempty"`
}

type SshKey struct {
	//Required
	Id int `json:"id,omitempty"`
}
