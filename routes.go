package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/osrg/gobgp/packet/bgp"
	"github.com/osrg/gobgp/table"
)

func apiShowRoutes(req *http.Request, params httprouter.Params) (ApiResponse, error) {
	neighbourId, err := validateNotEmpty(params.ByName("id"))
	if err != nil {
		return nil, err
	}

	// Get routes
	prefixes := make([]*table.LookupPrefix, 0)
	rib, err := GoBGP.GetLocalRIB(neighbourId, bgp.RF_IPv4_UC, prefixes)
	if err != nil {
		return nil, err
	}

	routes := rib.GetDestinations()

	return routes, nil
}
