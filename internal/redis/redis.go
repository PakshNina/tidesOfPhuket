package redis

import (
	"encoding/json"
	"fmt"
	"time"

	r "github.com/go-redis/redis"
	"gopkg.in/errgo.v2/fmt/errors"

	apiClient "tidesOfPhuket/internal/tidesofphuket/client"
	"tidesOfPhuket/internal/tidesofphuket/config"
)

var ErrKeyDoesNotExistsInRedis = errors.New("error with getting value from redis")

type Redis interface {
	SaveTideInfo(beach string, tidesInfo *apiClient.TidesExtremeResult) error
	GetExtremesForToday(beach string) (*apiClient.TidesExtremeResult, error)
}

type Client struct {
	redis *r.Client
}

func NewClient(cfg *config.Config) *Client {
	opt := r.Options{
		Addr: cfg.RedisAddr,
		DB:   1,
	}
	c := r.NewClient(&opt)
	return &Client{redis: c}
}

func (c Client) SaveTideInfo(beach string, tidesInfo *apiClient.TidesExtremeResult) error {
	key := getRedisKeyForToday(beach)
	byteData, errJson := json.Marshal(tidesInfo)
	if errJson != nil {
		return errors.Newf("Error with encoding key `%s`, value `%v` to JSON: %v", key, tidesInfo, errJson)
	}
	if errSet := c.redis.Set(key, byteData, 24*time.Hour).Err(); errSet != nil {
		return errors.Newf("Error with saving key `%s`, value `%v` to redis: %v", key, tidesInfo, errSet)
	}
	return nil
}

func getRedisKeyForToday(beach string) string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	dt := time.Now().In(loc).Format("2006-01-02")
	return fmt.Sprintf("%s_%s", beach, dt)
}

func (c Client) GetExtremesForToday(beach string) (*apiClient.TidesExtremeResult, error) {
	dateKey := getRedisKeyForToday(beach)
	jsonData, err := c.redis.Get(dateKey).Result()
	if err != nil {
		if err == r.Nil {
			return nil, ErrKeyDoesNotExistsInRedis
		}
		return nil, err
	}
	var arr apiClient.TidesExtremeResult
	jsonErr := json.Unmarshal([]byte(jsonData), &arr)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &arr, nil
}
