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
			lAkwas := false
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

					keyword  = ""
				} else if keyword == "low" {
					for i := index-1 ; i > 0 ; i--{
						if res[i] == ' ' {
							break
						}
						
					}
				}
			}
			fmt.Println(string(keyword))
		}
	}
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
