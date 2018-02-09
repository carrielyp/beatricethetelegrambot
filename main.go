package main

import (
	"fmt"
	"log"
	"strings"
	"telegram_bot/lunchvenue"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("492536683:AAGBJjFK3ZxbHtuX_iOMwq00ze0Fq8jBb2M")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		messages := strings.Fields(update.Message.Text)
		fmt.Println("---------------------------------------")
		fmt.Println(messages)
		var reply string

		switch messages[0] {
		case "/add_lunch_venue":
			fmt.Println("---------------------------------------")
			fmt.Println("entered /add_lunch_venue")
			lunchvenue.AddLunchVenue(messages[1])
			reply = messages[1] + " added as a new venue!"
		case "/list_lunch_venues":
			fmt.Println("---------------------------------------")
			fmt.Println("entered /list_lunch_venues")
			reply = strings.Join(lunchvenue.ListLunchVenues(), "\n")
		}

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		bot.Send(msg)
	}
}
