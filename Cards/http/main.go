package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct{

}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// read, err1 :=resp.Body.Read(bs)
	// fmt.Println(read, err1)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(p []byte)(n int, err error){
	fmt.Println(string(p))
	fmt.Println("Just did it")
	return len(p), nil
}
