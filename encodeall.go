package main

import (
	"fmt"
	"strconv"
)

func encoder(input string) string {
	output := ""
	i := 0
	for i < len(input) {
		// Decide if we can take 2 characters, otherwise just 1
		unit := ""
		if i+1 < len(input) {
			unit = input[i:i+2]
			// fmt.Println(input[i:i+2])
		} else {
			unit = string(input[i])
		}

		count := 1
		// Count how many times this unit repeats
		for {
			nextStart := i + count*len(unit)
			nextEnd := nextStart + len(unit)
			if nextEnd > len(input) {
				break
			}
			if input[nextStart:nextEnd] == unit {
				count++
			} else {
				break
			}
		}
		output += "[" + strconv.Itoa(count) + " " + unit + "]"
		i += count * len(unit)
	}

	return output
}

func main() {
	fmt.Println(encoder("aaaaaaa"))
}