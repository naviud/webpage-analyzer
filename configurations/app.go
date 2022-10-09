package configurations

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type appConfig struct {
	ChannelCount    int           `yaml:"channelCount"`
	LinkTimeoutInMs time.Duration `yaml:"linkTimeoutInMs"`
	ServicePort     string        `yaml:"servicePort"`
}

var config *appConfig

func GetAppConfig() *appConfig {
	if config == nil {
		config = &appConfig{}
	}
	return config
}

func (a *appConfig) Register() {
	file, err := ioutil.ReadFile("configs/app.yaml")
	if err != nil {
		log.Fatal("Error in file open", err)
	}
	err = yaml.Unmarshal(file, a)
	if err != nil {
		log.Fatal("Error in decoding", err)
	}

	a.LinkTimeoutInMs = a.LinkTimeoutInMs * time.Millisecond

	log.Printf("App config : %+v", a)
}
