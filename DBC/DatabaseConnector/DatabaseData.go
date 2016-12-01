package DatabaseConnector

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
