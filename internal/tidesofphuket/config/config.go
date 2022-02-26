package config

import "tidesOfPhuket/internal/tools"


type Config struct {
	WorldTideApiKey string
	WorldTideUrl    string
	TelegramToken   string
	RedisAddr       string
	RedisPassword   string
}

func NewTidesOfPhuketConfig() (*Config, error) {
	cfg := Config{}
	tools.EnvToString(&cfg.WorldTideApiKey, "WORLD_TIDE_API_KEY", "")
	tools.EnvToString(&cfg.WorldTideUrl, "WORLD_TIDE_URL", "https://www.worldtides.info/api/v3")
	tools.EnvToString(&cfg.TelegramToken, "WORLD_TIDE_TELEGRAM_TOKEN", "")
	tools.EnvToString(&cfg.RedisAddr, "WORLD_TIDE_REDIS_ADDR", "127.0.0.1:6379")
	tools.EnvToString(&cfg.RedisPassword, "WORLD_TIDE_REDIS_PASSWORD", "")
	return &cfg, nil
}