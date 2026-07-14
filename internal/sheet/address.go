package sheet

import (
	"strconv"
	"strings"

	"github.com/uplang/tsvsheet.go/internal/constants"
)

// AddressText is spreadsheet-address source text (`A1`, `F4`) accepted by
// ParseAddress. It is exported so callers in other packages can convert their
// string input at the call site.
type AddressText string

// Address is a cell coordinate in spreadsheet notation (`F4`): column letters
// plus a 1-based row. It carries 0-based indices internally.
type Address struct {
	Row int `json:"row"` // 0-based
	Col int `json:"col"` // 0-based
}

// ParseAddress parses spreadsheet notation (`A1`, `F4`, `AA10`) into an
// Address. The column is one or more ASCII uppercase letters, the row a
// positive integer; anything else is constants.ErrInvalidValue.
func ParseAddress(s AddressText) (Address, error) {
	letters, digits := splitLetters(s)
	if letters == "" || digits == "" {
		return Address{}, constants.ErrInvalidValue.With(nil, "address", string(s))
	}
	row, err := strconv.Atoi(digits)
	if err != nil || row < 1 {
		return Address{}, constants.ErrInvalidValue.With(nil, "address", string(s))
	}
	return Address{Row: row - 1, Col: lettersToIndex(columnLetters(letters))}, nil
}

// splitLetters splits a spreadsheet address into its leading uppercase-letter
// run and trailing digit run; a malformed shape (a non-digit in the trailing
// run) leaves the digit part empty.
func splitLetters(s AddressText) (letters, digits string) {
	i := 0
	for i < len(s) && s[i] >= 'A' && s[i] <= 'Z' {
		i++
	}
	rest := string(s[i:])
	for j := 0; j < len(rest); j++ {
		if rest[j] < '0' || rest[j] > '9' {
			return string(s[:i]), ""
		}
	}
	return string(s[:i]), rest
}

// String renders the Address in spreadsheet notation.
func (a Address) String() string {
	return indexToLetters(colIndex(a.Col)) + strconv.Itoa(a.Row+1)
}

// lettersToIndex converts spreadsheet column letters to a 0-based index
// (A→0, Z→25, AA→26), bijective base-26.
func lettersToIndex(letters columnLetters) int {
	index := 0
	for i := 0; i < len(letters); i++ {
		index = index*26 + int(letters[i]-'A') + 1
	}
	return index - 1
}

// indexToLetters converts a 0-based column index to spreadsheet letters.
func indexToLetters(index colIndex) string {
	var b strings.Builder
	for n := index + 1; n > 0; n = (n - 1) / 26 {
		_ = b.WriteByte(byte('A' + (n-1)%26))
	}
	return reverse(columnLetters(b.String()))
}

// reverse returns s with its bytes reversed (ASCII column letters only).
func reverse(s columnLetters) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
