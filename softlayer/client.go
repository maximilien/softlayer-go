package softlayer

type Client interface {
	GetService(name string) (Service, error)
}
