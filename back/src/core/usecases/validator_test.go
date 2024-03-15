package usecases

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: all those tests are white box tests, they are testing the implementation of the functions, we should move those tests to a separate file and do black box tests instead

func TestStringHasNumberAtStart(t *testing.T) {
	// Given
	stringToTest := "7est"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasNumberInMiddle(t *testing.T) {
	// Given
	stringToTest := "te5t"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasNumberAtEnd(t *testing.T) {
	// Given
	stringToTest := "tes7"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasNoNumber(t *testing.T) {
	// Given
	stringToTest := "test"

	// When
	result := stringHasNumber(stringToTest)

	// Then
	assert.Equal(t, result, false)
}

func TestStringHasUppercaseLetterAtStart(t *testing.T) {
	// Given
	stringToTest := "Test"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasUppercaseLetterInMiddle(t *testing.T) {
	// Given
	stringToTest := "teSt"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasUppercaseLetterAtEnd(t *testing.T) {
	// Given
	stringToTest := "tesT"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasNoUppercaseLetter(t *testing.T) {
	// Given
	stringToTest := "test"

	// When
	result := stringHasUppercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, false)
}

func TestStringHasLowercaseLetterAtStart(t *testing.T) {
	// Given
	stringToTest := "tEST"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasLowercaseLetterInMiddle(t *testing.T) {
	// Given
	stringToTest := "TEsT"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasLowercaseLetterAtEnd(t *testing.T) {
	// Given
	stringToTest := "TESt"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasNoLowercaseLetter(t *testing.T) {
	// Given
	stringToTest := "TEST"

	// When
	result := stringHasLowercaseLetter(stringToTest)

	// Then
	assert.Equal(t, result, false)
}

func TestStringHasSpecialCharacterAtStart(t *testing.T) {
	// Given
	stringToTest := "!est"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasSpecialCharacterInMiddle(t *testing.T) {
	// Given
	stringToTest := "te~t"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasSpecialCharacterAtEnd(t *testing.T) {
	// Given
	stringToTest := "tes!"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	assert.Equal(t, result, true)
}

func TestStringHasNoSpecialCharacter(t *testing.T) {
	// Given
	stringToTest := "test"

	// When
	result := stringHasSpecialCharacter(stringToTest)

	// Then
	assert.Equal(t, result, false)
}
