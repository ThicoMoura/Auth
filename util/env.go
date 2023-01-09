package util

import (
	"github.com/spf13/viper"
)

type env struct {
	Source  string `mapstructure:"DBSOURCE"`
	Migrate string `mapstructure:"MIGRATION_URL"`
	Key     string `mapstructure:"TOKEN_KEY"`
	GinMode string `mapstructure:"GIN_MODE"`
	Addr string `mapstructure:"ADDRESS"`
}

func NewEnv(path string) (env *env, err error) {
	vp := viper.New()

	vp.SetConfigFile(path + ".env")

	if err = vp.ReadInConfig(); err != nil {
		return
	}

	if err = vp.Unmarshal(&env); err != nil {
		return
	}

	return
}
