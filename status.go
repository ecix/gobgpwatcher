package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Get the gobgpd status and provide
// additional information about the gobgpwatcher

func apiShowStatus(req *http.Request, params httprouter.Params) (ApiResponse, error) {

	return map[string]string{
		"version": "23.42",
		"foo":     "implement me.",
	}, nil
}
