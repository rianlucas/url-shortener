package main

import (
	"net/http"
)

func main() {
	r := http.NewServeMux()

	err := http.ListenAndServe("3000", r)
	if err != nil {
		return
	}
}
