package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const URL = "https://api.github.com/"
var IS_TOKEN_SET bool = false
var TOKEN *string = nil

type ProcessedResponse struct {
	Header 	map[string][]string
	Body	[]map[string]interface{}
}

func ToRequestURL(path ...string) string {
	return URL + strings.Join(path, "/")
}

func getToken() *string {
	if !IS_TOKEN_SET {
		IS_TOKEN_SET = true
		err := godotenv.Load()

		if err != nil {
			fmt.Println("Failed to load .env: " + err.Error())
			return nil
		}

		token, found := os.LookupEnv("GITHUB_TOKEN")
		if !found {
			fmt.Println("There's no GITHUB_TOKEN in .env: " + err.Error())
			return nil
		}
		TOKEN = &token
		return TOKEN
	} else {
		return TOKEN
	}
}

func GetResponse(method string, reqUrl string, params map[string]string) ProcessedResponse {
	req, _ := http.NewRequest(method, reqUrl, nil)

	if token := getToken(); token != nil {
		req.Header.Add("Authorization", "token " + *token)
	}

	q := req.URL.Query()
	q.Add("accept", "application/vnd.github.v3+json")
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, _ := http.DefaultClient.Do(req)
	resBody, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	bodyString := string(resBody)
	if bodyString[0] != '[' {
		bodyString = "[" + bodyString + "]"
	}
	
	var parsed []map[string]interface{}
	json.Unmarshal([]byte(bodyString), &parsed)
	
	return ProcessedResponse { res.Header, parsed }
}