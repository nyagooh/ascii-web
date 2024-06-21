package asciiart

import (
	"unicode"
)

func IsPrintable(str string) bool {
	for i, value := range str {
		if i < len(str)-1 && str[i] == '\\' {
			if str[i+1] == 't' || str[i+1] == 'a' || str[i+1] == 'b' || str[i+1] == 'v' || str[i+1] == 'f' || str[i+1] == 'r' {
				return false
			}
		}
		if !unicode.IsPrint(value) {
			return false
		}
	}
	return true
}
