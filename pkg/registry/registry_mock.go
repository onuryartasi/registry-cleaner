package registry

import "log"

type RegistryTest struct{}

func (client RegistryTest) GetManifest(imageName, tag string) Manifests {
	tags := make(map[string]string)
	tags["linux"] = `{"created":"2021-03-05T01:25:25.230064203Z"}`
	tags["1.0.0"] = `{"created":"2021-03-05T04:25:25.230064203Z"}`
	tags["2.0.0"] = `{"created":"2021-03-05T03:03:25.230064203Z"}`
	tags["2.0-alpha"] = `{"created":"2021-03-05T22:12:00.232064203Z"}`
	tags["missing-history"] = ""
	tags["unmarshal"] = `{"created":":wrong-time"}`

	date := tags[tag]
	v1 := []struct{ V1Compatibility string }{{V1Compatibility: date}}
	manifest := Manifests{
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

func (client RegistryTest) GetCatalog() Catalog {
	return Catalog{Repositories: []string{""}}
}

func (client RegistryTest) DeleteTag(imageName, digest string) (int, error) {
	log.Println(imageName, digest)
	return 200, nil
}

func (client RegistryTest) GetDigest(imageName, tag string) (string, error) {
	log.Println(imageName, tag)
	return "", nil
}
