package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// Config contains the API keys
type Config struct {
	TelegramKey string `json:"telegram_key"`
	BysykkelKey string `json:"bysykkel_key"`
}

// FromYAML reads from a YAML file
func FromYAML(filepath string) Config {
	byt, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(errors.Wrap(err, "Could not load YAML file: "+filepath))
	}
	var c Config
	if err := yaml.Unmarshal(byt, &c); err != nil {
		panic(errors.Wrap(err, "Could not parse YAML file: "+filepath))
	}
	return c
}

// GetKeys returns keys, according to deploy kind
func GetKeys() (string, string) {
	switch os.Getenv("DEPLOY_KIND") {
	case "local":
		config := FromYAML("config/config.yaml")
		return config.TelegramKey, config.BysykkelKey
	case "cloud":
		return os.Getenv("TELEGRAM_KEY"), os.Getenv("BYSYKKEL_KEY")
	default:
		log.Println("DEPLOY_KIND is not set")
		panic(errors.New("DEPLOY_KIND is not set"))
	}
}
