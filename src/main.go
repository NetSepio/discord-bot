package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	BotPrefix string
	config    *configstruct
)

type configstruct struct {
	Token     string `json: "Token"`
	BotPrefix string `json: "BotPrefix"`
}

var BotId string
var goBot *discordgo.Session

//Check for links in channel


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