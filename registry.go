package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


func NewClient(host,port string)(Registry){
	return Registry{HOST:host,PORT:port}

}

func (registry *Registry) BasicAuthentication(user,password string){
	*registry = Registry{HOST:registry.HOST,PORT:registry.PORT,USER:user,PASSWORD:password}

}
func (registry Registry)GET(path string)(*http.Response,error){
	resp,err := http.Get(fmt.Sprintf("http://%s:%s%s",registry.HOST,registry.PORT,path))
	return resp,err
}

func (registry Registry) getCatalog()Catalog{
	var catalog Catalog
	resp,err := registry.GET("/v2/_catalog")
	if err != nil {
		log.Println("Error getting version",err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &catalog)

	if err != nil {
		log.Println("Unmarshall error ",err)
	}
	return catalog
}

func splitRepositories(repositories []string)map[string][]string{
	var group,repoName string
	var repoMap = make(map[string][]string)
	for _,repo:= range repositories{
		split := strings.Split(repo,"/")
		if len(split)>1{
			group = split[0]
			repoName = split[1]
		}else{
			group = "other"
			repoName = split[0]
		}
		groupRepo,ok := repoMap[group]
		if ok{
			groupRepo = append(groupRepo,repoName)
			repoMap[group] = groupRepo
		}else {
			repoMap[group] = []string{repoName}
		}
	}
	return repoMap

}


func (registry Registry) getDigest (imageName,tag string) string{
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/v2/%s/manifests/%s",registry.HOST,registry.PORT,imageName,tag)
	req, err := http.NewRequest("HEAD", url , nil)
	if err != nil {
		log.Println("Cannot get docker image digest",err)
	}
	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Cannot get digest",err)
	}
	return resp.Header["Docker-Content-Digest"][0]
}
func (registry Registry) getManifest(imageName,tag string) Manifests {
	var manifests Manifests
	url := fmt.Sprintf("http://%s:%s/v2/%s/manifests/%s",registry.HOST,registry.PORT,imageName,tag)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url , nil)
	if err != nil {
		log.Println("Cannot get docker image digest",err)
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Cannot get digest",err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &manifests)
	if err !=nil{
		log.Println("Unmarshal error mamifes",err)
	}
	return manifests
}


func (registry Registry) getTags(groupName,repoName string) Tag{
	var tags Tag
	url := fmt.Sprintf("http://%s:%s/v2/%s/%s/tags/list",registry.HOST,registry.PORT, groupName, repoName)
	resp, err := http.Get(url)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &tags)
	if err != nil {
		log.Println("Error unmarshal tags",err)
	}
	return tags
}


func (registry Registry) deleteTag(imageName,digest string) int{
	url := fmt.Sprintf("http://%s:%s/v2/%s/manifests/%s",registry.HOST,registry.PORT,imageName,digest)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url , nil)
	if err != nil {
		log.Println("Cannot get docker image digest",err)
	}
	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	if len(registry.USER)>0{
		req.SetBasicAuth(registry.USER,registry.PASSWORD)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Cannot get digest",err)
	}
	return resp.StatusCode
}














