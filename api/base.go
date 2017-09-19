package api

import (
	"io/ioutil"
	"net/http"
)

const baseURL = "tasks-rad.quadient.com:8080"

func callGet(endpoint string) ([]byte, error) {
	resp, err := http.Get("http://" + baseURL + "/" + endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func callPut() {

}
