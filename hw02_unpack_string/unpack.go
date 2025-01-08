package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	result := strings.Builder{}

	var lastSymbol rune

	for idx, symbol := range s {

		if unicode.IsNumber(symbol) && (idx == 0 || unicode.IsNumber(lastSymbol)) {
			return "", ErrInvalidString
		}

		if (!unicode.IsNumber(symbol)) && !unicode.IsNumber(lastSymbol) && idx > 0 {
			result.WriteRune(lastSymbol)
		}

		if unicode.IsNumber(symbol) && lastSymbol != 0 {
			countRepeat, _ := strconv.Atoi(string(symbol))
			result.WriteString(strings.Repeat(string(lastSymbol), countRepeat))
		}

		lastSymbol = symbol
	}

	if !unicode.IsNumber(lastSymbol) && len(s) > 0 {
		result.WriteRune(lastSymbol)
	}

	return result.String(), nil
}
