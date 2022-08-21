package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}