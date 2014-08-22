package softlayer

import (
	"bytes"
)

type Client interface {
	GetService(name string) (Service, error)
	GetSoftLayer_Account() (SoftLayer_Account, error)
	GetSoftLayer_Virtual_Guest() (SoftLayer_Virtual_Guest, error)

	DoRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, error)
	GenerateRequestBody(templateData interface{}) (*bytes.Buffer, error)
	HasErrors(body map[string]interface{}) error
}
