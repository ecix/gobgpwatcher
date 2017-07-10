package main

import (
	"net"

	"github.com/osrg/gobgp/packet/bgp"
)

func routeFamilyFromAddr(addr string) bgp.RouteFamily {
	ip := net.ParseIP(addr)
	if ip == nil {
		return bgp.RouteFamily(0)
	}

	if ip.To4() != nil {
		return bgp.RF_IPv4_UC
	}

	if ip.To16() != nil {
		return bgp.RF_IPv6_UC
	}

	return bgp.RouteFamily(0)
}
