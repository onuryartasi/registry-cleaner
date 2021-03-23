package policy

import (
	"io/ioutil"
	"log"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
	"gopkg.in/yaml.v2"
)

var client = registry.Registry{}

func Initiliaze() Policy {
	policy := Policy{}

	//todo: Add Environment variable for policy file.
	//todo: Read policy file to path from arguments.
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Cannot read config file")
	}

	err = yaml.Unmarshal(data, &policy)
	if err != nil {
		log.Println("Cannot unmarshal yaml file to struct")
	}

	return policy
}

func (policy Policy) Apply(cl registry.Registry, image registry.Tag) {
	//var deleteImage = true
	client = cl

	//if policy.RegexRule.Enable == true {
	//	//todo: Write RegexRule function and call it.
	//	// deleteImage = deleteImage && RegexRuleFunctionResult
	//}

	//if policy.ImageRule.Enable == true {
	//	//todo: Write ImageRule function and call it.
	//	//todo: deleteImage = deleteImage && ImageRuleFunctionResult
	//}
	//

	if policy.NRule.Enable {
		policy.nRuleCheck(image)
		//todo: Write NRule function and call it.
		//todo: deleteImage = deleteImage && NRuleFunctionResult
	}

	//if policy.OlderThanGivenDateRule.Enable == true {
	//	//todo: Write OlderThanGivenDateRule function and call it.
	//	//todo: deleteImage = deleteImage && NRuleFunctionResult
	//}
	//

	//if deleteImage == true {
	//	//todo: call delete image function
	//}
}
