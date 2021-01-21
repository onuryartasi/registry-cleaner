package policy

import (
	"encoding/json"
	"github.com/onuryartasi/registry-management/pkg/registry"
	"log"
	"sort"
	"time"
)

func (policy Policy) nRuleCheck(image registry.Tag) {
	var tagList []registry.SortTag
	var v1Compatibility registry.V1Compatibility
	var startedTime = time.Now()

	if len(image.Tags) > policy.NRule.Keep {
		for _, tag := range image.Tags {

			manifests := client.GetManifest(image.Name, tag)
			//v1comp,err := strconv.Unquote(manifests.History[0].V1Compatibility)
			if len(manifests.History) == 0 {
				log.Println("Image Manifest is broken.Skippin this tag.", image.Name, tag)
				continue
			}

			v1comp := manifests.History[0].V1Compatibility
			err := json.Unmarshal([]byte(v1comp), &v1Compatibility)
			if err != nil {
				log.Println("Error Unmarshal compatibility ", err)
			}

			digest := client.GetDigest(image.Name, tag)
			tagList = append(tagList, registry.SortTag{Tag: tag, TimeAgo: startedTime.Sub(v1Compatibility.Created).Hours(), Digest: digest, Name: image.Name})
			log.Printf("Image: %s, created date: %s", tag, startedTime.Sub(v1Compatibility.Created).Hours())
		}

		sort.SliceStable(tagList, func(i, j int) bool {
			return tagList[i].TimeAgo < tagList[j].TimeAgo
		})

		lastTags := tagList[policy.NRule.Keep:]
		log.Println(lastTags)

	}
}