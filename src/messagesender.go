package main

import (
	"strings"
	"mvdan.cc/xurls/v2"
	"github.com/bwmarrin/discordgo"
)

func DeciderType(dataString string,s *discordgo.Session, m *discordgo.MessageCreate){
	rxRelaxed := xurls.Relaxed()
	
		safe:=strings.Count(dataString, "Safe")
		phishing:=strings.Count(dataString, "Phishing")
		vsafe:=strings.Count(dataString, "Very safe")
		spyware:=strings.Count(dataString, "Spyware")
		genuine:=strings.Count(dataString, "genuine")
		malware:=strings.Count(dataString, "Malware")
		vsafeandsmooth:=strings.Count(dataString, "Very safe and smooth")
		if(safe>phishing && safe>vsafe && safe>spyware && safe>genuine && safe>malware && safe>vsafeandsmooth ){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ rxRelaxed.FindString(m.Content)+" is classified as Safe`")
		}else if(phishing>safe && phishing>vsafe && phishing>spyware && phishing>genuine && phishing>malware && phishing>vsafeandsmooth){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ rxRelaxed.FindString(m.Content)+" is classified as Phishing`")
		}else if(safe>phishing && safe>vsafe && safe>spyware && safe>genuine && safe>malware && safe>vsafeandsmooth){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ rxRelaxed.FindString(m.Content)+" is classified as Safe`")
		}else if(genuine>phishing && genuine>vsafe && genuine>spyware && genuine>safe && genuine>malware && genuine>vsafeandsmooth){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ rxRelaxed.FindString(m.Content)+" is classified as Genuine`")
		}else if(vsafeandsmooth>phishing && vsafeandsmooth>vsafe && vsafeandsmooth>spyware && vsafeandsmooth>safe && vsafeandsmooth>malware && vsafeandsmooth>genuine){
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ rxRelaxed.FindString(m.Content)+" is classified as vsafeandsmooth`")
		}else{
			_, _ = s.ChannelMessageSend(m.ChannelID, "`The link "+ rxRelaxed.FindString(m.Content)+" is not tested`")
		}
		return 
}