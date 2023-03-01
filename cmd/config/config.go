package config

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const defaultConfigPath = "cmd/config/config.yml"

func LoadConfig() error {
	viper.SetConfigFile(defaultConfigPath)
	err := viper.ReadInConfig()
	log.Info("Loading config variables..")
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	return err
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
