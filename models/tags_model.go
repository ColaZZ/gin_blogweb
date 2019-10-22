package models

import "strings"

func HandleTagsListData(tags []string) map[string]int {
	var tagMap map[string]int
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList{
			tagMap[value]++
		}
	}
	return tagMap
}
