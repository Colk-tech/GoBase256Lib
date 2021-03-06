package utils

import (
	"unicode/utf8"
)

func StrToRune(text string) (result rune, err error) {
	lengthOfText := utf8.RuneCountInString(text)

	if lengthOfText == 0 {
		return result, err
	}

	if lengthOfText != 1 {
		return result, err
	}
	result = []rune(text)[0]

	return result, err
}

func RuneToStr(r rune) (result string) {
	runes := []rune{r}

	return string(runes)
}
