package ast

import "strings"

// clean up padding/additional data characters that can appear in character
// literals such as '_'
func purgeNumericStrings(s string) string {
	return strings.Replace(s, "_", "", -1)
}
