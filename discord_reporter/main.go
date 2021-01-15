package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type configurationType struct {
	authenticationToken string
	guild               string
	channel             string
}

var (
	cfg configurationType
)

func main() {
	cfg.authenticationToken = ""
	cfg.guild = ""
	cfg.channel = "general"

	discord, err := discordgo.New("Bot " + cfg.authenticationToken)
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(ready)

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot is running. ctrl-c to exit.")

	// create a channel to pass information from testBackgroundProcess to a process waiting to send messages to the guild
	messages := make(chan string)
	channelID, err := getChannelID(discord)
	if err != nil {
		log.Fatal(err)
	}
	// start processes
	go programMessage(discord, channelID, messages)
	go testBackgroundProcess(messages)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()
}

// background listener simply receives string messages and sends to discord channel
func programMessage(s *discordgo.Session, channelID string, messages chan string) {
	for {
		msg := <-messages
		s.ChannelMessageSend(channelID, msg)
	}
}

// example background process
// simulate looped code that would want to communicate to a discord server
func testBackgroundProcess(messages chan<- string) {
	for {
		time.Sleep(time.Second * 8)
		messages <- "message from separate process"

	}
}

// might not be necessary
// for some reason s.State.Guilds[0].Name is unpopulated for me
// this is true even after discordgo.Ready event
func getChannelID(s *discordgo.Session) (string, error) {
	for _, guild := range s.State.Guilds {
		guild, err := s.Guild(guild.ID)
		if err != nil {
			return "", err
		}
		if guild.Name == cfg.guild {
			channels, err := s.GuildChannels(guild.ID)
			if err != nil {
				return "", err
			}
			for _, channel := range channels {
				if channel.Name == cfg.channel {
					return channel.ID, nil
				}
			}
		}
	}
	return "", errors.New("could not obtain channelID, ensure guild & channel names are correct")
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot is Ready")
}
