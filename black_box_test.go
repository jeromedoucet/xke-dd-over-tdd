package main_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"xke-dd-over-tdd"
)

func Test_Should_Loot_Fire_Ring_for_Akuko(t *testing.T) {
	item := main.LootSomeStuff("Akuko")
	assert.Equal(t, "Ring of Fire Resistance", item)
}

func Test_Should_Loot_Speed_boot_for_Kayin(t *testing.T) {
	item := main.LootSomeStuff("Kayin")
	assert.Equal(t, "Boots of Speed", item)
}
