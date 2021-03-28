package policy

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
	"gopkg.in/yaml.v2"
)

var client = registry.Registry{}
var imageRuleImages *[]Image

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

	policy.setImageRuleImages()

	return policy
}

//TODO: Change deletableImages to registry.Image type
//type Image struct {
//	Name string   `json:"name"`
//	Tags []string `json:"tags"`
//}

func (policy Policy) Apply(cl registry.Registry, image registry.Image) {
	//var deleteImage = true
	client = cl
	//var deletableImages []Image


	if policy.RegexRule.Enable == true {
		//todo: Write RegexRule function and call it.
		// deleteImage = deleteImage && RegexRuleFunctionResult
		 deletableImages := policy.regexRuleCheck(image)
		 log.Println("Deletable Regex Image",deletableImages)
	}

	if policy.ImageRule.Enable == true {
		//todo: Write ImageRule function and call it.
		//todo: deleteImage = deleteImage && ImageRuleFunctionResul

		 //imageRuleDeletable := policy.imageRuleCheck(image)
		 //deletableImages = append(deletableImages,imageRuleDeletable[:]...)

	}


	//if policy.OlderThanGivenDateRule.Enable == true {
	//	//todo: Write OlderThanGivenDateRule function and call it.
	//	//todo: deleteImage = deleteImage && NRuleFunctionResult
	//}
	//

	if policy.NRule.Enable {
		policy.nRuleCheck(image)
		//todo: Write NRule function and call it.
		//todo: deleteImage = deleteImage && NRuleFunctionResult
	}



	//if deleteImage == true {
	//	//todo: call delete image function
	//}
}


func (policy Policy) setImageRuleImages() {

	var imagess []Image
	for _, repo := range policy.ImageRule.Images {
		var image Image
		tag := strings.Split(repo, ":")
		if len(tag) > 1 {
			image.tag = tag[1]
		}else{
			image.tag = ""
		}
		image.name = tag[0]
	}
	imageRuleImages = &imagess

}