package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	n := len(str)
	for i := 0; i < n; i++ {
		first := 0
		second := 0
		for ; i < n && str[i] >= '0' && str[i] <= '9'; i++ {
			d := int(str[i] - '0')
			first = first*10 + d
		}
		if i < n && str[i] == '-' {
			i++
			for ; i < n && str[i] >= '0' && str[i] <= '9'; i++ {
				b := int(str[i] - '0')
				second = second*10 + b
			}
			for ; first <= second; first++ {
				fmt.Print(first)
				if first != second || i != n {
					fmt.Print(" ")
				}
			}
		} else {
			fmt.Print(first)
			if i != n {
				fmt.Print(" ")
			}
		}
	}
}
