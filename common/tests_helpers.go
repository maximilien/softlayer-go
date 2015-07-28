package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadJsonTestFixtures(packageName, fileName string) ([]byte, error) {
	wd, _ := os.Getwd()
	return ioutil.ReadFile(filepath.Join(wd, "..", "test_fixtures", packageName, fileName))
}
