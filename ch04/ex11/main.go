/*
	Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
	invoking their preferred text editor when substantial text input is required.
*/
// There is no delete option in the Github API

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// const targetURL = "https://api.github.com/repos/zzossig/test/issues/1" // patch
// const targetURL = "https://api.github.com/repos/zzossig/test/issues/1" // get
const targetURL = "https://api.github.com/repos/zzossig/test/issues" // post
const token = "your github api token"

func main() {
	if err := create(); err != nil {
		panic(err)
	}

	// b, err := read();
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", b)

	// if err := update(); err != nil {
	// 	panic(err)
	// }
}

func create() error {
	client := &http.Client{}
	
	params := make(map[string]interface{})
	params["title"] = "Test Title"
	params["body"] = "Test Body"
	b, err := json.Marshal(params)
	if err != nil {
		return err
	}
	
	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "bearer " + token)
	
	if _, err := client.Do(req); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func read() ([]byte, error) {
	client := &http.Client{}
	
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "bearer " + token)
	
	result, err := client.Do(req);

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return b, nil
}

func update() error {
	client := &http.Client{}
	
	params := make(map[string]interface{})
	params["title"] = "Updated Title"
	params["body"] = "Updated Body"
	b, err := json.Marshal(params)
	if err != nil {
		return err
	}
	
	req, err := http.NewRequest("PATCH", targetURL, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "bearer " + token)
	
	if _, err := client.Do(req); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
