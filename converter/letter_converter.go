package converter
import (
	"errors"
	"unicode"
)

func Convert(letter rune) (result int, err error) {
	var value, ok = reference[letter]
	if (ok) {
		return value, err
	}
	value, ok = reference[unicode.ToUpper(letter)]
	if (ok) {
		return value, err
	}
	err = errors.New("Invalid value")
	return -1, err
}