package policy

import (
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
	"testing"
)

func TestRegexRuleCheck(t *testing.T){
	policy := Policy{}
	policy.RegexRule.Pattern = []string{"foo/bar"}

	image := registry.Image{Name: "foo/bars"}
	type Image struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	actually := policy.regexRuleCheck(image)

	if actually.Name != image.Name {
		t.Errorf("regexRuleCheck failed. expected %v, got %v",image.Name,actually.Name)
	}

}
