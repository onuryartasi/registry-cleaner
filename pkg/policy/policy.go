package policy

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"

	"gopkg.in/yaml.v2"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

var client = registry.Registry{}
var imageRuleImages *[]Image
var logger *logrus.Logger

func init() {
	logger = logging.GetLogger()
}

func Initialize() Policy {
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

	if policy.RegexRule.Enable {
		image = policy.regexRuleCheck(image)

	}

	if policy.ImageRule.Enable {
		image = policy.imageRuleCheck(image)

	}

	if policy.OlderThanGivenDateRule.Enable {
		image = policy.olderThanGivenDateCheck(image)
	}

	if policy.NRule.Enable {
		image = policy.nRuleCheck(image)

	}

	if client.DryRun {
		logger.Infoln(image)
	} else {
		deleteTags(image)
	}

}

func (policy Policy) setImageRuleImages() {

	var imagess []Image
	for _, rawImage := range policy.ImageRule.Images {
		var image Image
		tag := strings.Split(rawImage, ":")
		if len(tag) > 1 {
			image.tag = tag[1]
		} else {
			image.tag = ""
		}
		image.name = tag[0]
		imagess = append(imagess, image)
	}
	imageRuleImages = &imagess

}

func deleteTags(image registry.Image) {

	for _, tag := range image.Tags {
		digest, err := client.GetDigest(image.Name, tag)
		if err != nil {
			logger.Errorf("Cannot getting Image Digest in deleteTags. Image: %s:%s", image.Name, tag)
			continue
		}

		statusCode, err := client.DeleteTag(image.Name, digest)
		if err != nil {
			logger.Errorf("Cannot delete Tag  status Code:%v, error: %s", statusCode, err)
			continue
		}
		if logger.GetLevel() > logrus.ErrorLevel {
			if statusCode == 202 {
				logger.Infof("Deleted image: %s:%s", image.Name, tag)
			}
		}
		if logger.GetLevel() > logrus.FatalLevel {
			if statusCode != 202 {
				logger.Infof("Cannot Delete image: %s:%s error: %v", image.Name, tag, err)
			}
		}

	}
}
