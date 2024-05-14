package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	resp, err := http.Get("https://animechan.xyz/api/quotes/anime?title=naruto")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	animeChan := []Animechan{}
	err = json.Unmarshal(body, &animeChan)
	if err != nil {
		return nil, err
	}

	return animeChan, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	url := "https://postman-echo.com/post"
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return Postman{}, err
	}
	defer resp.Body.Close()

	var result Postman
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Postman{}, err
	}

	return result, nil
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
