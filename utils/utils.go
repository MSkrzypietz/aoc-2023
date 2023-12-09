package utils

import (
	"strconv"
	"strings"
)

func IntFields(s string) []int {
	var result []int
	for _, field := range strings.Fields(s) {
		intField, _ := strconv.Atoi(field)
		result = append(result, intField)
	}
	return result
}
