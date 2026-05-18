package main
import (
	"os"
	"fmt"
	"strings"
	"flag"
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
	mode := flag.String("mode", "decode", "choose encode or decode")
	flag.Usage = func (){
		fmt.Println("Usage: ./ --mode=<encode|decode> <input string>")
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("input encode/decode")
		os.Exit(1)
	}
	InputLine := strings.Split(args[0], "\n")
	switch *mode {
	case "decode":
		for _, v := range InputLine{
			if !BracketCheck(v) {
				fmt.Println("Close the bracket")
				os.Exit(1)
			}
			Output, malformed:= decoder(v)
			if !malformed {
				os.Exit(1)
			}
			for _, v := range Output {
				if v == '[' || v == ']' {
					fmt.Println("input malformed")
					os.Exit(1)			
				}
			}
			fmt.Println(Output)
		}				

	case "encode":
		for _, v := range InputLine{
			Output := encoder(v)
			
			fmt.Println(Output)
		}
	default:
		fmt.Println("invalid mode: use encode or decode")
		os.Exit(1)
	}	
}
