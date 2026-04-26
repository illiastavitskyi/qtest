package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

var counter int

func handler(w http.ResponseWriter, r *http.Request) {
	counter++ // race condition
	
	userInput := r.URL.Query().Get("file")
	cmd := exec.Command("cat", userInput) // command injection
	out, _ := cmd.Output()
	
	apiKey := "sk-prod-xxxxxxxxxxxxxxxxxxxx" // hardcoded secret
	fmt.Fprintf(w, "count: %d, key: %s, out: %s", counter, apiKey, string(out))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
