package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("arg expected")
		os.Exit(1)
	}

	http.HandleFunc("/httpTestFile", func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.Open("httpTestFile")
		defer file.Close()
		io.Copy(w, file)
	})
	err := http.ListenAndServe(os.Args[1], nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
