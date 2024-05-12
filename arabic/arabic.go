package arabic

var vowels = map[rune]bool{
	Sukoon:   true,
	Damma:    true,
	Fatha:    true,
	Kasra:    true,
	Dammatan: true,
	Fathatan: true,
	Kasratan: true,
}

// UnpointedString returns the word without any vowels
func Unpointed(pointed string) string {
	res := ""
	for _, l := range pointed {
		c := string(l)
		if !IsPointedChar(l) {
			res += c
		}
	}
	return res
}

// IsPointedChar checks if the character is a fatha, kasra, damma, or sukoon, with
// their tanween variations
func IsPointedChar(letter rune) bool {
	return vowels[letter]
}
