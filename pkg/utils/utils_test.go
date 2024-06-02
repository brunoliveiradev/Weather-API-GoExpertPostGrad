package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidZipcode_ValidZipcode(t *testing.T) {
	assert.True(t, IsValidZipcode("12345678"), "IsValidZipcode failed for valid zipcode")
}

func TestIsValidZipcode_InvalidLength(t *testing.T) {
	assert.False(t, IsValidZipcode("1234567"), "IsValidZipcode did not fail for zipcode of invalid length")
}

func TestIsValidZipcode_InvalidCharacters(t *testing.T) {
	assert.False(t, IsValidZipcode("1234567a"), "IsValidZipcode did not fail for zipcode with invalid characters")
}

func TestIsValidZipcode_EmptyString(t *testing.T) {
	assert.False(t, IsValidZipcode(""), "IsValidZipcode did not fail for empty string")
}
