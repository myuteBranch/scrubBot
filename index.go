package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
// var (
// 	Token string
// )

// func init() {

// 	flag.StringVar(&Token, "t", "", "Bot Token")
// 	flag.Parse()
// }

func main() {
	Token := os.Getenv("bot_token")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "!dota_matches_all" {
		outString := chunkString(getFormatedMatches(getDotaMatches()), 1800)
		for _, chunk := range outString {
			sendMessage(s, m.ChannelID, fmt.Sprintf("```%s```", chunk))
		}
	}

	if m.Content == "!dota_matches" {
		outString := chunkString(getFormatedMatches(getDotaMatches()), 1800)
		sendMessage(s, m.ChannelID, fmt.Sprintf("```%s \n \t More ...```", outString[0]))
	}

	if m.Content == "!sendMe" {
		ch, _ := s.UserChannelCreate(m.Author.ID)
		sendMessage(s, ch.ID, "Help!")
	}

	if strings.HasPrefix(m.Content, "!linkMe") {
		ch, _ := s.UserChannelCreate(m.Author.ID)
		if len(strings.Split(m.Content, " ")) > 1 && strings.Split(m.Content, " ")[1] != "" {
			if strings.Split(m.Content, " ")[1] == "help" {
				keys := make([]string, 0, len(links))
				for key := range links {
					keys = append(keys, key)
				}
				sendMessage(s, ch.ID, fmt.Sprintf("```Available links are: \n!linkMe %s ```", strings.Join(keys, "\n!linkMe ")))
				return
			}
			if links[strings.Split(m.Content, " ")[1]] != "" {
				sendMessage(s, ch.ID, fmt.Sprintf("%s :  %s ", strings.Split(m.Content, " ")[1], links[strings.Split(m.Content, " ")[1]]))
				return
			}
		}
		sendMessage(s, ch.ID, "```Invalid argument for command !linkMe for valid options try \n try !linkMe help  ```")
	}
}

func sendMessage(s *discordgo.Session, channelID string, writeString string) {
	_, err := s.ChannelMessageSend(channelID, writeString)
	if err != nil {
		fmt.Println("error sending match data", err)
	}
}
