package utils

import "strconv"

func IsValidZipcode(zipcode string) bool {
	if len(zipcode) != 8 {
		return false
	}
	for _, char := range zipcode {
		if _, err := strconv.Atoi(string(char)); err != nil {
			return false
		}
	}
	return true
}
