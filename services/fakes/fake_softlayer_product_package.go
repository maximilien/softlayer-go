package test_helpers

import (
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
	"strings"
)

type FakeProductPackageService struct{}

func (fps *FakeProductPackageService) GetName() string {
	return "Mock_Product_Package_Service"
}

func (fps *FakeProductPackageService) GetItemsByType(packageType string) ([]datatypes.SoftLayer_Product_Item, error) {
	response, _ := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Product_Package_getItemsByType_virtual_server.json")

	productItems := []datatypes.SoftLayer_Product_Item{}
	json.Unmarshal(response, &productItems)

	return productItems, nil
}

func (fps *FakeProductPackageService) GetItemPrices(packageId int, filters string) ([]datatypes.SoftLayer_Product_Item_Price, error) {
	var textFixtrue = ""

	switch {
	case strings.Contains(filters, "keyName"):
		textFixtrue = "SoftLayer_Product_Package_getItemPrices.json"
	case strings.Contains(filters, "100") || strings.Contains(filters, "1000") || strings.Contains(filters, "2000"):
		textFixtrue = "SoftLayer_Product_Package_getItemPricesBySizeAndIops.json"
	}

	if textFixtrue == "" {
		return []datatypes.SoftLayer_Product_Item_Price{}, errors.New(fmt.Sprintf("No matched IOPS found for filters %s", filters))
	}

	response, _ := testhelpers.ReadJsonTestFixtures("services", textFixtrue)
	itemPrices := []datatypes.SoftLayer_Product_Item_Price{}
	json.Unmarshal(response, &itemPrices)

	return itemPrices, nil
}

func (fps *FakeProductPackageService) GetItems(packageId int, filters string) ([]datatypes.SoftLayer_Product_Item, error) {
	return []datatypes.SoftLayer_Product_Item{}, errors.New("Not supported")
}

func (fps *FakeProductPackageService) GetPackagesByType(packageType string) ([]datatypes.Softlayer_Product_Package, error) {
	return []datatypes.Softlayer_Product_Package{}, errors.New("Not supported")
}

func (fps *FakeProductPackageService) GetOnePackageByType(packageType string) (datatypes.Softlayer_Product_Package, error) {
	return datatypes.Softlayer_Product_Package{}, errors.New("Not supported")
}
