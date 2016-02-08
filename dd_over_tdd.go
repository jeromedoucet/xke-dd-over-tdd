package main

import (
	"strconv"
	"xke-dd-over-tdd/converter"
)

func LootSomeStuff(playerName string) string {
	number := convertToNumber(playerName)
	digit := reduceToDigit(number)
	itemType := convertToType(digit)
	return convertToItem(itemType)
}

func convertToNumber(name string) uint {
	var sum uint = 0
	for _, runeVal := range name {
		value, err := converter.Convert(runeVal)
		if err == nil {
			sum += uint(value)
		}
	}
	return sum
}

func reduceToDigit(value uint) (aggregate uint) {
	for _, runeVal := range strconv.Itoa(int(value)) {
		aggregate += uint(runeVal - '0')
	}
	if aggregate > 9 {
		aggregate = reduceToDigit(aggregate)
	}
	return aggregate
}

func convertToType(intType uint) rune {
	return converter.ItemsReferences[int(intType)]
}

func convertToItem(itemType rune) string {
	return converter.Items[itemType]
}