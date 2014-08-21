package data_types

import (
	"time"
)

type SoftLayer_Network_Storage struct {
	AccountId         int       `json:"accountId"`
	CapacityGb        int       `json:"capacityGb"`
	CreateDate        time.Time `json:"createDate"`
	GuestId           int       `json:"guestId"`
	HardwareId        int       `json:"hardwareId"`
	HostId            int       `json:"hostId"`
	Id                int       `json:"id"`
	NasType           string    `json:"nasType"`
	Notes             string    `json:"notes"`
	Password          string    `json:"password"`
	ServiceProviderId int       `json:"serviceProviderId"`
	UpgradableFlag    bool      `json:"upgradableFlag"`
	Username          string    `json:"username"`
}
