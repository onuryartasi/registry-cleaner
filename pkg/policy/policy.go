package policy

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"

	"gopkg.in/yaml.v2"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

var parsedDate time.Time
var logger *logrus.Logger

func init() {
	logger = logging.GetLogger()
}

// initialize convert config to struct.
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

	parsedDate, err = parseDate(policy.OlderThanGivenDateRule.Date)
	if err != nil {
		logger.Fatalf("Cannot parse given date with layout. Check layout table...")
	}

	return policy
}

// apply perform polices given config on unique images.
func (policy Policy) Apply(client registry.Registry, image registry.Image) {

	if policy.RegexRule.Enable {
		image = policy.regexRuleCheck(image)
	}

	if policy.OlderThanGivenDateRule.Enable {
		image = policy.olderThanGivenDateCheck(client, image, parsedDate)
	}

	if policy.NRule.Enable {
		image = policy.nRuleCheck(client, image)
	}

	if client.DryRun {
		logger.Infoln("Deleting images: ", image)
	} else {
		deleteTags(client, image)
	}

}

func (policy Policy) Start(client registry.Registry) {
	catalog := client.GetCatalog()
	logger.Infof("Founded %d unique images", len(catalog.Repositories))
	repoMap := registry.SplitRepositories(catalog.Repositories)

	//TODO: Get group part by part instead of all
	for gN, rL := range repoMap {
		for _, v := range rL {
			image := client.GetImageTags(gN, v)
			policy.Apply(client, image)
		}
	}
}

// deleteTags  get tags' digest given tags and delete them.
func deleteTags(client registry.RegistryInterface, image registry.Image) {

	for _, tag := range image.Tags {
		digest, err := client.GetDigest(image.Name, tag)
		if err != nil {
			logger.Errorf("Cannot getting Image Digest in deleteTags. Image: %s:%s", image.Name, tag)
			continue
		}
		statusCode, err := client.DeleteTag(image.Name, digest)
		if err != nil {
			logger.Errorf("Cannot delete Tag statusCode:%v, error: %s", statusCode, err)
			continue
		}

		if statusCode == 202 {
			logger.Infof("Deleted image: %s:%s", image.Name, tag)
		} else {
			logger.Warnf("Cannot Delete image: %s:%s error: %v", image.Name, tag, err)
		}
	}
}

// parseDate convert time.Time type given date in config
func parseDate(date string) (time.Time, error) {
	var err error
	var parsedDate time.Time
	var layouts = []string{
		"02.01.2006 15:04:05",
		"02.01.2006 15:04",
		"02.01.2006",
	}

	for _, layout := range layouts {
		parsedDate, err = time.Parse(layout, date)
		if err != nil {
			continue
		}
		return parsedDate, nil
	}
	return time.Time{}, err
}
