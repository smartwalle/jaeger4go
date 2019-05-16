package jaeger4go

import (
	"errors"
	"github.com/uber/jaeger-client-go/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ErrNotFoundConfigFile = errors.New("jaeger4go: not found config file")
)

type Config struct {
	config.Configuration
}

func Load(path string) (cfg *Config, err error) {
	if len(path) == 0 {
		return nil, ErrNotFoundConfigFile
	}

	cfgData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var jCfg config.Configuration
	if err = yaml.Unmarshal(cfgData, &jCfg); err != nil {
		return nil, err
	}

	cfg = &Config{}
	cfg.Configuration = jCfg

	return cfg, nil
}
