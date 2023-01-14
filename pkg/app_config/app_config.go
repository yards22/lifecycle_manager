package app_config

import (
	"log"
	"path"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

var Data = koanf.New(".")

func init() {
	os := runtime.GOOS
	re := regexp.MustCompile("windows")
	var basePath string
	switch {
	case re.MatchString(os):
		basePath = filepath.Join("../../", "")
	default:
		basePath = filepath.Join("")
	}

	if err := Data.Load(file.Provider(path.Join(basePath, "runner_duration.json")), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := Data.Load(file.Provider(path.Join(basePath, ".env")), dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

}
