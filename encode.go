package main
import (
"fmt"
"strconv"
)

func encoder(input string) string {
	store := ""
	output := ""
	num := ""
	for i:= 0; i < len(input); i++ {
		if i+1 < len(input) && input[i] == input[i+1] {
			store += string(input[i])
		}else {
			store += string(input[i])
			num = strconv.Itoa(len(store))
			output += string('[') + num + string(' ') + string(input[i]) + string(']')
			store = ""
		}
	}
	return output
}

func main() {
	str := "aabbbb"
	fmt.Println(encoder(str))
}






// [3 ab][4 b] >> ababab bbbb