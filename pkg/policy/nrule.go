package policy

import (
	"encoding/json"
	"log"
	"sort"
	"time"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func (policy Policy) nRuleCheck(image registry.Image) registry.Image {
	var tagList []registry.SortTag
	var v1Compatibility registry.V1Compatibility
	var startedTime = time.Now()
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

			digest := client.GetDigest(image.Name, tag)
			tagList = append(tagList, registry.SortTag{Tag: tag, TimeAgo: startedTime.Sub(v1Compatibility.Created).Hours(), Digest: digest, Name: image.Name})
			//log.Printf("Image %s, Tag: %s, created date: %v", image.Name ,tag, startedTime.Sub(v1Compatibility.Created).Hours())
		}

		sort.SliceStable(tagList, func(i, j int) bool {
			return tagList[i].TimeAgo < tagList[j].TimeAgo
		})

		//Return deletable tags
		lastTags := tagList[0:policy.NRule.Size]
		for _, v := range lastTags {
			deletableTags = append(deletableTags, v.Tag)
		}
		log.Println(deletableTags)
		return registry.Image{Name: image.Name, Tags: deletableTags}
	}
	return registry.Image{}
}
