package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const (
	ROOT_USER_NAME               = "root"
	EPHEMERAL_DISK_CATEGORY_CODE = "guest_disk1"
)

type softLayer_Virtual_Guest_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Virtual_Guest_Service(client softlayer.Client) *softLayer_Virtual_Guest_Service {
	return &softLayer_Virtual_Guest_Service{
		client: client,
	}
}

func (slvgs *softLayer_Virtual_Guest_Service) GetName() string {
	return "SoftLayer_Virtual_Guest"
}

func (slvgs *softLayer_Virtual_Guest_Service) CreateObject(template datatypes.SoftLayer_Virtual_Guest_Template) (datatypes.SoftLayer_Virtual_Guest, error) {
	err := slvgs.checkCreateObjectRequiredValues(template)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	parameters := datatypes.SoftLayer_Virtual_Guest_Template_Parameters{
		Parameters: []datatypes.SoftLayer_Virtual_Guest_Template{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s.json", slvgs.GetName()), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	err = slvgs.client.CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	softLayer_Virtual_Guest := datatypes.SoftLayer_Virtual_Guest{}
	err = json.Unmarshal(response, &softLayer_Virtual_Guest)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	return softLayer_Virtual_Guest, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetObject(instanceId int) (datatypes.SoftLayer_Virtual_Guest, error) {

	objectMask := []string{
		"accountId",
		"createDate",
		"dedicatedAccountHostOnlyFlag",
		"domain",
		"fullyQualifiedDomainName",
		"hostname",
		"id",
		"lastPowerStateId",
		"lastVerifiedDate",
		"maxCpu",
		"maxCpuUnits",
		"maxMemory",
		"metricPollDate",
		"modifyDate",
		"notes",
		"postInstallScriptUri",
		"privateNetworkOnlyFlag",
		"startCpus",
		"statusId",
		"uuid",

		"globalIdentifier",
		"managedResourceFlag",
		"primaryBackendIpAddress",
		"primaryIpAddress",

		"location.id",
		"datacenter.id",
		"operatingSystem.passwords.password",
		"operatingSystem.passwords.username",
	}

	response, err := slvgs.client.DoRawHttpRequestWithObjectMask(fmt.Sprintf("%s/%d/getObject.json", slvgs.GetName(), instanceId), objectMask, "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	virtualGuest := datatypes.SoftLayer_Virtual_Guest{}
	err = json.Unmarshal(response, &virtualGuest)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	return virtualGuest, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) EditObject(instanceId int, template datatypes.SoftLayer_Virtual_Guest) (bool, error) {
	parameters := datatypes.SoftLayer_Virtual_Guest_Parameters{
		Parameters: []datatypes.SoftLayer_Virtual_Guest{template},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return false, err
	}

	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/editObject.json", slvgs.GetName(), instanceId), "POST", bytes.NewBuffer(requestBody))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to edit virtual guest with id: %d, got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) DeleteObject(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", slvgs.GetName(), instanceId), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) GetPowerState(instanceId int) (datatypes.SoftLayer_Virtual_Guest_Power_State, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getPowerState.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Power_State{}, err
	}

	vgPowerState := datatypes.SoftLayer_Virtual_Guest_Power_State{}
	err = json.Unmarshal(response, &vgPowerState)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Power_State{}, err
	}

	return vgPowerState, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetPrimaryIpAddress(instanceId int) (string, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getPrimaryIpAddress.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return "", err
	}

	vgPrimaryIpAddress := strings.TrimSpace(string(response))
	if vgPrimaryIpAddress == "" {
		return "", errors.New(fmt.Sprintf("Failed to get primary IP address for instance with id '%d', got '%s' as response from the API.", instanceId, response))
	}

	return vgPrimaryIpAddress, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetActiveTransaction(instanceId int) (datatypes.SoftLayer_Provisioning_Version1_Transaction, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getActiveTransaction.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	activeTransaction := datatypes.SoftLayer_Provisioning_Version1_Transaction{}
	err = json.Unmarshal(response, &activeTransaction)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	return activeTransaction, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetActiveTransactions(instanceId int) ([]datatypes.SoftLayer_Provisioning_Version1_Transaction, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getActiveTransactions.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	activeTransactions := []datatypes.SoftLayer_Provisioning_Version1_Transaction{}
	err = json.Unmarshal(response, &activeTransactions)
	if err != nil {
		return []datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	return activeTransactions, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetSshKeys(instanceId int) ([]datatypes.SoftLayer_Security_Ssh_Key, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getSshKeys.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	sshKeys := []datatypes.SoftLayer_Security_Ssh_Key{}
	err = json.Unmarshal(response, &sshKeys)
	if err != nil {
		return []datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	return sshKeys, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) PowerCycle(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/powerCycle.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to power cycle instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) PowerOff(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/powerOff.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to power off instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) PowerOffSoft(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/powerOffSoft.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to power off soft instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) PowerOn(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/powerOn.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to power on instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) RebootDefault(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/rebootDefault.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to default reboot instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) RebootSoft(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/rebootSoft.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to soft reboot instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) RebootHard(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/rebootHard.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to hard reboot instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) SetMetadata(instanceId int, metadata string) (bool, error) {
	dataBytes := []byte(metadata)
	base64EncodedMetadata := base64.StdEncoding.EncodeToString(dataBytes)

	parameters := datatypes.SoftLayer_SetUserMetadata_Parameters{
		Parameters: []datatypes.UserMetadataArray{
			[]datatypes.UserMetadata{datatypes.UserMetadata(base64EncodedMetadata)},
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return false, err
	}

	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/setUserMetadata.json", slvgs.GetName(), instanceId), "POST", bytes.NewBuffer(requestBody))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to setUserMetadata for instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) ConfigureMetadataDisk(instanceId int) (datatypes.SoftLayer_Provisioning_Version1_Transaction, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/configureMetadataDisk.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	transaction := datatypes.SoftLayer_Provisioning_Version1_Transaction{}
	err = json.Unmarshal(response, &transaction)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	return transaction, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) AttachIscsiVolume(instanceId int, volumeId int) (string, error) {

	virtualGuest, err := slvgs.GetObject(instanceId)
	if err != nil {
		return "", err
	}

	networkStorageService, err := slvgs.client.GetSoftLayer_Network_Storage_Service()
	if err != nil {
		return "", err
	}

	volume, err := networkStorageService.GetIscsiVolume(volumeId)
	if err != nil {
		return "", err
	}

	deviceName, err := slvgs.attachVolumeBasedOnShellScript(virtualGuest, volume)
	if err != nil {
		return "", err
	}

	return deviceName, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) DetachIscsiVolume(instanceId int, volumeId int) error {
	virtualGuest, err := slvgs.GetObject(instanceId)
	if err != nil {
		return err
	}

	networkStorageService, err := slvgs.client.GetSoftLayer_Network_Storage_Service()
	if err != nil {
		return err
	}
	volume, err := networkStorageService.GetIscsiVolume(volumeId)
	if err != nil {
		return err
	}

	return slvgs.detachVolumeBasedOnShellScript(virtualGuest, volume)
}

func (slvgs *softLayer_Virtual_Guest_Service) GetUserData(instanceId int) ([]datatypes.SoftLayer_Virtual_Guest_Attribute, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getUserData.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Guest_Attribute{}, err
	}

	attributes := []datatypes.SoftLayer_Virtual_Guest_Attribute{}
	err = json.Unmarshal(response, &attributes)
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Guest_Attribute{}, err
	}

	return attributes, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) IsPingable(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/isPingable.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return false, err
	}

	res := string(response)

	if res == "true" {
		return true, nil
	}

	if res == "false" {
		return false, nil
	}

	return false, errors.New(fmt.Sprintf("Failed to checking that virtual guest is pingable for instance with id '%d', got '%s' as response from the API.", instanceId, res))
}

func (slvgs *softLayer_Virtual_Guest_Service) AttachEphemeralDisk(instanceId int, diskSize int) error {
	diskItemPrice, err := slvgs.findUpgradeItemPriceForEphemeralDisk(instanceId, diskSize)
	if err != nil {
		return err
	}

	service, err := slvgs.client.GetSoftLayer_Product_Order_Service()
	if err != nil {
		return err
	}

	order := datatypes.SoftLayer_Product_Order{
		VirtualGuests: []datatypes.VirtualGuest{
			datatypes.VirtualGuest{
				Id: instanceId,
			},
		},
		Prices: []datatypes.SoftLayer_Item_Price{
			datatypes.SoftLayer_Item_Price{
				Id: diskItemPrice.Id,
				Categories: []datatypes.Category{
					datatypes.Category{
						CategoryCode: EPHEMERAL_DISK_CATEGORY_CODE,
					},
				},
			},
		},
		ComplexType: "SoftLayer_Container_Product_Order_Virtual_Guest_Upgrade",
		Properties: []datatypes.Property{
			datatypes.Property{
				Name:  "MAINTENANCE_WINDOW",
				Value: time.Now().UTC().Format(time.RFC3339),
			},
			datatypes.Property{
				Name:  "NOTE_GENERAL",
				Value: "addingdisks",
			},
		},
	}

	_, err = service.PlaceOrder(order)

	return err
}

func (slvgs *softLayer_Virtual_Guest_Service) GetUpgradeItemPrices(instanceId int) ([]datatypes.SoftLayer_Item_Price, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getUpgradeItemPrices.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Item_Price{}, err
	}

	itemPrices := []datatypes.SoftLayer_Item_Price{}
	err = json.Unmarshal(response, &itemPrices)
	if err != nil {
		return []datatypes.SoftLayer_Item_Price{}, err
	}

	return itemPrices, nil
}

//Private methods

func (slvgs *softLayer_Virtual_Guest_Service) attachVolumeBasedOnShellScript(virtualGuest datatypes.SoftLayer_Virtual_Guest, volume datatypes.SoftLayer_Network_Storage) (string, error) {
	command := fmt.Sprintf(`
		export PATH=/etc/init.d:$PATH
		cp /etc/iscsi/iscsid.conf{,.save}
		sed '/^node.startup/s/^.*/node.startup = automatic/' -i /etc/iscsi/iscsid.conf
		sed '/^#node.session.auth.authmethod/s/#//' -i /etc/iscsi/iscsid.conf
		sed '/^#node.session.auth.username / {s/#//; s/ username/ %s/}' -i /etc/iscsi/iscsid.conf
		sed '/^#node.session.auth.password / {s/#//; s/ password/ %s/}' -i /etc/iscsi/iscsid.conf
		sed '/^#discovery.sendtargets.auth.username / {s/#//; s/ username/ %s/}' -i /etc/iscsi/iscsid.conf
		sed '/^#discovery.sendtargets.auth.password / {s/#//; s/ password/ %s/}' -i /etc/iscsi/iscsid.conf
		open-iscsi restart
		rm -r /etc/iscsi/send_targets
		open-iscsi stop
		open-iscsi start
		iscsiadm -m discovery -t sendtargets -p %s
		open-iscsi restart`,
		volume.Username,
		volume.Password,
		volume.Username,
		volume.Password,
		volume.ServiceResourceBackendIpAddress,
	)

	_, err := slvgs.client.ExecShellCommand(ROOT_USER_NAME, slvgs.getRootPassword(virtualGuest), virtualGuest.PrimaryIpAddress, command)
	if err != nil {
		return "", err
	}

	_, deviceName, err := slvgs.findIscsiDeviceNameBasedOnShellScript(virtualGuest, volume)
	if err != nil {
		return "", err
	}

	return deviceName, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) detachVolumeBasedOnShellScript(virtualGuest datatypes.SoftLayer_Virtual_Guest, volume datatypes.SoftLayer_Network_Storage) error {
	targetName, _, err := slvgs.findIscsiDeviceNameBasedOnShellScript(virtualGuest, volume)
	command := fmt.Sprintf(`
		iscsiadm -m node -T %s -u
		iscsiadm -m node -o delete -T %s`,
		targetName,
		targetName,
	)

	_, err = slvgs.client.ExecShellCommand(ROOT_USER_NAME, slvgs.getRootPassword(virtualGuest), virtualGuest.PrimaryIpAddress, command)
	if err != nil {
		return err
	}

	return err
}

func (slvgs *softLayer_Virtual_Guest_Service) findIscsiDeviceNameBasedOnShellScript(virtualGuest datatypes.SoftLayer_Virtual_Guest, volume datatypes.SoftLayer_Network_Storage) (targetName string, deviceName string, err error) {
	command := `
		sleep 1
		iscsiadm -m session -P3 | sed -n  "/Target:/s/Target: //p; /Attached scsi disk /{ s/Attached scsi disk //; s/State:.*//p}"`

	output, err := slvgs.client.ExecShellCommand(ROOT_USER_NAME, slvgs.getRootPassword(virtualGuest), virtualGuest.PrimaryIpAddress, command)
	if err != nil {
		return "", "", err
	}

	lines := strings.Split(strings.Trim(output, "\n"), "\n")

	for i := 0; i < len(lines); i += 2 {
		if strings.Contains(lines[i], strings.ToLower(volume.Username)) {
			return strings.Trim(lines[i], "\t"), strings.Trim(lines[i+1], "\t"), nil
		}
	}

	return "", "", errors.New(fmt.Sprintf("Can not find matched iSCSI device for user name: %s", volume.Username))
}

func (slvgs *softLayer_Virtual_Guest_Service) checkCreateObjectRequiredValues(template datatypes.SoftLayer_Virtual_Guest_Template) error {
	var err error
	errorMessage, errorTemplate := "", "* %s is required and cannot be empty\n"

	if template.Hostname == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Hostname for the computing instance")
	}

	if template.Domain == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Domain for the computing instance")
	}

	if template.StartCpus <= 0 {
		errorMessage += fmt.Sprintf(errorTemplate, "StartCpus: the number of CPU cores to allocate")
	}

	if template.MaxMemory <= 0 {
		errorMessage += fmt.Sprintf(errorTemplate, "MaxMemory: the amount of memory to allocate in megabytes")
	}

	for _, device := range template.BlockDevices {
		if device.DiskImage.Capacity <= 0 {
			errorMessage += fmt.Sprintf("Disk size must be positive number, the size of block device %s is set to be %dGB.", device.Device, device.DiskImage.Capacity)
		}
	}

	if template.Datacenter.Name == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Datacenter.Name: specifies which datacenter the instance is to be provisioned in")
	}

	if errorMessage != "" {
		err = errors.New(errorMessage)
	}

	return err
}

func (slvgs *softLayer_Virtual_Guest_Service) getRootPassword(virtualGuest datatypes.SoftLayer_Virtual_Guest) string {
	passwords := virtualGuest.OperatingSystem.Passwords

	for _, password := range passwords {
		if password.Username == ROOT_USER_NAME {
			return password.Password
		}
	}

	return ""
}

func (slvgs *softLayer_Virtual_Guest_Service) findUpgradeItemPriceForEphemeralDisk(instanceId int, ephemeralDiskSize int) (datatypes.SoftLayer_Item_Price, error) {
	if ephemeralDiskSize <= 0 {
		return datatypes.SoftLayer_Item_Price{}, errors.New(fmt.Sprintf("Ephemeral disk size can not be negative: %d", ephemeralDiskSize))
	}

	itemPrices, err := slvgs.GetUpgradeItemPrices(instanceId)
	if err != nil {
		return datatypes.SoftLayer_Item_Price{}, nil
	}

	var currentDiskCapacity int
	var currentItemPrice datatypes.SoftLayer_Item_Price

	for _, itemPrice := range itemPrices {

		flag := false
		for _, category := range itemPrice.Categories {
			if category.CategoryCode == EPHEMERAL_DISK_CATEGORY_CODE {
				flag = true
				break
			}
		}

		if flag && strings.Contains(itemPrice.Item.Description, "(LOCAL)") {

			capacity, _ := strconv.Atoi(itemPrice.Item.Capacity)

			if capacity >= ephemeralDiskSize {
				if currentItemPrice.Id == 0 || currentDiskCapacity >= capacity {
					currentItemPrice = itemPrice
					currentDiskCapacity = capacity
				}
			}
		}
	}

	if currentItemPrice.Id == 0 {
		return datatypes.SoftLayer_Item_Price{}, errors.New(fmt.Sprintf("No proper local disk for size %d", ephemeralDiskSize))
	}

	return currentItemPrice, nil
}
