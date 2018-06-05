package bot

import (
	"log"

	"git.wetofu.top/tonychee7000/tgBotPill/config"
	tgApi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Message for orig tgAPI
type Message tgApi.Message

// ProceedFunc recv and reply
type ProceedFunc func(*Message) string

// Bot a bot
type Bot struct {
	bot     *tgApi.BotAPI
	conf    config.Config
	uc      tgApi.UpdateConfig
	updates tgApi.UpdatesChannel
	Name    string
	ID      int
}

// NewBot for create
func NewBot(conf config.Config) *Bot {
	var b = new(Bot)
	b.conf = conf
	b.updates = make(tgApi.UpdatesChannel, 1024)
	return b
}

// Connect to bot
func (b *Bot) Connect() error {
	var err error
	b.bot, err = tgApi.NewBotAPI(b.conf.APIToken)
	if err != nil {
		return err
	}

	b.bot.Debug = b.conf.Debug
	b.Name = b.bot.Self.UserName
	b.ID = b.bot.Self.ID

	b.uc = tgApi.NewUpdate(0)
	b.uc.Timeout = b.conf.UpdateTimeout

	b.updates, err = b.bot.GetUpdatesChan(b.uc)
	if err != nil {
		return err
	}
	return nil
}

// Run to proceed
func (b *Bot) Run(f ProceedFunc) {
	for {
		select {
		case update := <-b.updates:
			var m Message
			if update.Message == nil {
				continue
			} else {
				m = Message(*update.Message)
			}

			reply := f(&m)
			if reply == "" {
				continue
			}
			msg := tgApi.NewMessage(m.Chat.ID, reply)
			msg.ReplyToMessageID = m.MessageID
			_, err := b.bot.Send(msg)
			if err != nil {
				log.Println("ERROR:", err)
			}
		}
	}
}
