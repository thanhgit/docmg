package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SiteMap []struct {
	Type     string `json:"type"`
	Name     string `json:"name,omitempty"`
	Contents []struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Content  string `json:"content"`
		Contents []struct {
			Type     string        `json:"type"`
			Name     string        `json:"name"`
			Content  string 		`json:"content"`
			Contents []interface{} `json:"contents,omitempty"`
		} `json:"contents"`
	} `json:"contents,omitempty"`
	Directories int `json:"directories,omitempty"`
	Files       int `json:"files,omitempty"`
}

func GetSiteMap(path string) *SiteMap  {
	sitemap := SiteMap{}
	file, _ := ioutil.ReadFile(path)
	json.Unmarshal([]byte(file), &sitemap)

	return &sitemap
}

func SaveSiteMap(sitemap SiteMap)  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("Error: %v", r)
		}
	}()
	data, err := json.Marshal(&sitemap)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(data))
}

func (sitemap *SiteMap)GetFile(paths []string) (name string, content string) {
	depth := len(paths)
	for  _, item := range *sitemap {
		if item.Name == paths[0] {
			for index:=0; index < depth; index++ {

			}
		}
	}
	return
}

func (sitemap *SiteMap) AddFile()  {

}

func (sitemap *SiteMap)	UpdateFile()  {

}

func (sitemap *SiteMap) DeleteFile() {

}

func (sitemap *SiteMap) MoveFile()  {

}

