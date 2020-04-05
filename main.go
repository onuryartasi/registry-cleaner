package main

import (
	"encoding/json"
	"github.com/cheggaaa/pb/v3"
	"log"
	"sort"
	"time"
)

func main(){

	registry := NewClient("host","port")
	registry.BasicAuthentication("user","pass")


	var v1Compatibility v1Compatibility


	catalog := registry.getCatalog()

	log.Printf("Founded %d image names",len(catalog.Repositories))
	repoMap := splitRepositories(catalog.Repositories)

	for gN,rL := range repoMap{
		log.Printf("Getting image from group: %s",gN)
		for _,v := range rL {
			tags := registry.getTags(gN,v)
			var tagList []SortTag
			if len(tags.Tags) > 10 {
				log.Printf("Getting tags from %s image tags: %d\n",tags.Name,len(tags.Tags))
				count := len(tags.Tags)
				bar := pb.StartNew(count)
				for _, tag := range tags.Tags {
					bar.Increment()
					manifests := registry.getManifest(tags.Name,tag)
					//v1comp,err := strconv.Unquote(manifests.History[0].V1Compatibility)
					v1comp := manifests.History[0].V1Compatibility
					err := json.Unmarshal([]byte(v1comp), &v1Compatibility)
					if err != nil {
						log.Println("Error Unmarshal compatibility ",err)
					}
					//log.Printf("Image Name:%s, Image Tag: %s, Created Date: %s",tags.Name,v,v1Compatibility.Created.Format("2006-01-02"))
					digest := registry.getDigest(tags.Name,tag)
					tagList = append(tagList, SortTag{Tag: tag, TimeAgo: time.Now().Sub(v1Compatibility.Created).Hours(),Digest:digest})
				}
				bar.Finish()
				sort.SliceStable(tagList, func(i, j int) bool {
					return tagList[i].TimeAgo < tagList[j].TimeAgo
				})

				lastTags := tagList[10:]
				log.Println(len(lastTags))
				//Remove old image keep last 10
				for _,image := range lastTags{
					statusCode := registry.deleteTag(tags.Name,image.Digest)
					if statusCode == 202 {
						log.Printf("%s image's %s tag's removed",tags.Name,image.Tag)
					}
				}
				//----
			}
		}
	}

}