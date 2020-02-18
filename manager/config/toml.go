package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// manager config = mc
// yes u r mc boy!

type Mc struct {
	Global global // global table options
}

type global struct {
	Listen   string // ip address or ethernet interface which service listening
	Port     int    // port
	LogLevel string // log level
}

var mConfig *Mc

// refresh config when loading
func init() {
	RefreshConfig()
}

// refresh config from toml file
func RefreshConfig() {
	mConfig = unmarshal()
}

// return manager config
func GetConf() Mc {
	return *mConfig
}

// unmarshal config from manager.toml to Mc struct
func unmarshal() *Mc {
	conf := new(Mc)
	// set default values
	conf.Global.Listen = "0.0.0.0"
	conf.Global.Port = 8292
	conf.Global.LogLevel = "info"

	if _, err := toml.DecodeFile("/etc/kiwi/kiwi_manager.toml", conf); err != nil {
		log.Fatalf("unmarshal kiwi manager config failed, %s", err)
		return nil
	}
	return conf
}
