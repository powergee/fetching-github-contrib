package main

import (
	"fmt"
	"strings"
	"time"
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
	result := make([]Commit, len(res.Body), len(res.Body))

	for index, history := range res.Body {
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

func getHashOfLatestCommit(owner string, repo string, branch string) string {
	url := ToRequestURL("repos", owner, repo, "commits")
	params := map[string]string{
        "sha": branch,
        "page": "1",
        "per_page": "1",
    }
	res := GetResponse("GET", url, params)

	return res.Body[0]["sha"].(string)
}

func getHashOfOldestCommit(owner string, repo string, branch string) string {
	url := ToRequestURL("repos", owner, repo, "commits")
	params := map[string]string{
        "sha": branch,
        "page": "1",
        "per_page": "1",
    }
	res := GetResponse("GET", url, params)

	if links, exist := res.Header["Link"]; exist {
		lastLinks := links[len(links) - 1]
		apiCallForLast := lastLinks[strings.LastIndex(lastLinks, "<") + 1:strings.LastIndex(lastLinks, ">")]
		res = GetResponse("GET", apiCallForLast, nil)
	}

	return res.Body[len(res.Body) - 1]["sha"].(string)
}

func CountAllCommits(owner string, repo string, branch string) int {
	oldest := getHashOfOldestCommit(owner, repo, branch)
	latest := getHashOfLatestCommit(owner, repo, branch)

	url := ToRequestURL("repos", owner, repo, "compare", oldest + "..." + latest)
	res := GetResponse("GET", url, nil)

	return int(res.Body[0]["total_commits"].(float64) + 1)
}