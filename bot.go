package main

import (
  "fmt"
  "os"

  "github.com/bwmarrin/discordgo"
)

const (
  token string = os.Getenv("DISCORD_BOT_TOKEN")
)

// var (
//   token string
// )

func init() {
  // token = os.Getenv("DISCORD_BOT_TOKEN")
}

func main() {

  init()

  discord, err := discordgo.New("Bot " + token)
  if err != nil {
    fmt.Println("Error creating a Discord session,", err.Error())
    return
  }

  // add handlers here
  // discord.AddHandler(onCommand)

  err = discord.Open()
  if err != nil {
    fmt.Println("Error opening connection,", err.Error())
  }

  fmt.Println("Bot is now running.")

}
