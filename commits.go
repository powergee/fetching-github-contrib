package main

import (
	"time"
	"fmt"
)

type Commit struct {
	Message string
	Date time.Time
	Author string
}

func GetCommits(owner string, repo string, branch string, perPage int, page int) []Commit {
	url := ToRequestURL("repos", owner, repo, "commits")
	params := map[string]string{
        "sha": branch,
        "page": fmt.Sprint(page),
        "per_page": fmt.Sprint(perPage),
    }
	res := GetResponse("GET", url, params)
	result := make([]Commit, len(res), len(res))

	for index, history := range res {
		commit, _ := history["commit"].(map[string]interface{})

		author, _ := commit["author"].(map[string]interface{})
		result[index].Author, _ = author["name"].(string)

		dateStr, _ := author["date"].(string)
		result[index].Date, _ = time.Parse(time.RFC3339, dateStr)

		message, _ := commit["message"].(string)
		result[index].Message = message
	}

	return result
}