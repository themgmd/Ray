package config

import (
	"errors"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	Package string `json:"package"`
}

var config = &Config{}

func Get() *Config {
	return config
}

func sync(file *os.File) error {
	fst, err := file.Stat()
	if err != nil {
		return err
	}

	data := make([]byte, fst.Size())
	_, err = file.Read(data)
	if err != nil {
		return err
	}

	return toml.Unmarshal(data, config)
}

func update() error {
	tomlConf, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile("config.toml", tomlConf, os.ModeAppend)
}

func Init() error {
	file, err := os.Open("config.toml")
	if err != nil {
		return err
	}

	defer func() {
		err = errors.Join(err, file.Close())
	}()

	return sync(file)
}

func SetPackage(pkg string) error {
	config.Package = pkg
	return update()
}
