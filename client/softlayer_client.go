package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	services "github.com/maximilien/softlayer-go/services"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const (
	SOFTLAYER_API_URL  = "api.softlayer.com/rest/v3"
	TEMPLATE_ROOT_PATH = "templates"
)

type softLayerClient struct {
	username string
	apiKey   string

	templatePath string

	httpClient *http.Client

	softLayerServices map[string]softlayer.Service
}

func NewSoftLayerClient(username, apiKey string) *softLayerClient {
	pwd, _ := os.Getwd()
	slc := &softLayerClient{
		username: username,
		apiKey:   apiKey,

		templatePath: filepath.Join(pwd, TEMPLATE_ROOT_PATH),

		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},

		softLayerServices: map[string]softlayer.Service{},
	}

	slc.initSoftLayerServices()

	return slc
}

//Client interface methods

func (slc *softLayerClient) GetService(serviceName string) (softlayer.Service, error) {
	slService, ok := slc.softLayerServices[serviceName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("softlayer-go does not support service '%s'", serviceName))
	}

	return slService, nil
}

func (slc *softLayerClient) GetSoftLayer_Account() (softlayer.SoftLayer_Account, error) {
	slService, err := slc.GetService("SoftLayer_Account")
	if err != nil {
		return nil, err
	}

	return slService.(softlayer.SoftLayer_Account), nil
}

//Private methods

func (slc *softLayerClient) initSoftLayerServices() {
	slc.softLayerServices["SoftLayer_Account"] = services.NewSoftLayer_Account(slc)
}

func (slc *softLayerClient) generateRequestBody(templateData interface{}) (*bytes.Buffer, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	bodyTemplate := template.Must(template.ParseFiles(filepath.Join(cwd, slc.templatePath)))
	body := new(bytes.Buffer)
	bodyTemplate.Execute(body, templateData)

	log.Printf("Generated request body %s", body)

	return body, nil
}

func (slc *softLayerClient) hasErrors(body map[string]interface{}) error {
	if errString, ok := body["error"]; !ok {
		return nil
	} else {
		return errors.New(errString.(string))
	}
}

func (slc *softLayerClient) doRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, error) {
	url := fmt.Sprintf("https://%s:%s@%s/%s", slc.username, slc.apiKey, SOFTLAYER_API_URL, path)
	log.Printf("Sending new request to softlayer: %s", url)

	var lastResponse http.Response
	switch requestType {
	case "POST", "DELETE":
		req, err := http.NewRequest(requestType, url, requestBody)

		if err != nil {
			return nil, err
		}
		resp, err := slc.httpClient.Do(req)

		if err != nil {
			return nil, err
		} else {
			lastResponse = *resp
		}
	case "GET":
		resp, err := http.Get(url)

		if err != nil {
			return nil, err
		} else {
			lastResponse = *resp
		}
	default:
		return nil, errors.New(fmt.Sprintf("Undefined request type '%s', only GET/POST/DELETE are available!", requestType))
	}

	responseBody, err := ioutil.ReadAll(lastResponse.Body)
	lastResponse.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Printf("Received response from SoftLayer: %s", responseBody)

	return responseBody, nil
}

func (slc *softLayerClient) doHttpRequest(path string, requestType string, requestBody *bytes.Buffer) (map[string]interface{}, error) {
	responseBody, err := slc.doRawHttpRequest(path, requestType, requestBody)
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to get proper HTTP response from SoftLayer API"))
		return nil, err
	}

	var decodedResponse map[string]interface{}
	err = json.Unmarshal(responseBody, &decodedResponse)
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to decode JSON response from SoftLayer: %s | %s", responseBody, err))
		return nil, err
	}

	return decodedResponse, nil
}
