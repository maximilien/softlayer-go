package services

import (
	datatypes "github.com/maximilien/softlayer-go"
)

type SoftLayer_Account interface {
	getVirtualGuests() []datatypes.SoftLayer_Virtual_Guest
}
