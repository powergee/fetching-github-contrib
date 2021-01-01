package main

import (
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const URL = "https://api.github.com/"

func ToRequestURL(path ...string) string {
	return URL + strings.Join(path, "/")
}

func GetResponse(method string, reqUrl string, params map[string]string) []map[string]interface{} {
	req, _ := http.NewRequest(method, reqUrl, nil)

	q := req.URL.Query()
	q.Add("accept", "application/vnd.github.v3+json")
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, _ := http.DefaultClient.Do(req)
	resBody, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var parsed []map[string]interface{}
	json.Unmarshal([]byte(resBody), &parsed)
	return parsed
}