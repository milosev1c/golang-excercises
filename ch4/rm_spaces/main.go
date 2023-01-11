package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func removeSpaces(runes []byte) []byte {
	res := runes[:0]
	var last rune
	for i := 0; i < len(runes); {
		rn, s := utf8.DecodeRune(runes[i:])
		if !unicode.IsSpace(rn) {
			res = append(res, runes[i:i+s]...)
		} else if unicode.IsSpace(rn) && !unicode.IsSpace(last) {
			res = append(res, ' ')
		}
		last = rn
		i += s
	}
	return res
}

func main() {
	str := []byte("dd dda d dd    dwadad")
	fmt.Println(string(removeSpaces(str)))
}
