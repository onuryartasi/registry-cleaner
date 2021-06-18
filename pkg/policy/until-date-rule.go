package policy

import (
	"encoding/json"
	"time"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func (policy Policy) untilDateCheck(client registry.RegistryInterface, image registry.Image, parsedDate time.Time) registry.Image {
	logger := logging.GetLogger()
	var v1Compatibility registry.V1Compatibility
	var deletableTags []string

	for _, tag := range image.Tags {

		manifests := client.GetManifest(image.Name, tag)

		if len(manifests.History) == 0 {
			logger.Warnf("Image Manifest is broken. Skipping this tag. image: %s:%s", image.Name, tag)
			continue
		}

		v1comp := manifests.History[0].V1Compatibility
		err := json.Unmarshal([]byte(v1comp), &v1Compatibility)
		if err != nil {
			logger.Errorf("Error Unmarshal compatibility error:%s", err)
		}

		if parsedDate.After(v1Compatibility.Created) {
			deletableTags = append(deletableTags, tag)
		}

	}

	return registry.Image{Name: image.Name, Tags: deletableTags}
}
