package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/docker/client/opts"
	"io/ioutil"
	"net/http"
)

type APIClient struct{}
type CmdJsonBody struct {
	Cmd string `json:"cmd"`
}

func (client *APIClient) ImageBuild() error {
	url := "http://localhost:" + opts.DefaultHTTPPort + "/commands"
	bodyStruct := CmdJsonBody{
		Cmd: "build",
	}
	bodyJson, _ := json.Marshal(bodyStruct)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		panic(err)
	}
	req.Header = http.Header{
		"Content-type": []string{"application/json"},
	}

	httpClient := new(http.Client)
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArray))
	return nil
}
