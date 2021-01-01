package main

import "strings"

const URL = "https://api.github.com/"

func ToRequestURL(path ...string) string {
	return URL + strings.Join(path, "/")
}