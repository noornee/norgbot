package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

func ReadConfig() error {

	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}
