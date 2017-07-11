package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/osrg/gobgp/table"
)

func apiShowRoutes(req *http.Request, params httprouter.Params) (ApiResult, error) {
	neighbourId, err := validateNotEmpty(params.ByName("id"))
	if err != nil {
		return nil, err
	}

	// Get route family
	family := routeFamilyFromAddr(neighbourId)

	// Get routes
	prefixes := make([]*table.LookupPrefix, 0)
	rib, err := GoBGP.GetLocalRIB(neighbourId, family, prefixes)
	if err != nil {
		return nil, err
	}

	routes := rib.GetDestinations()

	return routes, nil
}
