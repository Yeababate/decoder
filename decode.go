package main
import (
"fmt"
"strconv"
"strings"
"unicode"
)

func decoder(input string) (string, bool) {
	var store, output, rep string
	var i,j,k int
	
	for i = 0; i < len(input); i++{
		if input[i] == '[' {
			if !unicode.IsDigit(rune(input[i+1])) {
				return output, false
			}
			for j = i+1; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					if !unicode.IsDigit(rune(input[j+1])) && input[j+1] != ' ' {
						fmt.Println("should be space between the arguments")    
						return output, false
					}    
				store += string(input[j])
				} else if input[j] == ' ' {
					length, err := strconv.Atoi(store)
					if err != nil {
						fmt.Printf("Error changing to number\n")
					}
					for k = j+1; k < len(input); k++{
						if input[k] != ']' {
							rep += string(input[k])
							store = ""
						} else {
							i = k
							break
						}
					}
					output += strings.Repeat(rep, length )
					length = 0
					rep = ""
					break
				}
			}

		}else if input[i] == ']' {
			return output, false
		}else {
			
			output += string(input[i])
		}
	}
	return output, true
}
