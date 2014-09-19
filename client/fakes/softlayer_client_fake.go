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

	HasErrorsError, CheckForHttpResponseError error
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

		HasErrorsError:            nil,
		CheckForHttpResponseError: nil,
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
	slService, err := fslc.GetService("SoftLayer_Account")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Account_Service), nil
}

func (fslc *FakeSoftLayerClient) GetSoftLayer_Virtual_Guest_Service() (softlayer.SoftLayer_Virtual_Guest_Service, error) {
	slService, err := fslc.GetService("SoftLayer_Virtual_Guest")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Virtual_Guest_Service), nil
}

func (fslc *FakeSoftLayerClient) GetSoftLayer_Virtual_Disk_Image_Service() (softlayer.SoftLayer_Virtual_Disk_Image_Service, error) {
	slService, err := fslc.GetService("SoftLayer_Virtual_Disk_Image")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Virtual_Disk_Image_Service), nil
}

func (fslc *FakeSoftLayerClient) GetSoftLayer_Security_Ssh_Key_Service() (softlayer.SoftLayer_Security_Ssh_Key_Service, error) {
	slService, err := fslc.GetService("SoftLayer_Ssh_Key")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Security_Ssh_Key_Service), nil
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

func (fslc *FakeSoftLayerClient) CheckForHttpResponseErrors(data []byte) error {
	return fslc.CheckForHttpResponseError
}

//Private methods

func (fslc *FakeSoftLayerClient) initSoftLayerServices() {
	fslc.SoftLayerServices["SoftLayer_Account"] = services.NewSoftLayer_Account_Service(fslc)
	fslc.SoftLayerServices["SoftLayer_Virtual_Guest"] = services.NewSoftLayer_Virtual_Guest_Service(fslc)
	fslc.SoftLayerServices["SoftLayer_Virtual_Disk_Image"] = services.NewSoftLayer_Virtual_Disk_Image_Service(fslc)
	fslc.SoftLayerServices["SoftLayer_Ssh_Key"] = services.NewSoftLayer_Ssh_Key_Service(fslc)
}
