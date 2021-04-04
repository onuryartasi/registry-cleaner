package main

import (
	"flag"

	"github.com/onuryartasi/registry-cleaner/pkg/logging"
	"github.com/onuryartasi/registry-cleaner/pkg/policy"
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func main() {
	var host = flag.String("host", "localhost", "Registry host")
	var port = flag.String("port", "5000", "Registry Port")
	var username = flag.String("username", "", "Registry username")
	var password = flag.String("password", "", "Registry password")
	//var lastImages = *flag.Int("keep", 10, "Keep Last n images")
	var dryRun = flag.Bool("dry-run", false, "Print old images, don't remove.")
	var groupName = *flag.String("group", "", "Remove images from group")
	flag.Parse()

	logger := logging.GetLogger()

	policy := policy.Initialize()

	var isAllGroup = false

	//TODO: Add GroupName list and just apply policy this groups
	if len(groupName) == 0 {
		isAllGroup = true
	}

	client := registry.NewClient(*host, *port, *dryRun)
	client.BasicAuthentication(*username, *password)

	//var v1Compatibility registry.V1Compatibility
	catalog := client.GetCatalog()

	logger.Infof("Founded %d unique images", len(catalog.Repositories))

	repoMap := registry.SplitRepositories(catalog.Repositories)

	//TODO: Get group part by part instead of all
	for gN, rL := range repoMap {
		if isAllGroup || gN == groupName {
			for _, v := range rL {
				image := client.GetImageTags(gN, v)
				logger.Infoln(image)
				policy.Apply(client, image)
			}
		}
	}
}
