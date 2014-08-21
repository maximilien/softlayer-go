package data_types

import (
	"time"
)

type SoftLayer_Virtual_Guest struct {
	AccountId                    int
	CreateDate                   time.Time
	DedicatedAccountHostOnlyFlag bool
	Domain                       string
	FullyQualifiedDomainName     string
	Hostname                     string
	Id                           int
	LastPowerStateId             int
	LastVerifiedDate             time.Time
	MaxCpu                       int
	MaxCpuUnits                  string
	MaxMemory                    int
	MetricPollDate               time.Time
	ModifyDate                   time.Time
	Notes                        string
	PostInstallScriptUri         string
	PrivateNetworkOnlyFlag       bool
	StartCpus                    int
	StatusId                     int
	//SupplementalCreateObjectOptions SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions
	Uuid string
}
