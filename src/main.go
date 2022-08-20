package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

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

func DeciderType(dataString string,s *discordgo.Session, m *discordgo.MessageCreate){
	safe:=strings.Count(dataString, "Safe")
		phishing:=strings.Count(dataString, "Phishing")
		vsafe:=strings.Count(dataString, "Very safe")
		spyware:=strings.Count(dataString, "Spyware")
		genuine:=strings.Count(dataString, "genuine")
		malware:=strings.Count(dataString, "Malware")
		vsafeandsmooth:=strings.Count(dataString, "Very safe and smooth")
		if(safe>phishing && safe>vsafe && safe>spyware && safe>genuine && safe>malware && safe>vsafeandsmooth ){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ m.Content+" is classified as Safe`")
		}else if(phishing>safe && phishing>vsafe && phishing>spyware && phishing>genuine && phishing>malware && phishing>vsafeandsmooth){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ m.Content+" is classified as Phishing`")
		}else if(safe>phishing && safe>vsafe && safe>spyware && safe>genuine && safe>malware && safe>vsafeandsmooth){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ m.Content+" is classified as Safe`")
		}else if(genuine>phishing && genuine>vsafe && genuine>spyware && genuine>safe && genuine>malware && genuine>vsafeandsmooth){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ m.Content+" is classified as Genuine`")
		}else if(vsafeandsmooth>phishing && vsafeandsmooth>vsafe && vsafeandsmooth>spyware && vsafeandsmooth>safe && vsafeandsmooth>malware && vsafeandsmooth>genuine){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ m.Content+" is classified as vsafeandsmooth`")
		}else{
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ m.Content+" is not tested`")
		}
		return 
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
		client := &http.Client{Timeout: time.Second * 10}
		response, err := client.Do(request)
		//defer response.Body.Close()
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		data, _ := ioutil.ReadAll(response.Body)
		dataString:=string(data)
		DeciderType(dataString,s,m)
		
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