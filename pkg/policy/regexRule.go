package policy

import (
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
	"log"
	"regexp"
)

func (policy Policy) regexRuleCheck(image registry.Image) []Image {
	deletableImages := []Image{}
	for _, pattern := range policy.RegexRule.Pattern {

		//match, err :=regexp.MatchString(pattern,image.Name)
		r, err := regexp.Compile(pattern)
		if err != nil {
			log.Printf("regex error compile: %s", err)
		}

		match := r.MatchString(image.Name)
		if match {
			for _, tag := range image.Tags {
				deletableImages = append(deletableImages, Image{name: image.Name, tag: tag})
			}
		}
	}
	return deletableImages
}
