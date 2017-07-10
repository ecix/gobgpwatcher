package main

/*
 GoBGPwatcher

 Use the gobgp grpc api to provide a
 minimal rest / http API for use with Alice.
*/

import (
	"net/http"

	"fmt"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/osrg/gobgp/client"
)

var GoBGP *client.Client

func main() {

	fmt.Println("GoBGPwatcher                      v0.0.1")
	fmt.Println("  - Listening on:  0.0.0.0:50023")

	// Make client instance
	cli, err := client.New("")
	if err != nil {
		log.Fatal("Could not create gobgp client:", err)
	}
	GoBGP = cli

	router := httprouter.New()

	// Register API routes
	apiRegisterEndpoints(router)

	// Start http server
	log.Fatal(http.ListenAndServe("0.0.0.0:50023", router))
}
