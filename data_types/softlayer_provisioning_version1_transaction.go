package data_types

import (
	"time"
)

type SoftLayer_Provisioning_Version1_Transaction struct {
	CreateDate       *time.Time `json:"createDate"`
	ElapsedSeconds   int        `json:"elapsedSeconds"`
	GuestId          int        `json:"guestId"`
	HardwareId       int        `json:"hardwareId"`
	Id               int        `json:"id"`
	ModifyDate       *time.Time `json:"modifyDate"`
	StatusChangeDate *time.Time `json:"statusChangeDate"`
}
