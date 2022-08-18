package main
import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"regexp"
	"github.com/hasura/go-graphql-client"
	"context"
	"strings"

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
	goBot.AddHandler(validator)
	err=goBot.Open()
	if err!=nil{
		fmt.Println(err);
		return 
	}
	fmt.Println("Bot is running")
}

//Check for links in channel
func validator(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID==BotId{
		return
	}
	//Regex check for a link
	regexCheck:=`^((https?|ftp|smtp):\/\/)?(www.)?[a-z0-9]+\.[a-z]+(\/[a-zA-Z0-9#]+\/?)*$`
	match, _ := regexp.MatchString(regexCheck, m.Content)
    fmt.Println(match)

	if match==true{
		_,_=s.ChannelMessageSend(m.ChannelID,"This is a link")
		client := graphql.NewClient("https://query.graph.lazarus.network/subgraphs/name/NetSepio",nil)
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
		//fmt.Println(string(e))
		if err != nil {
			fmt.Println(err)
		}	
		var substr = "github.com"
		i := strings.Index(string(e), substr)
		fmt.Println(i)
		fmt.Println(string(e)[i:i+40])
		//get index of } and index of : and substring and display
	}
}

func main(){
	err:=ReadConfig();
	if err!=nil{
		fmt.Println(err);
		return 
	}
	Start()
	<-make(chan struct{})
	return
}