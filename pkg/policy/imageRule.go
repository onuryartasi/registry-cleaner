package policy

import (
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func (policy Policy) imageRuleCheck(image registry.Image) registry.Image {

	var deletableTags []string

	if policy.ImageRule.Keep {

		for _, v := range *imageRuleImages {
			if v.name == image.Name {

				if v.tag == "" {
					continue
				} else {
					for _, tag := range image.Tags {
						if v.tag == tag {
							continue
						} else {
							deletableTags = append(deletableTags, tag)
						}
					}
				}
			} else {
				deletableTags = append(deletableTags, image.Tags...)
			}
		}
	} else {
		for _, v := range *imageRuleImages {

			if v.name == image.Name {

				if v.tag == "" {
					deletableTags = append(deletableTags, image.Tags...)
				} else {
					for _, tag := range image.Tags {
						if v.tag == tag {
							deletableTags = append(deletableTags, tag)
						} else {
							continue
						}
					}
				}
			} else {
				continue
			}
		}
	}

	if len(deletableTags) > 0 {
		return registry.Image{Name: image.Name, Tags: deletableTags}
	}
	return registry.Image{}
}
