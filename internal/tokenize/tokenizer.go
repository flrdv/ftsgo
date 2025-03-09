package tokenize

import (
	"iter"
	"unicode"
)

func String(str string) iter.Seq[string] {
	return func(yield func(string) bool) {
		lastSplit := 0

		for i, char := range str {
			if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
				if i-lastSplit != 0 && !yield(str[lastSplit:i]) {
					return
				}

				lastSplit = i + 1
			}
		}

		if len(str)-lastSplit != 0 {
			yield(str[lastSplit:])
		}
	}
}
