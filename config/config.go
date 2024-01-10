package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP struct {
		Port         int    `yaml:"port" env-default:"4000"`
		Host         string `yaml:"host" env-default:"localhost"`
		ReadTimeout  string `yaml:"read_timeout" env-default:"10s"`
		IdleTimeout  string `yaml:"idle_timeout" env-default:"60s"`
		WriteTimeout string `yaml:"write_timeout" env-default:"30s"`
	} `yaml:"http"`

	ENV string `yaml:"env" env-default:"development"`

	DB struct {
		DSN              string `env:"APP_DB_DSN" env-required:"true"`
		MaxOpenConns     int    `yaml:"max_open_conns" env-default:"25"`
		MaxIdleConns     int    `yaml:"max_idle_conns" env-default:"25"`
		MaxIdleTime      string `yaml:"max_idle_time" env-default:"15m"`
		MigrationPath    string `yaml:"migration_path"`
		MigrationVersion uint   `yaml:"migration_version"`
	} `yaml:"db"`

	Token struct {
		SecretKey  string        `env:"TOKEN_SECRET_KEY" env-required:"true"`
		TimeToLive time.Duration `yaml:"time_to_live"`
	} `yaml:"token"`

	// limiter struct {
	// 	rps     float64
	// 	burst   int
	// 	enabled bool
	// }

	// smtp struct {
	// 	host     string
	// 	port     int
	// 	username string
	// 	password string
	// 	sender   string
	// }

	// cors struct {
	// 	trustedOrigins []string
	// }
}

func InitConfig(path string) (*Config, error) {
	config := new(Config)
	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
