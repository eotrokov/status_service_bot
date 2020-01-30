package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	os.Setenv("services", "http://ya.ru")
	env := os.Getenv("services")
	if len(env) == 0 {
		fmt.Println("нет данных")
		return
	}
	urls := strings.Split(env, " ")

	services := make([]service, len(urls))
	for index, url := range urls {
		services[index] = service{status: 0, url: url}
	}
	fmt.Println(services)
	run(services)
}

func run(services []service) {
	for now := range time.Tick(time.Second * 2) {
		for idx := range services {
			code := getStatusCode(services[idx].url)
			if code != services[idx].status {
				services[idx].status = code
				AddData(services[idx])
				fmt.Println(services[idx].url, "status", services[idx].status, http.StatusText(services[idx].status), now)
			}
		}
	}
}
