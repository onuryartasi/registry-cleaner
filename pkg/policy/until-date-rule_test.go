package policy

import (
	"log"
	"strings"
	"testing"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"

	"github.com/sirupsen/logrus/hooks/test"

	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func TestUntilDate(t *testing.T) {
	//assert := assert.New(t)
	policy := Policy{}

	policy.UntilDateRule.Date = "05.03.2021 04:04:05"
	client := registry.RegistryTest{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"linux", "1.0.0", "2.0.0", "2.0-alpha"}}
	parsedDate, err := parseDate("05.03.2021 04:04:05")
	if err != nil {
		log.Println("Parsed Date error")
	}
	result := policy.untilDateCheck(client, image, parsedDate)
	log.Println(result)
	//assert.ElementsMatch([]string{"linux", "2.0.0"}, result.Tags)
}

func TestUntilDateWithMissingHistory(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logging.SetLogger(testLogger)
	policy := Policy{}

	client := registry.RegistryTest{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"missing-history", "linux", "1.0.0", "2.0.0", "2.0-alpha"}}
	parsedDate, err := parseDate("05.03.2021 04:04:05")
	if err != nil {
		log.Println("Parsed Date error")
	}
	policy.untilDateCheck(client, image, parsedDate)
	if !strings.Contains(hook.LastEntry().Message, "Image Manifest is broken. Skipping this tag") {
		t.Errorf("Cannot logged error")
	}

}

func TestUntilDateWithUnmarshalError(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logging.SetLogger(testLogger)
	policy := Policy{}
	client := registry.RegistryTest{}
	image := registry.Image{Name: "foo/bar", Tags: []string{"unmarshal", "linux", "1.0.0", "2.0.0", "2.0-alpha"}}

	parsedDate, err := parseDate("05.03.2021 04:04:05")
	if err != nil {
		log.Println("Parsed Date error")
	}
	policy.untilDateCheck(client, image, parsedDate)
	if !strings.Contains(hook.LastEntry().Message, "Error Unmarshal compatibility error") {
		t.Errorf("Cannot logged error")
	}
}
