package policy

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

func Initiliaze() Policy{
	policy := Policy{}

	//todo: Add Environment variable for policy file.
	//todo: Read policy file to path from arguments.
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Cannot read config file")
	}

	err = yaml.Unmarshal([]byte(data), &policy)
	if err != nil {
		log.Println("Cannot unmarshal yaml file to struct")
	}

	return policy
}


func (policy Policy) filter(image string) bool{
	var deleteImage = true
	if policy.ImageRule.Enable == true {
		//todo: Write ImageRule function and call it.
		//todo: deleteImage = deleteImage && ImageRuleFunctionResult
	}

	if policy.NRule.Enable == true {
		//todo: Write NRule function and call it.
		//todo: deleteImage = deleteImage && NRuleFunctionResult
	}

	if policy.OlderThanGivenDateRule.Enable == true {
		//todo: Write OlderThanGivenDateRule function and call it.
		//todo: deleteImage = deleteImage && NRuleFunctionResult
	}

	if policy.RegexRule.Enable == true {
		//todo: Write RegexRule function and call it.
		// deleteImage = deleteImage && RegexRuleFunctionResult
	}

	if deleteImage == true {
		//todo: call delete image function
	}
	return true
}