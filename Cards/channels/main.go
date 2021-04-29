package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
		c <- "Hi"
	}

	go func() {
		fmt.Println("Hi")
	}()

	// for l := range c {
	// 	go func(x string) {

	// 		time.Sleep(5 * time.Second)
	// 		checkLink(x, c)
	// 	}(l)
	// }
}

func checkLink(link string, c chan string) {
	fmt.Println(link, "All Good!")

	fmt.Println(<-c)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error : ", err)
		//c <- link
		return
	}
	fmt.Println(link, "All Good!")
	//c <- link
}
