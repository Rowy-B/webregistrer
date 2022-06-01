package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	login()
}

func register() {
	url := "http://localhost:8080/api/register/"
	method := "POST"

	payload := strings.NewReader(`{
	  "name":"tim",
	  "role":"doctor",
	  "email":"tim@mail.com",
	  "password":"password"
  
  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

type feedback struct {
	Msg   string `json:"msg"` //Hier wordt de teruggestuurde data opgeslagen
	Token string `json:"token"`
}

func login() {

	url := "http://localhost:8080/api/signin/"
	method := "POST"

	payload := strings.NewReader(`{
	  "email":"rowy@mail.com",
	  "password":"password"
  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var data feedback
	json.Unmarshal(body, &data)
	fmt.Println("dit is de data", data.Msg) //hier wordt de data opgeroepen uit de struct

}
