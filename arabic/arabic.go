/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package arabic

import (
	"errors"
	"strings"
)

var letters = map[rune]bool{
	Hamza:              true,
	AlefWithMadda:      true,
	AlefWithHamzaAbove: true,
	WawWithHamza:       true,
	AlefWithHamzaBelow: true,
	YehWithHamzaAbove:  true,
	Alef:               true,
	Beh:                true,
	TehMarbuta:         true,
	Teh:                true,
	Theh:               true,
	Jeem:               true,
	Hah:                true,
	Khah:               true,
	Dal:                true,
	Thal:               true,
	Reh:                true,
	Zain:               true,
	Seen:               true,
	Sheen:              true,
	Sad:                true,
	Dad:                true,
	Tah:                true,
	Zah:                true,
	Ain:                true,
	Ghain:              true,
	Feh:                true,
	Qaf:                true,
	Kaf:                true,
	Lam:                true,
	Meem:               true,
	Noon:               true,
	Heh:                true,
	Waw:                true,
	AlefMaksura:        true,
	Yeh:                true,
	AlefWaslah:         true,
}

var vowels = map[rune]bool{
	Sukoon:   true,
	Damma:    true,
	Fatha:    true,
	Kasra:    true,
	Dammatan: true,
	Fathatan: true,
	Kasratan: true,
}

// TODO(Amr Ojjeh): Consider making Vowel []rune so that more than one correct answer can be supported
type LetterPack struct {
	Letter          rune
	Vowel           rune
	Shadda          bool
	SuperscriptAlef bool
}

func (l LetterPack) String() string {
	builder := strings.Builder{}
	if IsVowel(l.Letter) {
		panic("letter cannot be a vowel")
	}
	if IsLetter(l.Vowel) {
		panic("vowel cannot be a letter")
	}
	if l.Letter != 0 {
		builder.WriteRune(l.Letter)
	}
	if l.Vowel != 0 {
		builder.WriteRune(l.Vowel)
	}
	if l.Shadda {
		builder.WriteRune(Shadda)
	}
	if l.SuperscriptAlef {
		builder.WriteRune(SuperscriptAlef)
	}
	return builder.String()
}

func ParseLetterPack(str string) (LetterPack, error) {
	lp := LetterPack{}
	for _, l := range str {
		switch l {
		case Shadda:
			lp.Shadda = true
		case SuperscriptAlef:
			lp.SuperscriptAlef = true
		default:
			if IsVowel(l) {
				if lp.Vowel != 0 {
					return LetterPack{}, errors.New("arabic: cannot have more than one vowel")
				}
				lp.Vowel = l
			} else if IsLetter(l) {
				if lp.Letter != 0 {
					return LetterPack{}, errors.New("arabic: cannot have more than one letter")
				}
				lp.Letter = l
			} else {
				return LetterPack{}, errors.New("arabic: unexpected character")
			}
		}
	}
	return lp, nil
}

// UnpointedString returns the word without any vowels
func Unpointed(pointed string) string {
	res := ""
	for _, l := range pointed {
		c := string(l)
		if !IsVowel(l) {
			res += c
		}
	}
	return res
}

// IsVowel checks if the character is a fatha, kasra, damma, or sukoon, with
// their tanween variations
func IsVowel(letter rune) bool {
	return vowels[letter]
}

func IsLetter(letter rune) bool {
	return letters[letter]
}
