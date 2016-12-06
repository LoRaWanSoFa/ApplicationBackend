package components

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	orgFile := "/src/github.com/LoRaWanSoFa/LoRaWanSoFa/config.yaml"
	newFile := "/src/github.com/LoRaWanSoFa/LoRaWanSoFa/config_TEST_FILE.yaml"
	os.Rename(filepath.Join(goPath, orgFile), filepath.Join(goPath, newFile))
	settings := GetConfiguration()
	if settings.Db.Port != 0 {
		t.Error("File was found?!")
	}
	settings = ReloadConfig()
	if settings.Db.Port != 0 {
		t.Error("File was found?!")
	}
	os.Rename(filepath.Join(goPath, newFile), filepath.Join(goPath, orgFile))
	settings = ReloadConfig()
	if settings.Db.Port == 0 {
		t.Error("Port: can not be 0")
	}

}
