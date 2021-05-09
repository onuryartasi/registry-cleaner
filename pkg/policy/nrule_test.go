package policy

import (
	"strings"
	"testing"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"
	"github.com/onuryartasi/registry-cleaner/pkg/registry"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

type testclient struct{}

func TestNRuleWithTags(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.NRule.Size = 2
	client := testclient{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"linux", "1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.nRuleCheck(client, image)
	assert.ElementsMatch([]string{"linux", "2.0.0"}, result.Tags)
}

func TestNRuleWithLowerSize(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.NRule.Size = 3
	client := testclient{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.nRuleCheck(client, image)
	assert.ElementsMatch([]string{"1.0.0", "2.0.0", "2.0-alpha"}, result.Tags)
}

func TestNRuleWithMissingHistory(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logging.SetLogger(testLogger)
	policy := Policy{}
	policy.NRule.Size = 2
	client := testclient{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"missing-history", "linux", "1.0.0", "2.0.0", "2.0-alpha"}}

	policy.nRuleCheck(client, image)
	if !strings.Contains(hook.LastEntry().Message, "Image Manifest is broken. Skipping this tag") {
		t.Errorf("Cannot logged error")
	}

}

func TestNRuleWithUnmarshalError(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logging.SetLogger(testLogger)
	policy := Policy{}
	policy.NRule.Size = 2
	client := testclient{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"unmarshal", "linux", "1.0.0", "2.0.0", "2.0-alpha"}}

	policy.nRuleCheck(client, image)
	if !strings.Contains(hook.LastEntry().Message, "Error Unmarshal compatibility error") {
		t.Errorf("Cannot logged error")
	}
}

func (client testclient) GetManifest(imageName, tag string) registry.Manifests {
	tags := make(map[string]string)
	tags["linux"] = `{"created":"2021-03-05T01:25:25.230064203Z"}`
	tags["1.0.0"] = `{"created":"2021-03-05T04:25:25.230064203Z"}`
	tags["2.0.0"] = `{"created":"2021-03-05T03:03:25.230064203Z"}`
	tags["2.0-alpha"] = `{"created":"2021-03-05T22:12:00.232064203Z"}`
	tags["missing-history"] = ""
	tags["unmarshal"] = `{"created":":wrong-time"}`

	date := tags[tag]
	v1 := []struct{ V1Compatibility string }{{V1Compatibility: date}}
	manifest := registry.Manifests{
		SchemaVersion: 1,
		Name:          imageName,
		Tag:           tag,
		Architecture:  "amd64",
		FsLayers:      nil,
		History:       nil,
	}

	manifest.History = []struct {
		V1Compatibility string `json:"v1Compatibility"`
	}(v1)

	if len(date) == 0 {
		manifest.History = nil
	}

	return manifest
}

func (client testclient) GetCatalog() registry.Catalog {
	return registry.Catalog{Repositories: []string{""}}
}

func (client testclient) GetDigest(imageName, tag string) (string, error) {
	return "", nil
}

func (client testclient) DeleteTag(imageName, digest string) (int, error) {
	return 200, nil
}
