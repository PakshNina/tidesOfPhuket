package service

import (
	"fmt"
	"log"
	"time"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tidesOfPhuket/internal/redis"
	c "tidesOfPhuket/internal/tidesofphuket/client"
	"tidesOfPhuket/internal/tidesofphuket/coordinates"
)

const (
	PatongCommand = "/patong"
	MaiKao = "/maikao"
	Aonang = "/aonang"
	High = "High"
	Low = "Low"
)

func RunService(bot *telebot.BotAPI, redisRepo redis.Redis, client c.WorldTidesClient) error {
	u := telebot.NewUpdate(0)
	u.Timeout = 10
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			command := update.Message.Text
			log.Printf("Received message: %s", command)
			coords := getCoordsByCommand(command)
			var replyMessage string
			if coords != nil {
				result := make(chan string)
				go getTidesInfo(command, coords, result, client, redisRepo)
				replyMessage = <- result
			} else {
				replyMessage = "Hello! Please choose command to see tides times. Available commands: /patong, /maikao, /aonang"
			}
			msg := telebot.NewMessage(update.Message.Chat.ID, replyMessage)
			msg.ReplyToMessageID = update.Message.MessageID
			_, errBot := bot.Send(msg)
			if errBot != nil {
				return errBot
			}
		}
	}
	return nil
}

func getTidesInfo(beach string, coords *coordinates.Coordinates, result chan string, c c.WorldTidesClient, redisRepo redis.Redis) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)
	today := now.Format("2006-01-02")
	arrToday, errTides := redisRepo.GetExtremesForToday(beach)
	if errTides != nil {
		if errTides == redis.ErrKeyDoesNotExistsInRedis {
			log.Printf("Key was not found in redis")
			arrToday, errTides = c.GetExtremes(today, coords.Lat, coords.Lon)
			if errTides != nil {
				log.Printf("Err with getting info: %v", errTides)
				result <- "Service currently is not available. Please try later"
				return
			}
			errRedis := redisRepo.SaveTideInfo(beach, arrToday)
			if errRedis != nil {
				log.Printf("Err with saving to redis: %v", errRedis)
				result <- fmt.Sprintf("Service currently is not available. Please try later: %v", errRedis)
				return
			}
		} else {
			log.Printf("Err with getting info from redis: %v", errTides)
			result <- "Service currently is not available. Please try later"
			return
		}
	}
	highTides := ""
	lowTides := ""
	count := 0
	for _, v := range arrToday.Extremes {
		if count < 6 {
			dt := time.Unix(v.Date, 0).UTC().In(loc).Format("02.01 15:04")
			switch v.TideType {
			case High:
				highTides = fmt.Sprintf("%s%s, maximum height %v\n", highTides, dt, v.Height)
			case Low:
				lowTides = fmt.Sprintf("%s%s, minumum height %v\n", lowTides, dt, v.Height)
			}
			count += 1
		}
	}
	result <- fmt.Sprintf("Now it is %s (Bangkok time). Latest tides on %s\n\nHigh:\n%s\nLow:\n%s", now.Format("15:04"), coords.Beach, highTides, lowTides)
	return
}

func getCoordsByCommand(command string) *coordinates.Coordinates {
	var coords *coordinates.Coordinates
	switch command {
	case PatongCommand:
		coords = coordinates.GetPatongCoordinates()
	case Aonang:
		coords = coordinates.GetAonangCoordinates()
	case MaiKao:
		coords = coordinates.GetMaiKaoCoordinates()
	}
	return coords
}