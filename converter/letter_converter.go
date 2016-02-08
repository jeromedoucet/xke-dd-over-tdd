package converter

import (
	"errors"
	"unicode"
)

func Convert(letter rune) (result int, err error) {
	value, ok := reference[unicode.ToUpper(letter)]
	if ok {
		return value, err
	} else {
		err = errors.New("Invalid value")
		return -1, err
	}
}