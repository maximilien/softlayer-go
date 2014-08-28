package data_types

import (
	"time"
)

type SoftLayer_Ssh_Key struct {
	CreateDate  *time.Time `json:"createDate"`
	Fingerprint string     `json:"fingerprint"`
	Id          int        `json:"id"`
	Key         string     `json:"key"`
	Label       string     `json:"label"`
	ModifyDate  *time.Time `json:"modifyDate"`
	Notes       string     `json:"notes"`
}
