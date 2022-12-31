package main

import "strings"

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	words := strings.Fields(input)

	return words
}
