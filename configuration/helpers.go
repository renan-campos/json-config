package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load() (Configuration, error) {
	var configuration Configuration
	if len(os.Args) != 2 {
		return configuration, fmt.Errorf("configuration json expected as command line argument")
	}
	filename := os.Args[1]

	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return configuration, err
	}

	err = json.Unmarshal(jsonData, &configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil
}
