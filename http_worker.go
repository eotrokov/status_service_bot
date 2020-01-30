package main

import (
	"log"
	"net/http"
)

func getStatusCode(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode
}
