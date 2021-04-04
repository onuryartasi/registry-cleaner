package registry

import "time"

type Catalog struct {
	Repositories []string `json:"repositories"`
}

//TODO: change struct name to Image
type Image struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type Manifests struct {
	SchemaVersion int    `json:"schemaVersion"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Architecture  string `json:"architecture"`
	FsLayers      []struct {
		BlobSum string `json:"blobSum"`
	} `json:"fsLayers"`
	History []struct {
		V1Compatibility string `json:"v1Compatibility"`
	} `json:"history"`
}

type V1Compatibility struct {
	Created time.Time `json:"created"`
	ID      string    `json:"id"`
	Parent  string    `json:"parent"`
}

type Tag struct {
	Name        string
	CreatedDate time.Time
	ImageName   string
}

type Registry struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DryRun   bool
}
