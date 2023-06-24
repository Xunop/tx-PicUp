package main

import (
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Url       string
	SecretID  string
	SecretKey string
}

var (
	cfg  *tomlConfig
	once sync.Once
)

func Config() *tomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("./config.toml")
		if err != nil {
			panic(err)
		}
		println("config file path: ", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
