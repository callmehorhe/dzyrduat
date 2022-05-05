package telegram

import (
	"log"

	translater "github.com/callmehorhe/dzyrduat/translater"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	language = make(map[int64]string)
)

type Bot struct {
	bot   *tgbotapi.BotAPI
	dicts *translater.Dicts
}

func NewBot(bot *tgbotapi.BotAPI, dicts *translater.Dicts) *Bot {
	return &Bot{
		bot:   bot,
		dicts: dicts,
	}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates := b.updatesInit()
	b.handleUpdates(updates)
}

func (b *Bot) updatesInit() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			b.handleNewMessages(update.Message)
		}
	}
}

func (b *Bot) handleNewMessages(message *tgbotapi.Message) {
	if message.IsCommand() {
		b.handleCommands(message)
	} else {
		b.handleMessage(message)
	}
}

func (b *Bot) handleCommands(message *tgbotapi.Message) {
	command := message.Command()
	switch command {
	case "start":
		language[message.Chat.ID] = "iron"
	case "iron":
		language[message.Chat.ID] = "iron"
		b.sendMessgae(message.Chat.ID, "Перевод с осетинского на русский!")
	case "yryssag":
		language[message.Chat.ID] = "yryssag"
		b.sendMessgae(message.Chat.ID, "Перевод с русского на осетинский!")
	case "about":
	default:
		break
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	translate := b.dicts.Translate(message.Text, language[message.Chat.ID])
	b.sendMessgae(message.Chat.ID, translate)
}

func (b *Bot) sendMessgae(chatID int64, text string) {
	b.bot.Send(tgbotapi.NewMessage(chatID, text))
}
