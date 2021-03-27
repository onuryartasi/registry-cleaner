package policy

import (
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)



func (policy Policy) imageRuleCheck(image registry.Image)  []Image{
	var deletableImage  []Image
		group,name := registry.SplitImage(image.Name)

		if policy.ImageRule.Keep  {
			for _,v := range *imageRuleImages {
				if v.group == group && v.name == name  {
					if v.tag == "" {
							continue
						} else {
							for _, tag := range image.Tags {
								if v.tag == tag {
										continue
									}else {
										deletableImage = append(deletableImage,Image{name: name,group: group,tag: tag})
									}
								}
							}
						}else {
								for _, tag := range image.Tags {
									deletableImage = append(deletableImage,Image{name: name,group: group,tag: tag})
								}
						}
					}
				}else{
					for _,v := range *imageRuleImages {
					if v.group == group && v.name == name  {
						if v.tag == "" {
							for _, tag := range image.Tags {
								deletableImage = append(deletableImage,Image{name: name,group: group,tag: tag})
							}
						} else {
							for _, tag := range image.Tags {
								if v.tag == tag {
									deletableImage = append(deletableImage,Image{name: name,group: group,tag: tag})
								}else {
									continue
								}
							}
						}
					} else {
						continue
					}
				}
		}
	return deletableImage
}
