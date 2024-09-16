package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	value := make([]int, n)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&value[i])
	}
	for i := 0; i < n; i++ {
		if value[i] == -1 {
			if i > 0 {
				value[i] = value[i-1] + 1
			} else {
				value[i] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			result[i] = value[i] - value[i-1]
		} else {
			result[i] = value[i]
		}

	}

	toString := func(mySlice []int) string {
		if len(mySlice) == 0 {
			return ""
		}
		var buff string
		buff += strconv.Itoa(mySlice[0])

		for i := 1; i < len(mySlice); i++ {
			buff += " " + strconv.Itoa(mySlice[i])
		}
		return buff
	}

	flag := "YES"
	for i := 0; i < n-1; i++ {
		if value[i] >= value[i+1] {
			flag = "NO"
			break
		}
	}
	if flag == "YES" {
		fmt.Println("YES")
		fmt.Println(toString(result))
	} else if flag == "NO" {
		fmt.Println("NO")
	}

}
