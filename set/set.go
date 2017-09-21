package main

import (
	"fmt"
)

type Set map[string]struct{}

func main() {
	s := make(Set)
	s["Item1"] = struct{}{}
	s["Item2"] = struct{}{}
	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string

	for k, _ := range s {
		retVal = append(retVal, k)
	}

	return retVal
}
