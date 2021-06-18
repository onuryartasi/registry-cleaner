package main

import (
	"flag"

	"github.com/onuryartasi/registry-cleaner/pkg/policy"
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func main() {
	var host = flag.String("host", "localhost", "Registry host")
	var username = flag.String("username", "", "Registry username")
	var password = flag.String("password", "", "Registry password")
	//var lastImages = *flag.Int("keep", 10, "Keep Last n images")
	var dryRun = flag.Bool("dry-run", false, "Print old images, don't remove.")
	flag.Parse()

	policy := policy.Initialize()

	client := registry.NewClient(*host, *username, *password, *dryRun)
	client.CheckAuth()

	policy.Start(client)

}
