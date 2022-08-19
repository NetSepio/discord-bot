package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/hasura/go-graphql-client"
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

var BotId string
var goBot *discordgo.Session

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
	goBot.AddHandler(validator)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bot is live")
}

//Check for links in channel
func validator(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	//Regex check for a link
	regexCheck := `^((https?|ftp|smtp):\/\/)?(www.)?[a-z0-9]+\.[a-z]+(\/[a-zA-Z0-9#]+\/?)*$`
	match, _ := regexp.MatchString(regexCheck, m.Content)
	if match == true {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hang on! NetSepio is verifying the link")
		client := graphql.NewClient("https://query.graph.lazarus.network/subgraphs/name/NetSepio", nil)
		var q struct {
			Reviews []struct {
				DomainAddress string `json:"domainAddress"`
				SiteSafety    string `json:"siteSafety"`
			} `json:"reviews"`
		}
		err := client.Query(context.Background(), &q, nil)
		if err != nil {
			fmt.Println(err)
		}
		e, err := json.Marshal(q)
		if err != nil {
			fmt.Println(err)
		}
		var substr = m.Content
		i := strings.Index(string(e), substr)
		if i != -1 {
			b := strings.Index(string(e)[i:i+80], ":")
			c := strings.Index(string(e)[i:i+80], "}")
			initPrint := i + b
			initPrint2 := i + c
			var sendMessage = string(e)[initPrint+2 : initPrint2-1]
			_, _ = s.ChannelMessageSend(m.ChannelID, "`"+m.Content+" is classified as: "+sendMessage+"`")

		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "`"+m.Content+" is not in our database`")
		}
	}
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
