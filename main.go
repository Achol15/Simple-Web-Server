package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! Everyone Whatsup? Welcome to my very first golang web server. ")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting Server on :9090...")
	if err := http.ListenAndServe(":9090", nil); err !=nil {
		fmt.Println("Error Starting Server:", err)
	}

}