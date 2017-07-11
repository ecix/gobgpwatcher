package main

import (
	"compress/gzip"
	"encoding/json"
	"net/http"

	"log"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// GoBGP http API
//
// This API provides some read-only endpoints
// for querying GoBGP via a HTTP interface.
//
// This is intended for use in a looking glass integration
// context.
//
// Endpoints:
//
//    Neighbours
//       List          /v1/neighbours
//       Show          /v1/neighbours/:id
//
//    Routes
//       Show Id       /v1/neighbours/:id/routes
//

type ErrorResponse struct {
	Error string `json: "error"`
}

type ApiResult interface{}
type apiEndpoint func(*http.Request, httprouter.Params) (ApiResult, error)

// Wrap handler for access controll, throtteling and compression
func endpoint(wrapped apiEndpoint) httprouter.Handle {
	return func(res http.ResponseWriter,
		req *http.Request,
		params httprouter.Params) {

		// Get result from handler
		result, err := wrapped(req, params)
		if err != nil {
			result = ErrorResponse{
				Error: err.Error(),
			}
			payload, _ := json.Marshal(result)
			http.Error(res, string(payload), http.StatusInternalServerError)
			return
		}

		// Encode json
		payload, err := json.Marshal(result)
		if err != nil {
			msg := "Could not encode result as json"
			http.Error(res, msg, http.StatusInternalServerError)
			log.Println(err)
			log.Println("This is most likely due to an older version of go.")
			log.Println("Consider upgrading to golang > 1.8")
			return
		}

		// Set response header
		res.Header().Set("Content-Type", "application/json")

		// Check if compression is supported
		if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			// Compress response
			res.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(res)
			defer gz.Close()
			gz.Write(payload)
		} else {
			res.Write(payload) // Fall back to uncompressed response
		}
	}
}

// Register api endpoints
func apiRegisterEndpoints(router *httprouter.Router) {
	router.GET("/v1/status", endpoint(apiShowStatus))

	router.GET("/v1/neighbours", endpoint(apiListNeighbours))
	router.GET("/v1/neighbours/:id", endpoint(apiShowNeighbour))

	router.GET("/v1/neighbours/:id/routes", endpoint(apiShowRoutes))
}
