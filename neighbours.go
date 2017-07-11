package main

// Get list of neighbours

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func apiListNeighbours(req *http.Request, params httprouter.Params) (ApiResult, error) {

	// Get list of neighbours
	neighbours, err := GoBGP.ListNeighbor()
	return neighbours, err
}

func apiShowNeighbour(req *http.Request, params httprouter.Params) (ApiResult, error) {
	id, err := validateNotEmpty(params.ByName("id"))
	if err != nil {
		return nil, err
	}

	// Get a single neighbour
	neighbour, err := GoBGP.GetNeighbor(id)
	return neighbour, err
}
