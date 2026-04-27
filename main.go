package main
import (
"fmt"
"strconv"
// "os"
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

func decoder(input string) string {
	store := ""
	output := ""
	rep := ""
	var i,j,k int
	for i = 0; i < len(input); i++{
		if input[i] == '[' {
			i++
			for j = i; j < len(input); j++ {
				// fmt.Printf("%d", i)
				if (input[j] >= '0' && input[j] <= '9') {
					store += string(input[j])
					// fmt.Println(store)
				} else if input[j] == ' ' {
					j++
					length, err := strconv.Atoi(store)
					// fmt.Printf("%d", length)
					if err != nil {
						fmt.Printf("Error changing to number")
					}

					for k = j; k < len(input); k++{
						if input[k] != ']' {
							rep += string(input[k])
						} else {
							i = k
							break
						}
					}
					output += WriteRep(length,rep )
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
	fmt.Println(decoder("[5 d][3 f]"))
}