package main

import (
	"fmt"
	"github.com/myuteBranch/scrubBot/utils"
	"github.com/sirupsen/logrus"
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

var log = logrus.New()

func init() {
	log.SetLevel(logrus.InfoLevel)
}

func main() {

	Token := os.Getenv("bot_token")
	log.Trace("Bot Token = ", Token)
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Info("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Debug("Message Recieved : ", m.Content, " from : ", m.Author.Username)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.Contains(m.Content, "!") {
		log.Debug("Message Recieved : does not have an !")
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "!dota_matches_all" {
		outString := utils.ChunkString(utils.GetFormatedMatches(utils.GetDotaMatches()), 1800)
		for _, chunk := range outString {
			sendMessage(s, m.ChannelID, fmt.Sprintf("```%s```", chunk))
		}
	}

	if m.Content == "!dota_matches" {
		outString := utils.ChunkString(utils.GetFormatedMatches(utils.GetDotaMatches()), 1800)
		sendMessage(s, m.ChannelID, fmt.Sprintf("```%s \n \t More ...```", outString[0]))
	}

	if strings.HasPrefix(m.Content, "!linkMe") {
		ch, _ := s.UserChannelCreate(m.Author.ID)
		sendMessage(s, ch.ID, utils.GetLinkForMessage(m.Content))
	}
}

func sendMessage(s *discordgo.Session, channelID string, writeString string) {
	_, err := s.ChannelMessageSend(channelID, writeString)
	if err != nil {
		log.Error("error sending match data", err)
	}
}
