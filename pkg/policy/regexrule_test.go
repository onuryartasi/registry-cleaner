package policy

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func TestRegexRuleCheckWithName(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.RegexRule.Pattern = []string{"foo/bar"}

	image := registry.Image{Name: "foo/bar", Tags: []string{"latest", "1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.regexRuleCheck(image)
	assert.ElementsMatch([]string{"latest", "1.0.0", "2.0.0", "2.0-alpha"}, result.Tags)

}

func TestRegexRuleCheckWithTags(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.RegexRule.Pattern = []string{"foo/bar:2.*"}

	image := registry.Image{Name: "foo/bar", Tags: []string{"latest", "1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.regexRuleCheck(image)
	assert.ElementsMatch([]string{"2.0.0", "2.0-alpha"}, result.Tags)

}

func TestRegexRuleCheckNotMatch(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.RegexRule.Pattern = []string{"foo/xyz"}

	image := registry.Image{Name: "foo/bar", Tags: []string{"latest", "1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.regexRuleCheck(image)
	assert.ElementsMatch([]string{}, result.Tags)

}

func TestRegexRuleCheckNilPattern(t *testing.T) {
	assert := assert.New(t)
	policy := Policy{}
	policy.RegexRule.Pattern = []string{""}

	image := registry.Image{Name: "foo/bar", Tags: []string{"latest", "1.0.0", "2.0.0", "2.0-alpha"}}

	result := policy.regexRuleCheck(image)
	assert.ElementsMatch([]string{}, result.Tags)
}
