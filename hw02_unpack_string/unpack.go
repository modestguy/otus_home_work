package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(value string) (string, error) {
	var result strings.Builder
	var prevCharBuilder strings.Builder
	var prevNumberCharBuilder strings.Builder
	var countsBuilder strings.Builder
	var lastChar rune

	if len(value) == 0 {
		return "", nil
	}

	for pos, currentRune := range value {
		if pos == 0 {
			if unicode.IsDigit(currentRune) {
				return "", ErrInvalidString
			}
			prevCharBuilder.WriteRune(currentRune)
			continue
		}

		if unicode.IsDigit(currentRune) {
			_, err := strconv.Atoi(prevNumberCharBuilder.String())
			if err == nil {
				return "", ErrInvalidString
			}
			countsBuilder.WriteRune(currentRune)
			count, _ := strconv.Atoi(countsBuilder.String())
			if count != 0 {
				result.WriteString(strings.Repeat(prevCharBuilder.String(), count))
			}
			prevCharBuilder.Reset()
			prevNumberCharBuilder.WriteRune(currentRune)
		} else {
			result.WriteString(prevCharBuilder.String())
			prevCharBuilder.Reset()
			prevCharBuilder.WriteRune(currentRune)
			prevNumberCharBuilder.Reset()
		}
		countsBuilder.Reset()
		lastChar = currentRune
	}

	if !unicode.IsDigit(lastChar) {
		prevCharBuilder.Reset()
		prevCharBuilder.WriteRune(lastChar)
		result.WriteString(prevCharBuilder.String())
	}
	return result.String(), nil
}
