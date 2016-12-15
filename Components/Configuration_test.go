package components

import "testing"

func TestGetConfiguration(t *testing.T) {
	orgFile := "/src/github.com/LoRaWanSoFa/LoRaWanSoFa/config.yaml"
	settings := GetConfiguration()
	if settings.Db.Port == 0 {
		t.Errorf("File was not found")
	}
	settings = ReloadConfig("nonexisting/path_/??!")
	if settings.Db.Port != 0 {
		t.Errorf("File was suprisingly found: %+v", settings)
	}
	settings = ReloadConfig(orgFile)
	if settings.Db.Port == 0 {
		t.Error("Port: can not be 0")
	}

}
