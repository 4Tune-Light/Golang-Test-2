package test0

import (
	"strings"
)

func Test0(word string) map[string]int{
	var consonant bool
	var alive int
	var dead int

	data := make(map[string]int)
	vowels := []int32{'a', 'i', 'u', 'e', 'o'}
	check := []bool{false, false, false, false, false}

	for _, char := range strings.ToLower(word) {
		consonant = true

		for i, vowel := range vowels {
			if char == vowel {
				consonant = false
				if check[i] == false {
					alive++
					check[i] = true
				}
			}
		}

		if consonant == true {
			dead++
		}
	}

	data["vowels"] = alive
	data["consonants"] = dead

	return data
}