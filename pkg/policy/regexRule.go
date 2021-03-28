package policy

import (
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
	"log"
	"regexp"
)

func (policy Policy) regexRuleCheck(image registry.Image) registry.Image {
	var deletableTags []string
	for _, pattern := range policy.RegexRule.Pattern {

		//match, err :=regexp.MatchString(pattern,image.Name)
		r, err := regexp.Compile(pattern)
		if err != nil {
			log.Printf("regex error compile: %s", err)
		}

		match := r.MatchString(image.Name)
		if match {
			for _, tag := range image.Tags {
				deletableTags = append(deletableTags, tag)
			}
		}
	}

	if len(deletableTags) > 0 {
		return registry.Image{Name: image.Name, Tags: deletableTags}
	}
	return registry.Image{}
}
