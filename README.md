softlayer-go [![Build Status](https://travis-ci.org/maximilien/softlayer-go.svg?branch=master)](https://travis-ci.org/maximilien/softlayer-go#)
============

An *incomplete* SoftLayer (SL) client API written in Go language

The best way to get started would be to look at the integration tests for creating a virtual guest. Here is a snippet of what is needed.

```go
//Add necessary imports, e.g., os, slclient, datatypes
// "os"
// slclient "github.com/maximilien/softlayer-go/client"
// datatypes "github.com/maximilien/softlayer-go/data_types"

//Access SoftLayer username and API key from environment variable or hardcode here
username := os.Getenv("SL_USERNAME")
apiKey := os.Getenv("SL_API_KEY")
	
//Create a softLayer-go client
client := slclient.NewSoftLayerClient(username, apiKey)

//Get the SoftLayer account service object
accountService, err := client.GetSoftLayer_Account_Service()
if err != nil {
  return err
}

//Create a template for the virtual guest (changing properties as needed)
virtualGuestTemplate := datatypes.SoftLayer_Virtual_Guest_Template{
  Hostname:  "some-hostname",
	Domain:    "softlayergo.com",
	StartCpus: 1,
	MaxMemory: 1024,
	Datacenter: datatypes.Datacenter{
		Name: "ams01",
	},
	SshKeys:                      [],  //or get the necessary keys and add here
	HourlyBillingFlag:            true,
	LocalDiskFlag:                true,
	OperatingSystemReferenceCode: "UBUNTU_LATEST",
}
	
//Get the SoftLayer virtual guest service
virtualGuestService, err := client.GetSoftLayer_Virtual_Guest_Service()
if err != nil {
  return err
}
	
//Create the virtual guest with the service
virtualGuest, err := virtualGuestService.CreateObject(virtualGuestTemplate)
if err != nil {
	return err
}
	
//Use the virtualGuest or other services...
```

**NOTE**: this client is created to support the [bosh-softlayer-cpi](https://github.com/maximilien/bosh-softlayer-cpi) project and only implements the portion of the SL APIs needed to complete the implementation of the BOSH CPI. You are welcome to use it in your own projects and as you do if you find areas we have not yet implemented but that you need, please submit [Pull Requests](https://help.github.com/articles/using-pull-requests/) or engage with us in discussions.
