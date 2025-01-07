package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/rivo/uniseg" //nolint:depguard
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	result := strings.Builder{}
	var prevSymbol []rune
	var currentSymbol []rune

	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		currentSymbol = gr.Runes()

		if len(currentSymbol) == 1 && unicode.IsNumber(currentSymbol[0]) &&
			(len(prevSymbol) == 0 || unicode.IsNumber(prevSymbol[0])) {
			return "", ErrInvalidString
		}

		if (!unicode.IsNumber(currentSymbol[0])) && len(prevSymbol) > 0 && !unicode.IsNumber(prevSymbol[0]) {
			result.WriteRune(prevSymbol[0])
		}

		if unicode.IsNumber(currentSymbol[0]) && len(prevSymbol) > 0 && prevSymbol[0] != 0 {
			for t, _ := strconv.Atoi(string(currentSymbol)); t > 0; t-- {
				result.WriteRune(prevSymbol[0])
			}
		}
		_, b := gr.Positions()

		if !unicode.IsNumber(currentSymbol[0]) && len(s) == b {
			result.WriteRune(currentSymbol[0])
		}

		prevSymbol = currentSymbol
	}

	return result.String(), nil
}
