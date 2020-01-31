package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
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
	for now := range time.Tick(time.Second * 10) {
		for idx := range services {
			code := getStatusCode(services[idx].url)
			services[idx].status = code
			AddData(services[idx])
		}
	}
}
