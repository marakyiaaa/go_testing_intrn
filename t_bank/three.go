package main

import (
	"fmt"
	"strings"
)

func main() {
	var set, value string
	var num, tmp, result, copy_index, value_index int
	fmt.Scan(&set, &value, &num)

	copySetTwo := value
	start_point := len(set)

	for i := len(set) - 1; i >= 0; i-- {
		value_index = strings.IndexRune(value, rune(set[i]))
		if value_index != -1 {
			if tmp == 0 {
				start_point = i
			}
			copy_index = strings.IndexRune(copySetTwo, rune(set[i]))
			if copySetTwo != "" && copy_index != -1 {
				copySetTwo = copySetTwo[:copy_index] + copySetTwo[copy_index+1:]
			}
			if copySetTwo == "" && copy_index == -1 && num-tmp == 1 {
				tmp = 0
				copySetTwo = value
				i = start_point
			}
			if (copySetTwo != "" && num-tmp > 1) || copy_index != -1 {
				tmp++
			}

		} else if value_index == -1 {
			tmp = 0
			copySetTwo = value
			i = start_point
		}
		if tmp == num {
			result = i
			break
		}
	}
	fmt.Println(set[result : result+num])

}
