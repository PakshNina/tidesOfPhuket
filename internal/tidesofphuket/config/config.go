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
	tools.EnvToString(&cfg.WorldTideApiKey, "WORLD_TIDE_API_KEY", "be177abf-daca-4d7e-9dc7-20cb2acfa982")
	tools.EnvToString(&cfg.WorldTideUrl, "WORLD_TIDE_URL", "https://www.worldtides.info/api/v3")
	tools.EnvToString(&cfg.TelegramToken, "WORLD_TIDE_TELEGRAM_TOKEN", "5090763874:AAFVybH2N-31hyc5nS_-u22YQxuYQH6_Mjc")
	tools.EnvToString(&cfg.RedisAddr, "WORLD_TIDE_REDIS_ADDR", "127.0.0.1:6379")
	tools.EnvToString(&cfg.RedisPassword, "WORLD_TIDE_REDIS_PASSWORD", "")
	return &cfg, nil
}
