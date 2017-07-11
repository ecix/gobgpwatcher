package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Get the gobgpd status and provide
// additional information about the gobgpwatcher

func apiShowStatus(req *http.Request, params httprouter.Params) (ApiResponse, error) {

	server, err := GoBGP.GetServer()

	status := map[string]interface{}{
		"gobgp": server,
		"api": map[string]interface{}{
			"version": "0.1.0",
		},
	}

	return status, err
}
