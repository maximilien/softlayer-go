package softlayer

import (
	"bytes"
)

type Client interface {
	GetService(name string) (Service, error)

	GetSoftLayer_Account_Service() (SoftLayer_Account_Service, error)
	GetSoftLayer_Virtual_Guest_Service() (SoftLayer_Virtual_Guest_Service, error)
	GetSoftLayer_Ssh_Key_Service() (SoftLayer_Ssh_Key_Service, error)

	DoRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, error)
	GenerateRequestBody(templateData interface{}) (*bytes.Buffer, error)
	HasErrors(body map[string]interface{}) error

	CheckForHttpResponseErrors(data []byte) error
}
