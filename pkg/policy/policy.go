package policy

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

func Init(){
	policy := Policy{}
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Cannot read config file")
	}

	err = yaml.Unmarshal([]byte(data), &policy)
	if err != nil {
		log.Println("Cannot unmarshal yaml file to struct")
	}

}
