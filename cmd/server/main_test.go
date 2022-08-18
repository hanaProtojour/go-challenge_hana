package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	resp, err := http.Get("http://localhost:3333")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
