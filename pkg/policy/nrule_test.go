package policy

import (
	"strings"
	"testing"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"
	"github.com/onuryartasi/registry-cleaner/pkg/registry"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestNRuleWithTags(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.NRule.Size = 2
	client := registry.RegistryTest{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"linux", "1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.nRuleCheck(client, image)
	assert.ElementsMatch([]string{"linux", "2.0.0"}, result.Tags)
}

func TestNRuleWithLowerSize(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.NRule.Size = 3
	client := registry.RegistryTest{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.nRuleCheck(client, image)
	assert.ElementsMatch([]string{"1.0.0", "2.0.0", "2.0-alpha"}, result.Tags)
}

func TestNRuleWithMissingHistory(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logging.SetLogger(testLogger)
	policy := Policy{}
	policy.NRule.Size = 2
	client := registry.RegistryTest{}
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
	client := registry.RegistryTest{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"unmarshal", "linux", "1.0.0", "2.0.0", "2.0-alpha"}}

	policy.nRuleCheck(client, image)
	if !strings.Contains(hook.LastEntry().Message, "Error Unmarshal compatibility error") {
		t.Errorf("Cannot logged error")
	}
}
