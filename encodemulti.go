package main
import (
"fmt"
"strconv"
)

func encoder(input string) string {
	output := ""
	check := ""
	count := 1
	if len(input) == 0 {
		return output
	}
	if len(input) >= 2 {
		check = input[:2]
	}else {
		check = input
	}

	for i:= 2; i < len(input); i+=2 {
		nextCheck := ""
		if i+1 < len(input) {
			nextCheck = input[i:i+2]
		}else {
			nextCheck = input[i:]
		}

		if nextCheck == check{
			count++
		}else{
			output += "[" + strconv.Itoa(count) + " " + check + "]"
			check = nextCheck
			count = 1
		}
	}
	output += "[" + strconv.Itoa(count)+ " " + check + "]"
	return output
}

func main() {
	str := "abaab"
	fmt.Println(encoder(str))
}

// abababcc >> [3 ab][2 b] 