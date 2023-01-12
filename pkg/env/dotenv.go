package env

import (
	"log"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/spf13/viper"
)

func ViperGetEnvVar(key string) string {
	os := runtime.GOOS
	re := regexp.MustCompile("windows")
	var envPath string
	switch {
	case re.MatchString(os):
		envPath = filepath.Join("../../", ".env")
	default:
		envPath = filepath.Join(".env")
	}
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
