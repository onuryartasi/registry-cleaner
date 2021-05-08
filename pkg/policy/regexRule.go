package policy

import (
	"fmt"
	"log"
	"regexp"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func (policy Policy) regexRuleCheck(image registry.Image) registry.Image {
	var deletableTags []string
	for _, pattern := range policy.RegexRule.Pattern {

		//match, err :=regexp.MatchString(pattern,image.Name)
		r, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatalf("regex error compile: %s", err)
		}

		if r.MatchString(image.Name) {
			deletableTags = append(deletableTags, image.Tags...)
		} else {
			for _, tag := range image.Tags {
				if r.MatchString(fmt.Sprintf("%s:%s", image.Name, tag)) {
					deletableTags = append(deletableTags, tag)
				}
			}
		}
	}

	if len(deletableTags) > 0 {
		return registry.Image{Name: image.Name, Tags: deletableTags}
	}
	return image
}
