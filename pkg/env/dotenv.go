package env

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

func ViperGetEnvVar(key string) string {
	envPath := filepath.Join(".env")
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
