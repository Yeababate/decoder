package main
import (
"fmt"
"strconv"
"strings"
)

func decoder(input string) (string, bool) {
	var store, output, rep string
	var i,j,k int
	
	for i = 0; i < len(input); i++{
		if input[i] == '[' {
			// countright++
			if !((input[i+1] >= '0') && (input[i+1] <= '9')) {
				fmt.Println("test")
				return output, false
			}
			for j = i+1; j < len(input); j++ {
				if (input[j] >= '0' && input[j] <= '9') {
					if !Check(input[j+1]){
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
		} else {
			
			output += string(input[i])
		}
	}

	return output, true
}
func Check(ToBeChecked byte) bool {
	if !((ToBeChecked >= '0') && (ToBeChecked <= '9')) && (ToBeChecked != ' ') { 
		return false
	}else {
		return true
	}
}