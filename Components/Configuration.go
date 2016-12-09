package components

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

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
		Ip     string `yaml:"Ip"`
		ApiKey string `yaml:"ApiKey"`
	}
}

var once sync.Once
var settings Configuration

func GetConfiguration() Configuration {
	once.Do(func() {
		settings = ReloadConfig()
	})
	return settings
}

func ReloadConfig() Configuration {
	// START: yaml config block
	goPath := os.Getenv("GOPATH")
	yamlFile, err := ioutil.ReadFile(filepath.Join(goPath, "/src/github.com/LoRaWanSoFa/LoRaWanSoFa/config.yaml"))
	if err != nil {
		return settings
	}
	settings = Configuration{}
	err = yaml.Unmarshal(yamlFile, &settings)
	if err != nil {
		fmt.Println(err)
	}
	// END: yaml config block
	return settings
}
