package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	ConfigVersion       int                     `json:"configversion"`
	Port                int                     `json:"port"`
	Webdev              bool                    `json:"webdev"`
	EnablePrometheus    bool                    `json:"enableprometheus"`
	MapBlockAccessorCfg *MapBlockAccessorConfig `json:"mapblockaccessor"`
}

type MapBlockAccessorConfig struct {
	Expiretime string `json:"expiretime"`
	Purgetime  string `json:"purgetime"`
	MaxItems   int    `json:"maxitems"`
}

var lock sync.Mutex

const ConfigFile = "mtworldview.json"

func (cfg *Config) Save() error {
	return WriteConfig(ConfigFile, cfg)
}

func WriteConfig(filename string, cfg *Config) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	str, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return err
	}

	f.Write(str)

	return nil
}

func ParseConfig(filename string) (*Config, error) {

	mapblockaccessor := MapBlockAccessorConfig{
		Expiretime: "15s",
		Purgetime:  "30s",
		MaxItems:   500,
	}

	cfg := Config{
		ConfigVersion:       1,
		Port:                8080,
		Webdev:              false,
		EnablePrometheus:    true,
		MapBlockAccessorCfg: &mapblockaccessor,
	}

	info, err := os.Stat(filename)
	if info != nil && err == nil {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
