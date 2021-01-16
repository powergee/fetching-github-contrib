package main

import (
	"fmt"
)

const owner = "CSUOS"
const repo = "KOS"
const recent = 3

func main() {
	fmt.Println("Fetching CSUOS/KOS")

	branches := GetBranches(owner, repo)
	fmt.Println("Branches:", branches)
	fmt.Println()

	for i, branch := range branches {
		fmt.Printf("[Branch #%v: Recent %v commits of %v]\n", i+1, recent, branch)
		fmt.Println("Total count of commits to", branch, ":", CountAllCommits(owner, repo, branch))
		commits := GetCommits(owner, repo, branch, recent, 1)

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