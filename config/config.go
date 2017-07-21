package config

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

type config struct {
	TelegramKey string `json:"telegram_key"`
	BysykkelKey string `json:"bysykkel_key"`
}

// FromYAML reads from a YAML file
func FromYAML(filepath string) config {
	byt, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(errors.Wrap(err, "Could not load YAML file: "+filepath))
	}
	var c config
	if err := yaml.Unmarshal(byt, &c); err != nil {
		panic(errors.Wrap(err, "Could not parse YAML file: "+filepath))
	}
	return c
}
