package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	var pingUrl string
	flag.StringVar(&pingUrl, "u", "", "url to GET")
	var timeout int
	flag.IntVar(&timeout, "t", 5, "timeout between pings (seconds)")
	var jwtPath string
	flag.StringVar(&jwtPath, "o", "bin/xposer.jwt", "path to signed jwt")
	flag.Parse()
	jwt, err := ioutil.ReadFile(jwtPath)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{Timeout: time.Duration(5) * time.Second}
	req, err := http.NewRequest("GET", pingUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+string(jwt))

	for {
		resp, err := client.Do(req)
		if err != nil {
			log.Print(fmt.Sprintf("Failed to connect: %s", err))
		} else {
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				log.Print(fmt.Sprintf("Failed to connect: %s", body))
			}
		}
		time.Sleep(time.Duration(timeout) * time.Second)
	}
}