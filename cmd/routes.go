package main

import "net/http"

func routes() http.Handler {
	m := http.NewServeMux()

	return m
}
