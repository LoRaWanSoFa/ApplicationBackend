package components

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type DatabaseData struct {
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
}

var once sync.Once
var settings DatabaseData

func GetConfiguration() DatabaseData {
	once.Do(func() {
		// START: yaml config block
		goPath := os.Getenv("GOPATH")
		yamlFile, err := ioutil.ReadFile(filepath.Join(goPath, "/src/github.com/LoRaWanSoFa/LoRaWanSoFa/config.yaml"))
		if err != nil {
			panic(err)
		}
		//log.Printf("%+v", string(yamlFile))
		settings = DatabaseData{}
		err = yaml.Unmarshal(yamlFile, &settings)
		if err != nil {
			panic(err)
		}
		//log.Printf("%+v", dbData)
		// END: yaml config block
	})
	return settings
}
