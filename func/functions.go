package functions

import (
	"strings"
	"unicode"
)

func TextFormated(s []string) string {
	var res string
	ponc := ".?:!;,"
	for i, ch := range s {
		if ch == "," || ch == "!" || ch == "?" || ch == ":" || ch == ";" || ch == "." {
			if len(res) > 0 && res[len(res)-1] == ' ' {
				res = res[:len(res)-1]
			}
			if len(res) > 0 && !strings.ContainsAny(res, ponc) {
				res += ch
			} else {
				res += string(ch)
			}
			if i+1 < len(s) && s[i+1] != " " {
				res += " "
			}
		} else {
			res += ch + " "
		}
	}
	return strings.TrimSpace(res)
}

func Quote(s []string) (bool, string) {
	var result string
	quoteOpen := false 

	for i := 0; i < len(s); i++ {
		word := s[i]
		if strings.HasPrefix(word, "'") {
			result += " " + word
			quoteOpen = true
		} else if quoteOpen{
			if strings.HasPrefix(word,"'"){
				result += word + " "
				quoteOpen = false
			}
		} else {
				result += " " +word + " "
			
		}
	}

	return true, result
}



func IsWord(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) { 
			return false
		}
	}
	return true
}
func Capitalize(word string) string {
	word = ToLower(word)
	for i := 0; i < len(word); i++ {
		word = ToUpper(string(word[0])) + word[1:]
	}
	return word
}

// func Split(s, sep string) []string {
// 	var result []string

// 	if len(sep) == 0 {
// 		return []string{s}
// 	}

// 	start := 0
// 	for i := 0; i < len(s); i++ {
// 		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
// 			result = append(result, s[start:i])
// 			// Move the start position past the separator
// 			start = i + len(sep)
// 			// Skip over the separator
// 			i += len(sep) - 1
// 		}
// 	}

// Add the last segment after the last separator
// 	result = append(result, s[start:])

// 	return result
// }

func ToUpper(s string) string {
	var res []rune
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			res = append(res, i-32)
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}

func ToLower(s string) string {
	var res []rune
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			res = append(res, i+32)
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}
