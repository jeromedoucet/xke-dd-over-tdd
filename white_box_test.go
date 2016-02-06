package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Should_Add_Letters_B_PLUS_C_Equal_TO_3(t *testing.T) {
	number := convertToNumber("BC")
	assert.Equal(t, uint(3), number)
}

func Test_Should_ReAdd_If_More_Than_9(t *testing.T) {
	number := convertToNumber("9B")
	assert.Equal(t, uint(10), number)
}

func Test_Should_Compute_9_For_GIDEON(t *testing.T) {
	number := convertToNumber("GIDEON")
	assert.Equal(t, uint(18), number)
}

func Test_Should_Compute_0_For_Akuku(t *testing.T) {
	number := convertToNumber("Akuko")
	assert.Equal(t, uint(4), number)
}

func Test_Should_Convert_In_Type(t *testing.T) {
	ddtype := convertToType(uint(0))
	assert.Equal(t, 'A', ddtype)
}

func Test_Should_Convert_Type_to_Medal(t *testing.T) {
	item := convertToItem('A')
	assert.Equal(t, "Medal", item)
}

func Test_Should_Reduce_to_One_Digit(t *testing.T) {
	digit := reduceToDigit(uint(66))
	assert.Equal(t, uint(3), digit)
}