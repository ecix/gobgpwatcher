package main

// Get list of neighbours

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func apiListNeighbours(req *http.Request, params httprouter.Params) (ApiResponse, error) {

	// Get list of neighbours
	neighbours, err := GoBGP.ListNeighbor()
	return neighbours, err
}

func apiShowNeighbour(req *http.Request, params httprouter.Params) (ApiResponse, error) {

	// Get a single neighbour
	return "implement me", nil
}
