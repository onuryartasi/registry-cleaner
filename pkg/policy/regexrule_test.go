package policy

import (
	"testing"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func TestRegexRuleCheck(t *testing.T) {
	policy := Policy{}
	policy.RegexRule.Pattern = []string{"foo/bar"}

	image := registry.Image{Name: "foo/bars"}
	actually := policy.regexRuleCheck(image)

	if actually.Name != image.Name {
		t.Errorf("regexRuleCheck failed. expected %v, got %v", image.Name, actually.Name)
	}

}
