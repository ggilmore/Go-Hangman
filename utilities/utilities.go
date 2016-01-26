package utilities

import (
	"fmt"
	"strings"
)

type Set map[rune]bool

func Join(runes []rune) string{
	stringSlice := make([]string, len(runes))
	for i, r := range runes {
		stringSlice[i] = fmt.Sprintf("%c", r)
	}
	return strings.Join(stringSlice, "")
}

func SetCopy(source Set) *Set{
	dest := make(Set)
	for k, v := range source{
		dest[k] = v
	}
	return &dest
}