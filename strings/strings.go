package strings

import "unicode"

func OnlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}
