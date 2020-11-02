package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving by %s for URL \"%s\".\n", os.Getenv("SERVICE_NAME"), r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on 8083...")
	http.ListenAndServe(":8083", nil)
}
