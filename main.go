package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urls, err := ioutil.ReadFile("Dirs.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, url := range urls {
		resp, err := http.Head("https://www.Teachable.com")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 305 {
			fmt.Println("exists", url)
		} else {
			fmt.Println(url, "does not exist")
		}
	}
}
