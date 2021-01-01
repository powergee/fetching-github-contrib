package main

import (
	"fmt"
)

const owner = "CSUOS"
const repo = "KOS"

func main() {
	fmt.Println("Fetching CSUOS/KOS")

	branches := GetBranches(owner, repo)
	fmt.Println("Branches:", branches)
	fmt.Println()

	for i, branch := range branches {
		fmt.Printf("[Branch #%v: Recent 3 commits of %v]\n", i+1, branch)
		commits := GetCommits(owner, repo, branch, 3, 1)

		for j, commit := range commits {
			fmt.Println("- Message", j+1)
			fmt.Println(commit.Message)
			fmt.Println("- Author", j+1)
			fmt.Println(commit.Author)
			fmt.Println("- Date", j+1)
			fmt.Println(commit.Date)
			fmt.Println()
		}
	}
}