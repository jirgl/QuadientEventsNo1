package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/*
base.go file contains basic http requests which are necessary for this task.
*/

const baseURL = "tasks-rad.quadient.com:8080"

func parseBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func callGet(endpoint string) ([]byte, error) {
	response, err := http.Get("http://" + baseURL + "/" + endpoint)
	if err != nil {
		return nil, err
	}

	return parseBody(response)
}

func callPut(endpoint string, data []byte) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", "http://"+baseURL+"/"+endpoint, bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return parseBody(response)
}
