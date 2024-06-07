package main

import (
	"net/http"
)

func handleReq(w http.ResponseWriter, r *http.Request) {
	respond(w, 200, struct{}{})
}
