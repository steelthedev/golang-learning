package main

import (
	"strings"
)

func getInitails(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	initials := []string{}

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
