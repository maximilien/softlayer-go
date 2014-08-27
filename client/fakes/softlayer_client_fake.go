package client_fakes

import (
	"bytes"

	services "github.com/maximilien/softlayer-go/services"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const (
	SOFTLAYER_API_URL  = "api.softlayer.com/rest/v3"
	TEMPLATE_ROOT_PATH = "templates"
)

type FakeSoftLayerClient struct {
	Username string
	ApiKey   string

	TemplatePath string

	SoftLayerServices map[string]softlayer.Service

	DoRawHttpRequestResponse []byte
	DoRawHttpRequestError    error

	RequestBodyBuffer *bytes.Buffer
	RequestBodyError  error

	HasErrorsError error
}

func NewFakeSoftLayerClient(username, apiKey string) *FakeSoftLayerClient {
	pwd, _ := os.Getwd()
	return &FakeSoftLayerClient{
		Username: username,
		ApiKey:   apiKey,

		TemplatePath: filepath.Join(pwd, TEMPLATE_ROOT_PATH),

		softLayerServices: map[string]softlayer.Service{},

		DoRawHttpRequestResponse: []byte{},
		DoRawHttpRequestError:    nil,

		GenerateRequestBodyBuffer: &bytes.NewBuffer(),
		GenerateRequestBodyError:  nil,

		HasErrorsError: nil,
	}
}

//softlayer.Client interface methods

func (fslc *FakeSoftLayerClient) GetService(serviceName string) (softlayer.Service, error) {
	slService, ok := fslc.softLayerServices[serviceName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("softlayer-go does not support service '%s'", serviceName))
	}

	return slService, nil
}

func (slc *FakeSoftLayerClient) GetSoftLayer_Account_Service() (softlayer.SoftLayer_Account_Service, error) {
	slService, err := slc.GetService("SoftLayer_Account")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Account_Service), nil
}

func (fslc *FakeSoftLayerClient) GetSoftLayer_Virtual_Guest_Service() (softlayer.SoftLayer_Virtual_Guest_Service, error) {
	slService, err := slc.GetService("SoftLayer_Virtual_Guest")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Virtual_Guest_Service), nil
}

//Public methods

func (fslc *FakeSoftLayerClient) DoRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, error) {
	return slfc.DoRawHttpRequestResponse, slfc.DoRawHttpRequestError
}

func (fslc *FakeSoftLayerClient) GenerateRequestBody(templateData interface{}) (*bytes.Buffer, error) {
	return fslc.GenerateRequestBodyBuffer, fslc.GenerateRequestBodyError
}

func (fslc *FakeSoftLayerClient) HasErrors(body map[string]interface{}) error {
	return fslc.HasErrorsError
}
