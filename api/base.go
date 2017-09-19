package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func callPut(endpoint, data string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", "http://"+baseURL+"/"+endpoint, strings.NewReader(data))
	//request.SetBasicAuth("admin", "admin")
	//request.ContentLength = 23
	response, err := client.Do(request)
	var content []byte
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(content)
	}

	return content, nil
}
