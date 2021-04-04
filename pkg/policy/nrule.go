package policy

import (
	"encoding/json"
	"log"
	"sort"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func (policy Policy) nRuleCheck(image registry.Image) registry.Image {

	var tagList []registry.Tag
	var v1Compatibility registry.V1Compatibility
	var deletableTags []string

	if len(image.Tags) > policy.NRule.Size {
		for _, tag := range image.Tags {

			manifests := client.GetManifest(image.Name, tag)
			//v1comp,err := strconv.Unquote(manifests.History[0].V1Compatibility)
			if len(manifests.History) == 0 {
				log.Println("Image Manifest is broken.Skipping this tag.", image.Name, tag)
				continue
			}

			v1comp := manifests.History[0].V1Compatibility
			err := json.Unmarshal([]byte(v1comp), &v1Compatibility)
			if err != nil {
				log.Println("Error Unmarshal compatibility ", err)
			}

			digest, err := client.GetDigest(image.Name, tag)
			if err != nil {
				logger.Errorf("Cannot get digest in nrule")
			}
			tagList = append(tagList, registry.Tag{Name: tag, CreatedDate: v1Compatibility.Created, Digest: digest, ImageName: image.Name})
			//log.Printf("Image %s, Tag: %s, created date: %v", image.Name ,tag, startedTime.Sub(v1Compatibility.Created).Hours())
		}

		sort.SliceStable(tagList, func(i, j int) bool {
			return tagList[i].CreatedDate.After(tagList[j].CreatedDate)
		})

		//Return deletable tags
		lastTags := tagList[policy.NRule.Size:]
		for _, tag := range lastTags {
			deletableTags = append(deletableTags, tag.Name)
		}
		return registry.Image{Name: image.Name, Tags: deletableTags}
	}
	return registry.Image{}
}
