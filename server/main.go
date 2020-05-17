package main

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {
	for _, header := range req.Header {
		for _, h := range header {
			fmt.Fprint(w, h)
		}
	}

	fmt.Fprint(w, "Hello")
}

func main() {
	http.HandleFunc("/hello", HelloWorldHandler)
	http.ListenAndServe(":8090", nil)

}
