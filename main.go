package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	length := len(args)
	if length < 2 {
		fmt.Println("less arguments you need to enter the input and the output files")
		os.Exit(1)
	} else if length > 2 {
		fmt.Println("too much arguments")
		os.Exit(1)
	} else {
		in := args[0]
		out := args[1]
		inExt := ""
		index := -1
		outExt := ""
		for i := len(in) - 1; i > 0; i-- {
			if in[i] == '.' {
				index = i
			}
		}
		inExt = in[index+1:]
		for j := len(out) - 1; j > 0; j-- {
			if out[j] == '.' {
				index = j
			}
		}
		outExt = out[index+1:]
		if inExt != "txt" || outExt != "txt" {
			fmt.Println("the extension must be .txt")
			os.Exit(1)
		} else {
			file, err := os.Open(in)
			if err != nil {
				fmt.Println("Err msg: ", err)
				return
			}
			defer file.Close()
			res, err := io.ReadAll(file)
			if err != nil {
				fmt.Println("err msg :", err)
				return
			}
			t := ""
			insideParenthese := false
			for _, v := range res {
					if v == '(' {
						t += " " + string(v)
						insideParenthese = true
					} else if v == ')' {

						t += string(v) + " "
						insideParenthese = false

					} else {
						if insideParenthese {
							if v ==','{
								t += string(v)
							} else if v != ' '{
								t += string(v)
							} 
						} else {
							if v <32 && v <48 {
								t += " " + string(v) + " "
							} else {
								t += string(v)
							}
						}
					}
				
			}
			arr1 := strings.Fields(string(t))
			res2 := ""
			for _, item := range arr1{
				if strings.HasPrefix(item, "(") && strings.HasSuffix(item, ")") {
					content := item[1:len(item)-1]
					if strings.Contains(content, ",") {
						res2 += "(" + content + ") "
					} else {
						// Apply rules to update the content
						switch content {
						case "cap":
							res2 += "(cap,1) "
						case "low":
							res2 += "(low,1)"
						case "up":
							res2 += "(up,1)"
						case "hex", "bin":
							res2 += "(" + content + ") "
						default:
							res2 += "(" + content + ") "
						}
					}
				} else {
					res2 += item + " "
				}		
				
			}
			arr := strings.Fields(res2)
			for i := 0; i < len(arr); i++ {
				if arr[i] == "(cap)" || arr[i] == "low" || arr[i] == "up" {
					// if i+2 < len(arr) && arr[i+1] == "," {
					// 	num, err := strconv.Atoi(arr[i+2])
					// 	if err != nil {
					// 		fmt.Println("msg err : not a number ", err)
					// 		return
					// 	}
					// }
					if arr[i] == "cap" {
						arr[i-2] = Capitalize(arr[i-2])
						arr[i] = ""
						arr[i+1] = ""
						arr[i-1] = ""
					} else if arr[i] == "low" {
						arr[i-2] = ToLower(arr[i-2])
						arr[i] = ""
						arr[i+1] = ""
						arr[i-1] = ""

					} else if arr[i] == "up" {
						arr[i-2] = ToUpper(arr[i-2])
						arr[i] = ""
						arr[i+1] = ""
						arr[i-1] = ""

					}
				} else if arr[i] == "bin" {
					integer, err := strconv.ParseInt(arr[i-2], 2, 64)
					if err != nil {
						fmt.Println("you can't convert")
						return
					}
					arr[i-2] = strconv.FormatInt(integer, 10)
					arr[i] = ""
					arr[i+1] = ""
					arr[i-1] = ""
				} else if arr[i] == "hex" {
					integer, err := strconv.ParseInt(arr[i-2], 16, 64)
					if err != nil {
						fmt.Println("you can't convert")
						return
					}
					arr[i-2] = strconv.FormatInt(integer, 10)
					arr[i] = ""
					arr[i+1] = ""
					arr[i-1] = ""
				}
			}
			/*lAkwas := false
			keyword := ""
			for i := 0; i < len(res); i++ {
				index := 0
				if res[i] == '(' {
					lAkwas = true
					index = res[i]
					continue
				}
				if res[i] == ')' {
					lAkwas = false
					continue
				}
				if lAkwas {
					keyword += string(res[i])
				}
				if keyword == "cap" {
					for i := index-1 ; i > 0 ; i--{
						if res[i] == ' ' {
							break
						}

					}category:formatters go
					keyword  = ""
				} else if keyword == "low" {
					for i := index-1 ; i > 0 ; i--{
						if res[i] == ' ' {
							break
						}

					}
					keyword = ""
				}
			}*/
			str := ""
			for _, v := range arr {
				str += v + " "
			}
			arr2 := strings.Fields(string(str))
			fmt.Printf("%v\n", arr2)
		}
	}
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
