package configuration

import (
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

type Configuration struct {
	AuthToken string
}

var configInstance *Configuration

func (c *Configuration) GetInstance() *Configuration {
	if configInstance != nil {
		return configInstance
	}

	lock.Lock()
	defer lock.Unlock()

	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("AUTH_TOKEN is undefined")
	}

	configInstance = &Configuration{AuthToken: authToken}

	return configInstance
}
