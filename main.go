package main

import (
	"flag"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"
	"github.com/onuryartasi/registry-cleaner/pkg/policy"
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func main() {
	var host, port, username, password, groupName string
	var dryRun bool
	flag.StringVar(&host, "host", "localhost", "Registry host")
	flag.StringVar(&port, "port", "5000", "Registry Port")
	flag.StringVar(&username, "username", "", "Registry username")
	flag.StringVar(&password, "password", "", "Registry password")
	flag.StringVar(&groupName, "group", "", "Remove images from group")
	//var lastImages = *flag.Int("keep", 10, "Keep Last n images")
	flag.BoolVar(&dryRun, "dry-run", false, "Print old images, don't remove.")

	flag.Parse()

	logger := logging.GetLogger()

	policy := policy.Initialize()

	var isAllGroup = false

	//TODO: Add GroupName list and just apply policy this groups
	if len(groupName) == 0 {
		isAllGroup = true
	}

	client := registry.NewClient(host, port, dryRun)
	client.BasicAuthentication(username, password)

	catalog := client.GetCatalog()

	logger.Infof("Founded %d unique images", len(catalog.Repositories))

	repoMap := registry.SplitRepositories(catalog.Repositories)

	//TODO: Get group part by part instead of all
	for gN, rL := range repoMap {
		if isAllGroup || gN == groupName {
			for _, v := range rL {
				image := client.GetImageTags(gN, v)
				policy.Apply(client, image)
			}
		}
	}
}
