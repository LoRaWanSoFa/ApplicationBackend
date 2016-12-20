package components

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

// Configuration is the Struct that will contain the values that are read from
// config.yaml file in the root path of the github repository.
type Configuration struct {
	Mqtt struct {
		AppEUI   string `yaml:"AppEUI"`
		Password string `yaml:"Password"`
		Address  string `yaml:"Address"`
	}
	Db struct {
		User            string `yaml:"User"`
		Password        string `yaml:"Password"`
		Name            string `yaml:"Name"`
		Network         string `yaml:"Network"`
		Port            int    `yaml:"Port"`
		NumberOfWorkers int    `yaml:"NumberOfWorkers"`
	}
	Rest struct {
		IP     string `yaml:"Ip"`
		APIKey string `yaml:"ApiKey"`
	}
}

var once sync.Once
var settings Configuration

// GetConfiguration reads the configuration from the standard path if ReloadConfig
// has not been run beforehand. If ReloadConfig has been run beforehand it will
// not set it to the standard path.
func GetConfiguration() Configuration {
	once.Do(func() {
		settings = ReloadConfig("/src/github.com/LoRaWanSoFa/ApplicationBackend/config.yaml")
	})
	return settings
}

// ReloadConfig reloads the config with a custom path. Calling GetConfiguration
// afterwards will allow the user to get the configuration file from the newly
// set path.
func ReloadConfig(path string) Configuration {
	go once.Do(func() {})
	// START: yaml config block
	goPath := os.Getenv("GOPATH")
	yamlFile, err := ioutil.ReadFile(filepath.Join(goPath, path))
	if err != nil {
		return Configuration{}
	}
	settings = Configuration{}
	err = yaml.Unmarshal(yamlFile, &settings)
	if err != nil {
		fmt.Println(err)
	}
	// END: yaml config block
	return settings
}
