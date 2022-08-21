package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)
var BotId string
var goBot *discordgo.Session

var (
	Token     string
	BotPrefix string
	config    *configstruct
)

type configstruct struct {
	Token     string `json: "Token"`
	BotPrefix string `json: "BotPrefix"`
}



func main() {
	err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	Start()
	<-make(chan struct{})
	return
}