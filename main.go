package main

import (
	"log"
	"strings"

	"git.wetofu.top/tonychee7000/tgBotPill/bot"
	"git.wetofu.top/tonychee7000/tgBotPill/config"
	"git.wetofu.top/tonychee7000/tgBotPill/consts"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("Hello!")
	bot := bot.NewBot(config.DefaultConfig)
	err := bot.Connect()
	if err != nil {
		log.Fatalln("FATAL:", err)
	}

	log.Printf("Bot authoirzed by name: %s(%d)", bot.Name, bot.ID)

	bot.Run(proceed)

}

func proceed(message *bot.Message) string {
	log.Printf("From=%s Text=\"%s\"", message.From.UserName, message.Text)
	args := strings.Split(message.Text, " ")
	cmd := strings.Split(args[0], "@")[0]
	switch cmd {
	case "/null":
		return consts.ReplyLue
	case "/help":
		return "help" // TODO: help txt
	default:
		return ""
	}
}
