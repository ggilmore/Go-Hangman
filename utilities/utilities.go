package utilities

import (
	"fmt"
	"strings"
)

/*
random collection of helper functions and values
 */


const ALPHABET = "abcdefghijklmnopqrstuvwxyz "

type RuneSet map[rune]bool

var SETOFLETTERS = genLetterSet() //size of alphabet

func genLetterSet() RuneSet {
	letterSet := RuneSet{}
	for _, l := range ALPHABET {
		letterSet[l] = true
	}
	return letterSet
}


func RuneToStr(r rune) string{
	return fmt.Sprintf("%c", r)
}



func JoinRunes(runes []rune) string{
	stringSlice := make([]string, len(runes))
	for i, r := range runes {
		stringSlice[i] = RuneToStr(r)
	}
	return strings.Join(stringSlice, "")
}

func RuneSetCopy(source RuneSet) *RuneSet {
	dest := make(RuneSet)
	for k, v := range source{
		dest[k] = v
	}
	return &dest
}