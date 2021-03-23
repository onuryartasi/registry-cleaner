package main

import (
	"flag"
	"log"

	"github.com/onuryartasi/registry-cleaner/pkg/policy"
	"github.com/onuryartasi/registry-cleaner/pkg/registry"
)

func main() {
	var host = flag.String("host", "localhost", "Registry host")
	var port = flag.String("port", "80", "Registry Port")
	var username = flag.String("username", "", "Registry username")
	var password = flag.String("password", "", "Registry password")
	//var lastImages = *flag.Int("keep", 10, "Keep Last n images")
	//var dryRun = *flag.Bool("dry-run",false,"Print old images, don't remove.")
	var groupName = *flag.String("group", "", "Remove images from group")
	flag.Parse()

	policy := policy.Initiliaze()
	//policy.Apply(Image)

	//var startedTime = time.Now()
	var isAllGroup = false

	if len(groupName) == 0 {
		isAllGroup = true
	}

	client := registry.NewClient(*host, *port)
	client.BasicAuthentication(*username, *password)

	//var v1Compatibility registry.V1Compatibility
	catalog := client.GetCatalog()
	log.Printf("Founded %d unique images", len(catalog.Repositories))
	repoMap := registry.SplitRepositories(catalog.Repositories)

	for gN, rL := range repoMap {
		if isAllGroup || gN == groupName {
			for _, v := range rL {

				tags := client.GetTags(gN, v)

				policy.Apply(client, tags)
				//var tagList []registry.SortTag
				//
				//if len(tags.Tags) > lastImages {
				//	log.Printf("Getting image from group: %s",gN)
				//	log.Printf("Getting tags from %s image tags: %d\n",tags.Name,len(tags.Tags))
				//	for _, tag := range tags.Tags {
				//
				//		manifests := client.GetManifest(tags.Name,tag)
				//		//v1comp,err := strconv.Unquote(manifests.History[0].V1Compatibility)
				//		if len(manifests.History) == 0 {
				//			log.Println("Image Manifest is broken.Skippin this tag.",tags.Name,tag)
				//			continue
				//		}
				//
				//		v1comp := manifests.History[0].V1Compatibility
				//		err := json.Unmarshal([]byte(v1comp), &v1Compatibility)
				//		if err != nil {
				//			log.Println("Error Unmarshal compatibility ",err)
				//		}
				//
				//		digest := client.GetDigest(tags.Name,tag)
				//		tagList = append(tagList, registry.SortTag{Tag: tag, TimeAgo: startedTime.Sub(v1Compatibility.Created).Hours(),Digest:digest,Name: tags.Name})
				//		log.Printf("Image: %s, created date: %s",tag,startedTime.Sub(v1Compatibility.Created).Hours())
				//	}
				//
				//	sort.SliceStable(tagList, func(i, j int) bool {
				//		return tagList[i].TimeAgo < tagList[j].TimeAgo
				//	})
				//
				//	lastTags := tagList[lastImages:]
				//	log.Println(len(lastTags))

				////Remove old image keep last 10
				//for _,image := range lastTags{
				//	if dryRun {
				//		log.Printf("Image: %s, tagName: %s",gN+image.Name,image.Tag)
				//	} else {
				//
				//		ok := policy.Filter(image)
				//		statusCode := client.DeleteTag(tags.Name,image.Digest)
				//		if statusCode == 202 {
				//			log.Printf("%s image's %s tag's removed",tags.Name,image.Tag)
				//		} else{
				//			log.Printf("Error to delete image. Status Code: %d",statusCode)
				//		}
				//	}
				//}
			}
		}
	}
}
