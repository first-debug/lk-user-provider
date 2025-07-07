package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Env    string `env:"ENV" env-default:"local"`
	URL    string `yaml:"url"`
	Port   string `yaml:"port" env-default:"80"`
	DB_URL string `env:"DB_URL"`
	Logger struct {
		Level        *slog.Level `yaml:"level"`
		ShowPathCall bool        `yaml:"show_path_call" env-default:"false"`
	} `yaml:"logger"`
	Shutdown struct {
		Period     time.Duration `yaml:"period" env-default:"15s"`
		HardPeriod time.Duration `yaml:"hard_period" env-default:"3s"`
	} `yaml:"shutdown"`
	Readiness struct {
		DrainDelay time.Duration `yaml:"drain_delay" env-default:"5s"`
	} `yaml:"readiness"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(configPath); err != nil {
		panic(err)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic(err.Error())
	}
	return cfg

}

func fetchConfigPath() (res string) {
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = "config/config_local.yml"
	}
	return
}
