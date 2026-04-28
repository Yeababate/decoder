package main
import (
"fmt"
"strconv"
"os"
)


func WriteRep(length int, rep string) string{
	Changed := ""
	for i := 0; i < length; i++ {
		for j := 0; j < len(rep); j++ {
			Changed += string(rep[j])
		}
	}
	return Changed
}

func Check(ToBeChecked byte) bool {
	if !((ToBeChecked >= '0') && (ToBeChecked <= '9')) && (ToBeChecked != ' ') {
		return true
	}else {
		return false
	}
}
func BracketCheck(input string) bool{
	count := 0
	for _,v:= range(input){
		if v == '[' || v == ']' {
			count++
		}
	}
	if !(count % 2 == 0) {
		fmt.Println("input malformed\n")
		os.Exit(0)
	}
	return false
}

func decoder(input string) string {
	store := ""
	output := ""
	rep := ""
	var i,j,k int
	for i = 0; i < len(input); i++{
		if input[i] == '[' {
			if !((input[i+1] >= '0') && (input[i+1] <= '9')) {
					fmt.Println("input in the right format\n")
					os.Exit(0)
				}
			for j = i+1; j < len(input); j++ {
				if (input[j] >= '0' && input[j] <= '9') {
					if Check((input[j+1])) == true {
						fmt.Println("Arguments should be separated by comma\n")
						os.Exit(0)
					}
					// if !((input[j+1] >= '0') && (input[j+1] <= '9')) && (input[j+1] != ' ') {
					// 	// fmt.Println(string(input[j]))
					// 	// fmt.Println(string(input[j+1]))
					// 	fmt.Println("Arguments should be separated by comma")
					// 	os.Exit(0)
					// }
					store += string(input[j])
				} else if input[j] == ' ' {
					length, err := strconv.Atoi(store)
					if err != nil {
						fmt.Printf("Error changing to number\n")
					}

					for k = j+1; k < len(input); k++{
						if input[k] != ']' {
							// EmptyCheck(input[k-1])
							rep += string(input[k])
							store = ""
						} else {
							i = k
							break
						}
					}
					output += WriteRep(length,rep )
					length = 0
					rep = ""
					break
				}
			}
		}else {
			output += string(input[i])
		}
	}
	return output
}
func main() {
	input := "[5 a][5 -_]-[5 #]"
	BracketCheck(input)
	fmt.Println(decoder(input))
}