package main

import (
	"bufio"
	"log"
	"os"

	telegram "github.com/callmehorhe/dzyrduat/telegram"
	"github.com/callmehorhe/dzyrduat/translater"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	yryssagDict []string
	ironDict    []string
)

//preparing
func main() {
	yryssagDict = openDict("files/yryssag.txt")
	ironDict = openDict("files/iron.txt")
	if err := godotenv.Load(); err != nil {
		log.Panic("Env vars load error")
		return
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Panic("Bot start error")
		return
	}
	translaterService := translater.NewDicts(yryssagDict, ironDict)
	telegramService := telegram.NewBot(bot, translaterService)
	telegramService.Start()
}

func openDict(name string) []string {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		log.Print("ERR")
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var dict []string
	for sc.Scan() {
		dict = append(dict, sc.Text())
	}
	return dict
}
