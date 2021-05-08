package policy

import (
	"encoding/json"
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

			// Broken manifest causes error. If tag's manifest is broken then continue next tag.
			if len(manifests.History) == 0 {
				logger.Warnf("Image Manifest is broken.Skipping this tag. image: %s:%s", image.Name, tag)
				continue
			}

			v1comp := manifests.History[0].V1Compatibility

			err := json.Unmarshal([]byte(v1comp), &v1Compatibility)
			if err != nil {
				logger.Errorf("Error Unmarshal compatibility error:%s", err)
			}

			// Collect tags in slice for sorting.
			tagList = append(tagList, registry.Tag{Name: tag, CreatedDate: v1Compatibility.Created, ImageName: image.Name})
		}

		// sort slice with tag's CreatedDate.
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
	return image
}
