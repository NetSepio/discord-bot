package main
import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

var (
	Token string
	BotPrefix string
	config *configstruct
)

type configstruct struct{
	Token string `json: "Token"`
	BotPrefix string `json: "BotPrefix"`
}

func ReadConfig() error{
	file,err:=ioutil.ReadFile("./config.json")
	if err!=nil{
		fmt.Println(err);
		return err;
	}
	fmt.Println(string(file))
	err= json.Unmarshal(file,&config)
	if err!=nil{
		fmt.Println(err);
		return err;
	}
	Token=config.Token
	BotPrefix=config.BotPrefix
	return nil;
}

var BotId string;
var goBot *discordgo.Session

func Start(){
	goBot,err:=discordgo.New("Bot "+ config.Token)
	if err!=nil{
		fmt.Println(err);
		return 
	}
	u,err:=goBot.User("@me")
	if err!=nil{
		fmt.Println(err);
		return 
	}
	BotId=u.ID;
	goBot.AddHandler(messageHandler)
	err=goBot.Open()
	if err!=nil{
		fmt.Println(err);
		return 
	}
	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID==BotId{
		return
	}
	if m.Content==BotPrefix+"test"{
		_,_=s.ChannelMessageSend(m.ChannelID,"test")
	}
}

func main(){
	err:=ReadConfig();
	if err!=nil{
		fmt.Println(err);
		return 
	}
	Start()
}