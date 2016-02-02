package main

import (
	"strconv"
	"xke-dd-over-tdd/converter"
)

func reduce(value int) (computed int, err error) {
	for _, runeVal := range strconv.Itoa(value) {
		computed += int(runeVal - '0')
	}
	if (computed > 9) {
		computed, err = reduce(computed)
	}
	return computed, err
}

func convertToNumber(name string) (int, error) {
	var sum int = 0
	for _, runeVal := range name {
		value, err := converter.Convert(runeVal)
		// todo test case
		if err != nil {
			return sum, err
		}
		sum += value
	}
	return reduce(sum)
}

func convertToType(name string) (rune, error) {
	num, err := convertToNumber(name)
	// todo handle errors !
	return converter.ItemsReferences[num], err
}

func convertToItem(name string) (string, error) {
	dndType, err := convertToType(name)
	return converter.Items[dndType], err
}