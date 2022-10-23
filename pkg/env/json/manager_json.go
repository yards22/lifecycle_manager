package json

import (
	"encoding/json"
	"os"
)

type ManagerConfigsReader struct {
	configs map[string]ManagerConfigs
}

func New(path string) (*ManagerConfigsReader, error) {
	var jr ManagerConfigsReader
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &jr.configs)
	if err != nil {
		return nil, err
	}
	return &jr, nil
}

func (jr *ManagerConfigsReader) Get(key string) ManagerConfigs {
	return jr.configs[key]
}
