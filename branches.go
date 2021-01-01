package main

func GetBranches(owner string, repo string) []string {
	url := ToRequestURL("repos", owner, repo, "branches")
	res := GetResponse("GET", url, nil)
	result := make([]string, len(res), len(res))
	
	for index, branch := range res {
		result[index], _ = branch["name"].(string)
	}

	return result
}