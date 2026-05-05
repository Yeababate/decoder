package main
import (
"fmt"
"strconv"
"os"
"bufio"
"strings"
)

// func WriteRep(length int, rep string) string{
// 	Changed := ""
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < len(rep); j++ {
// 			Changed += string(rep[j])
// 		}
// 	}
// 	return Changed
// }

func Check(ToBeChecked byte) bool {
	if !((ToBeChecked >= '0') && (ToBeChecked <= '9')) && (ToBeChecked != ' ') {
		fmt.Println("input malformed\n")
		os.Exit(1)
		return true
	}else {
		return false
	}
}

func BracketCountCheck() bool{     // edit count for both brackets eg, [2 a[[3 b] shouldn't work
	var countleft, countright int
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)               // Fix to return to main only no Exit here
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		for _,v:= range(line){
			if v == '[' {
				countright++
			}else if v == ']' {
				countleft++
			}
		}
		if countleft != countright {
			// fmt.Println("input malformed\n")
			return true
		}
		if errorCheck(line) {
			return true
		}
	}
	return false
}
// func BracketCheck(input string) bool{     // edit count for both brackets eg, [2 a[[3 b] shouldn't work
// 	var countleft, countright int
// 	for _,v:= range(input){
// 		if v == '[' {
// 			countright++
// 		}else if v == ']' {
// 			countleft++
// 		}
// 	}
// 	if countleft != countright {
// 		fmt.Println("input malformed\n")
// 		os.Exit(1)
// 	}
// 	return false
// }
func errorCheck (input string) bool {
	var i,j,k int
	for i = 0; i < len(input); i++{
		if input[i] == '[' {
			if !((input[i+1] >= '0') && (input[i+1] <= '9')) {
					// fmt.Println("input in the right format\n")
					return true
				}
			for j = i+1; j < len(input); j++ {
				if (input[j] >= '0' && input[j] <= '9') {
					Check(input[j+1])

				} else if input[j] == ' ' {
					for k = j+1; k < len(input); k++{
						if input[k] != ']' {
							continue
						} else {
							i = k
							break
						}
					}
					break
				}
			}
		}
	}
	return false
}
func decoder(input string) string {
	var store, output, rep string
	var i,j,k int
	for i = 0; i < len(input); i++{
		if input[i] == '[' {
			if !((input[i+1] >= '0') && (input[i+1] <= '9')) {
					fmt.Println("input in the right format\n")
					os.Exit(1)
				}
			for j = i+1; j < len(input); j++ {
				if (input[j] >= '0' && input[j] <= '9') {
					Check(input[j+1])

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
	return output
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}
	// var inputLines [] string
	// var art [] string
	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if BracketCountCheck() == true {
		fmt.Println("input malformed\n")
		os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	// for scanner.Scan(){
	// 	line := scanner.Text()
	// 	inputLines = append(inputLines, line)
	// }
	// for _, line := range inputLines {
	// 		BracketCheck(inputFile)
	// 		art = append(art, (decoder(line) + "\n"))
	// 	}
	// fmt.Println(art)
	for scanner.Scan(){
		line := scanner.Text()
		// BracketCheck(line)
		output := decoder(line)
		fmt.Println(output)

	}
}