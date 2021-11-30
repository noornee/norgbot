package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configStruct struct {
	Token string `json:"Token"`
}

var (
	Token  string
	config *configStruct
)

func ReadConfig() error {

	fmt.Println("Reading from config")

	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	Token = config.Token

	return nil

}
