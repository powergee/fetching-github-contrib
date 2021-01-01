package main

import (
	"fmt"
)

const owner = "CSUOS"
const repo = "KOS"

func main() {
	fmt.Println(GetBranches(owner, repo))
}