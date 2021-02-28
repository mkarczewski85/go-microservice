package config

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		DatabaseName   string `yaml:"name", envconfig:"DB_NAME"`
		CollectionName string `yaml:"collection", envconfig:"DB_COLLECTION"`
		ConnectionUri  string `yaml:"uri", envconfig:"DB_URI"`
	} `yaml:"database"`
}
