package app

import (
	"log"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	Host       string `env:"HOST" envDefault:"0.0.0.0"`
	Port       int    `env:"PORT" envDefault:"8080"`
	Address    string `env:"ADDRESS,expand" envDefault:"$HOST:${PORT}"`
	APIPrefix  string `env:"API_PREFIX" envDefault:"/"`
	MongodbURI string `env:"MONGODB_URI"`
	RedisURL   string `env:"REDIS_URL"`
	MySqlDNS   string `env:"MYSQL_DNS"`
	JWTSecret  string `env:"JWT_SECRET"`
}

func (app *App) initAppConfig() {
	cfg, err := env.ParseAsWithOptions[AppConfig](env.Options{
		RequiredIfNoDef: true,
	})

	if err != nil {
		log.Panicln("Configuration error", err)
	}

	app.cfg = &cfg
}
