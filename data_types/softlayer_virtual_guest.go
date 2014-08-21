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
