package config

import (
	"log"

	"github.com/spf13/viper"
)

func Get() *M {
	var V = &M{}

	viper.AutomaticEnv()
	viper.RegisterAlias(EnvHerokuSlug, "HEROKU_APP_NAME")
	viper.RegisterAlias(EnvHerokuDataBaseHost, "DATABASE_URL")
	viper.SetDefault(EnvHerokuPort, "3000")

	if err := viper.Unmarshal(V); err != nil {
		log.Fatalf("read config: %w", err)
	}

	return V
}
