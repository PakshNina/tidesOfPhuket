package main

import (
	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tidesOfPhuket/internal/tidesofphuket/client"
	"tidesOfPhuket/internal/tidesofphuket/service"

	"tidesOfPhuket/internal/tidesofphuket/config"
	//"tidesOfPhuket/internal/tidesofphuket/service"
	"tidesOfPhuket/internal/redis"
)

func main() {
	cfg, err := config.NewTidesOfPhuketConfig()
	if err != nil {
		log.Fatal("Error")
	}
	r := redis.NewClient(cfg)
	conn := client.Connection{
		Url:    cfg.WorldTideUrl,
		ApiKey: cfg.WorldTideApiKey,
	}
	c := client.NewWorldTidesClient(conn)
	bot, err := telebot.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Error with new telegram bot: %v", err)
	}
	serviceError := service.RunService(bot, r, c)
	if serviceError != nil {
		log.Fatal("Error")
	}
}
