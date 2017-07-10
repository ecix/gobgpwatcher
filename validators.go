package main

// Validate query parameters

import (
	"fmt"
)

func validateNotEmpty(param string) (string, error) {
	if param == "" {
		return param, fmt.Errorf("Missing or empty parameter")
	}

	return param, nil
}
