package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func GetBranches(owner string, repo string) []string {
	res, _ := http.Get(ToRequestURL("repos", owner, repo, "branches"))
	resBody, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var parsed []map[string]interface{}
	json.Unmarshal([]byte(resBody), &parsed)
	result := make([]string, len(parsed), len(parsed))
	
	for index, branch := range parsed {
		result[index], _ = branch["name"].(string)
	}

	return result
}