package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// NewActivity is the struct used for creating new activities, which do not have an id yet
type NewActivity struct {
	Type     string `json:"type"`
	Activity string `json:"activity"`
	Ref      string `json:"ref"`
}

func logActivity(t string, activity string, ref string) {

	act := NewActivity{t, activity, ref}
	var buf = new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&act)

	fmt.Println("***** URL: *****", activitiesUrl)

	req, err := http.NewRequest("POST", activitiesUrl, buf)
	if err != nil {
		log.Panic("CANNOT CREATE HTTP REQUEST")
		fmt.Println(err)
		panic(err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("CANNOT EXECUTE HTTP REQUEST")
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("request URL:", activitiesUrl)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
