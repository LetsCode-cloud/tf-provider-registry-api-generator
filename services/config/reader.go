package config

import (
	"encoding/json"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/types"
	"log"
	"os"
)

func ReadConfig(cfgFile string) types.Config {
	c := types.Config{}
	file, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Printf("file.Get err#%v ", err)
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
