package policy

import (
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
			log.Printf("regex error compile: %s", err)
		}

		match := r.MatchString(image.Name)
		if match {
			deletableTags = append(deletableTags, image.Tags...)
		}
	}

	if len(deletableTags) > 0 {
		return registry.Image{Name: image.Name, Tags: deletableTags}
	}
	return registry.Image{}
}
