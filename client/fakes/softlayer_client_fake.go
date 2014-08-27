package client_fakes

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

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

	GenerateRequestBodyBuffer *bytes.Buffer
	GenerateRequestBodyError  error

	HasErrorsError error
}

func NewFakeSoftLayerClient(username, apiKey string) *FakeSoftLayerClient {
	pwd, _ := os.Getwd()
	fslc := &FakeSoftLayerClient{
		Username: username,
		ApiKey:   apiKey,

		TemplatePath: filepath.Join(pwd, TEMPLATE_ROOT_PATH),

		SoftLayerServices: map[string]softlayer.Service{},

		DoRawHttpRequestResponse: []byte{},
		DoRawHttpRequestError:    nil,

		GenerateRequestBodyBuffer: new(bytes.Buffer),
		GenerateRequestBodyError:  nil,

		HasErrorsError: nil,
	}

	fslc.initSoftLayerServices()

	return fslc
}

//softlayer.Client interface methods

func (fslc *FakeSoftLayerClient) GetService(serviceName string) (softlayer.Service, error) {
	slService, ok := fslc.SoftLayerServices[serviceName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("softlayer-go does not support service '%s'", serviceName))
	}

	return slService, nil
}

func (fslc *FakeSoftLayerClient) GetSoftLayer_Account_Service() (softlayer.SoftLayer_Account_Service, error) {
	slService, err := fslc.GetService("SoftLayer_Account_Service")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Account_Service), nil
}

func (fslc *FakeSoftLayerClient) GetSoftLayer_Virtual_Guest_Service() (softlayer.SoftLayer_Virtual_Guest_Service, error) {
	slService, err := fslc.GetService("SoftLayer_Virtual_Guest_Service")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Virtual_Guest_Service), nil
}

//Public methods

func (fslc *FakeSoftLayerClient) DoRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, error) {
	return fslc.DoRawHttpRequestResponse, fslc.DoRawHttpRequestError
}

func (fslc *FakeSoftLayerClient) GenerateRequestBody(templateData interface{}) (*bytes.Buffer, error) {
	return fslc.GenerateRequestBodyBuffer, fslc.GenerateRequestBodyError
}

func (fslc *FakeSoftLayerClient) HasErrors(body map[string]interface{}) error {
	return fslc.HasErrorsError
}

//Private methods

func (fslc *FakeSoftLayerClient) initSoftLayerServices() {
	fslc.SoftLayerServices["SoftLayer_Account_Service"] = services.NewSoftLayer_Account_Service(fslc)
	fslc.SoftLayerServices["SoftLayer_Virtual_Guest_Service"] = services.NewSoftLayer_Virtual_Guest_Service(fslc)
}
