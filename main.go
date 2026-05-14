package main
import (
	"os"
	"fmt"
	"strings"
)

var countleft, countright int
func BracketCheck(str string) bool {
		for _, v := range str {
			if v == ']' {
				countleft++
			}else if v == '[' {
				countright++
			}
		}
		if countleft == countright {
			return true
		} else {
			return false
		}
	}

func main() {
if len(os.Args) != 2 {
	fmt.Println("Usage: go run main.go encoded stuff")
	os.Exit(1)
}

InputFile := os.Args[1]

InputLine := strings.Split(InputFile, "\n")
for _, v := range InputLine{
	if !BracketCheck(v) {
		fmt.Println("Close the bracket")
		os.Exit(1)
	}
	Output, malformed:= decoder(v)
	for _, v := range Output {
		if v == '[' || v == ']' {
			fmt.Println("input malformed test")
			os.Exit(1)			
	}
	}
	if !malformed {
		fmt.Println("input malformed test")
		os.Exit(1)
	} 
	fmt.Println(Output)
}
}
