package main

import "net/http"

func Response_readiness(w http.ResponseWriter, r *http.Request) {
	ResponsewithJSON(w, 200, struct{}{})
}
