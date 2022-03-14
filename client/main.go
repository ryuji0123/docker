package main

import (
	"docker/client/opts"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:" + opts.DefaultHTTPPort
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArray))
}
