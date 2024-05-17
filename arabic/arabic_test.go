/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package arabic

import (
	"reflect"
	"testing"
)

func TestLetterPackString(t *testing.T) {
	tests := []struct {
		name     string
		lp       LetterPack
		expected string
	}{
		{
			name: "Basic",
			lp: LetterPack{
				Letter:          Lam,
				Vowel:           Fatha,
				Shadda:          true,
				SuperscriptAlef: false,
			},
			expected: FromBuckwalter("la~"),
		},
		{
			name: "Missing letter",
			lp: LetterPack{
				Letter:          0,
				Vowel:           Fatha,
				Shadda:          true,
				SuperscriptAlef: false,
			},
			expected: FromBuckwalter("a~"),
		},
		{
			name: "Missing vowel",
			lp: LetterPack{
				Letter:          Lam,
				Vowel:           0,
				Shadda:          true,
				SuperscriptAlef: false,
			},
			expected: FromBuckwalter("l~"),
		},
		{
			name: "Superscript",
			lp: LetterPack{
				Letter:          Heh,
				Vowel:           Fatha,
				Shadda:          false,
				SuperscriptAlef: true,
			},
			expected: FromBuckwalter("ha`"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.lp.String()
			if actual != tt.expected {
				t.Errorf("expected %s; actual: %s", tt.expected, actual)
			}
		})
	}

	t.Run("Vowel as letter", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		_ = LetterPack{
			Letter:          Fatha,
			Vowel:           0,
			Shadda:          false,
			SuperscriptAlef: false,
		}.String()
	})

	t.Run("Letter as vowel", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		_ = LetterPack{
			Letter:          0,
			Vowel:           Lam,
			Shadda:          false,
			SuperscriptAlef: false,
		}.String()
	})
}

func TestLetterPackFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected LetterPack
		err      bool
	}{
		{
			name:  "Basic",
			input: FromBuckwalter("la~"),
			expected: LetterPack{
				Letter:          Lam,
				Vowel:           Fatha,
				Shadda:          true,
				SuperscriptAlef: false,
			},
		},
		{
			name:  "Missing letter",
			input: FromBuckwalter("a~"),
			expected: LetterPack{
				Letter:          0,
				Vowel:           Fatha,
				Shadda:          true,
				SuperscriptAlef: false,
			},
		},
		{
			name:  "Missing vowel",
			input: FromBuckwalter("l~"),
			expected: LetterPack{
				Letter:          Lam,
				Vowel:           0,
				Shadda:          true,
				SuperscriptAlef: false,
			},
		},
		{
			name:  "Superscript",
			input: FromBuckwalter("ha`"),
			expected: LetterPack{
				Letter:          Heh,
				Vowel:           Fatha,
				Shadda:          false,
				SuperscriptAlef: true,
			},
		},
		{
			name:  "Double letters",
			input: FromBuckwalter("ha`l"),
			err:   true,
		},
		{
			name:  "Double vowel",
			input: FromBuckwalter("ha`a"),
			err:   true,
		},
		{
			name:  "Double shadda",
			input: FromBuckwalter("ha~~"),
			expected: LetterPack{
				Letter:          Heh,
				Vowel:           Fatha,
				Shadda:          true,
				SuperscriptAlef: false,
			},
			err: false,
		},
		{
			name:  "Double superscript",
			input: FromBuckwalter("ha``"),
			expected: LetterPack{
				Letter:          Heh,
				Vowel:           Fatha,
				Shadda:          false,
				SuperscriptAlef: true,
			},
			err: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ParseLetterPack(tt.input)
			if tt.err {
				if err == nil {
					t.Errorf("expected error")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected err: %s", err)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected: %#v; actual: %#v", tt.expected, actual)
			}
		})
	}
}
