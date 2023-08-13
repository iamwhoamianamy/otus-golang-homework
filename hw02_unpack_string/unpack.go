package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

type ParsedSymbolType int

const (
	Symbol ParsedSymbolType = iota
	Digit
)

func parseSymbol(symbol byte) (ParsedSymbolType, int) {
	if val, err := strconv.Atoi(string(symbol)); err == nil {
		return Digit, val
	}

	return Symbol, 0
}

func Unpack(encoded string) (string, error) {
	if len(encoded) == 0 {
		return "", nil
	}

	builder := strings.Builder{}

	hasLastSymbol := false
	var lastSymbol byte

	for i := 0; i < len(encoded); i++ {
		symbol := encoded[i]

		symbolType, digit := parseSymbol(symbol)

		switch symbolType {
		case Symbol:
			if hasLastSymbol {
				builder.WriteString(string(lastSymbol))
			}

			lastSymbol = symbol
			hasLastSymbol = true
		case Digit:
			if hasLastSymbol {
				repeatedString := strings.Repeat(string(lastSymbol), digit)
				builder.WriteString(repeatedString)

				hasLastSymbol = false
			} else {
				return "", ErrInvalidString
			}
		}
	}

	if hasLastSymbol {
		builder.WriteString(string(lastSymbol))
	}

	return builder.String(), nil
}
