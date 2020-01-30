package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type service struct {
	status int
	url    string
}

func main() {
	urls := strings.Split(os.Getenv("services"), " ")
	services := []service{}
	for _, url := range urls {
		services = append(services, service{status: 0, url: url})
		fmt.Println(services)
	}
	run(services)
}

func run(services []service) {
	for now := range time.Tick(time.Second * 20) {
		for idx := range services {
			code := getStatusCode(services[idx].url)
			if code != services[idx].status {
				services[idx].status = code
				fmt.Println(services[idx].url, " status ", services[idx].status, http.StatusText(services[idx].status), now)
			}
		}
	}
}

func getStatusCode(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode
}
