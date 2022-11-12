package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type PinConfig struct {
	RowPins []int `json:"rows"`
	ColPins []int `json:"cols"`
}

func LoadConfig(fileName string) PinConfig {
	file, _ := ioutil.ReadFile(fmt.Sprintf("config-files/%s", fileName))

	config := PinConfig{}

	err := json.Unmarshal([]byte(file), &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("col pins")
	for _, pin := range config.ColPins {
		log.Println(pin)
	}

	log.Println("row pins")
	for _, pin := range config.RowPins {
		log.Println(pin)
	}
	return config
}

func (cfg *PinConfig) RowCount() int {
	return len(cfg.RowPins)
}

func (cfg *PinConfig) ColCount() int {
	return len(cfg.ColPins)
}
