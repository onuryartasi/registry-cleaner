package policy

import (
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
	"log"
)

func (policy Policy) imageRuleCheck(image registry.Image) registry.Image {

	var deletableTags []string

	if policy.ImageRule.Keep {

		for _, v := range *imageRuleImages {
			log.Println("xyz2 ", v.name, image.Name)
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
				for _, tag := range image.Tags {
					deletableTags = append(deletableTags, tag)
				}
			}
		}
	} else {
		for _, v := range *imageRuleImages {

			if v.name == image.Name {

				if v.tag == "" {
					for _, tag := range image.Tags {
						deletableTags = append(deletableTags, tag)
					}
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
