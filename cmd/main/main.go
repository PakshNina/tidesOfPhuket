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
	cfg, errCfg := config.NewTidesOfPhuketConfig()
	if errCfg != nil {
		log.Fatalf("Error with config: %v", errCfg)
	}
	redisClient := redis.NewClient(cfg)
	conn := client.Connection{
		Url:    cfg.WorldTideUrl,
		ApiKey: cfg.WorldTideApiKey,
	}
	tidesClient := client.NewWorldTidesClient(conn)
	bot, errBot := telebot.NewBotAPI(cfg.TelegramToken)
	if errBot != nil {
		log.Fatalf("Error with getting telegram bot: %v", errBot)
	}
	errService := service.RunService(bot, redisClient, tidesClient)
	if errService != nil {
		log.Fatalf("Error with service: %v", errService)
	}
}
