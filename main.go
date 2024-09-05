package main

import (
	"fmt"
	"io"
	"os"
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
			t:=""
			for _, v := range res {
				if string(v) == "," {
					t+=" "+string(v)+" "
				} else {
					t+=string(v)
				}
			}
			arr := Split(string(t)," ")
			for i := 0; i < len(arr); i++ {
				if arr[i] == "(cap)"{
					arr[i-1]=ToUpper(arr[i-1])
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

					}
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
			fmt.Printf("%v\n",arr)
		}
	}
}



func Split(s, sep string) []string {
	var result []string

	if len(sep) == 0 {
		return []string{s} 
	}

	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			// Move the start position past the separator
			start = i + len(sep)
			// Skip over the separator
			i += len(sep) - 1
		}
	}

	// Add the last segment after the last separator
	result = append(result, s[start:])

	return result
}


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
