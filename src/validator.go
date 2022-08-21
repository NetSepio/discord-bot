package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
	"github.com/bwmarrin/discordgo"
	"mvdan.cc/xurls/v2"
)

func Validator(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	//Regex check for a link
	rxRelaxed := xurls.Relaxed()
	fmt.Print(rxRelaxed.FindString("Do gophers live in golan?") )
	if(rxRelaxed.FindString(m.Content)!=""){
		fmt.Println(rxRelaxed.FindString(m.Content) )
	}
	regexCheck := `^((https?|ftp|smtp):\/\/)?(www.)?[a-z0-9]+\.[a-z]+(\/[a-zA-Z0-9#]+\/?)*$`
	match, _ := regexp.MatchString(regexCheck, m.Content)
	if match == true {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hang on! NetSepio is verifying the link")
		urlCheck:=m.Content
		 queryy:=`
		{
			reviews(where: {domainAddress_contains:"`+urlCheck+`"}) {
			  siteURL
			  siteSafety
			}
		  }
		`
		jsonData := map[string]string{
			"query": queryy,
		}
		jsonValue, _ := json.Marshal(jsonData)
		request, err := http.NewRequest("POST", "https://query.graph.lazarus.network/subgraphs/name/NetSepio", bytes.NewBuffer(jsonValue))
		client := &http.Client{Timeout: time.Second * 100}
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		dataString:=string(data)
		DeciderType(dataString,s,m)
	
	}
}