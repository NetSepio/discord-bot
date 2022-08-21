package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)


func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err)
		return
	}
	BotId = u.ID
	goBot.AddHandler(Validator)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bot is live")
}