package softlayer

type Client interface {
	GetService(name string) (Service, error)
	GetSoftLayer_Account() (SoftLayer_Account, error)
}
