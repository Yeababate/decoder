package main
import (
"fmt"
"strconv"
)

func encoder(input string) string {
	store := ""
	output := ""
	num := ""
	check := ""
	count := 0
	for i:= 0; i < len(input); i+=2 {
		j := i+1
		if j < len(input) {
			check = string(input[i]) + string(input[j])
			// fmt.Println(check)
			if j+2 > len(input) {
				// store += string(input[i])
				num = strconv.Itoa(count)
				output += string('[') + num + string(' ') + check + string(']')
				check = ""
				count = 0				
				break
			}else if check == input[j+1:j+2] {
				count++
				fmt.Println(count)
			}else {
				check = ""
				store += string(input[i])
				num = strconv.Itoa(count)
				output += string('[') + num + string(' ') + store + string(']')
				store = ""
			}
		}else {
			break
		}
	}
	return output
}

func main() {
	str := "ababab"
	fmt.Println(encoder(str))
}






// [3 ab][4 b] >> ab ab ab bbbb