package policy

import (
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)



func (policy Policy) imageRuleCheck(image registry.Image)  []Image{
	var deletableImage  []Image
		group,name := registry.SplitImage(image.Name)
		for _,v := range *imageRuleImages {
			if v.group == group && v.name == name  {

				if v.tag == "" {

					if policy.ImageRule.Keep{
						continue
					}else {
						for _, tag := range image.Tags {
							deletableImage = append(deletableImage,Image{name: name,group: group,tag: tag})
						}
					}
				} else {
					for _,tag := range image.Tags{
						if v.tag == tag {
							if policy.ImageRule.Keep{
								continue
							}else {
								deletableImage = append(deletableImage,Image{name: name,group: group,tag: v.tag})
							}
						}
					}
				}
			}
		}
	return deletableImage
}
