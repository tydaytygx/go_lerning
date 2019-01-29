package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(rs http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(rs, "URL.Path = %q\n", rq.URL.Path)

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
