package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		if n <= 3 {
			return true
		}
		if n%2 == 0 || n%3 == 0 {
			return false
		}
		for i := 5; i*i <= n; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				return false
			}
		}
		return true
	}

	countDivisors := func(n int) int {
		count := 0
		for i := 1; i*i <= n; i++ {
			if n%i == 0 {
				if i*i == n {
					count++
				} else {
					count += 2
				}
			}
		}
		return count
	}

	flag := 0
	for i := l; i <= r; i++ {
		if i > 1 {
			divCount := countDivisors(i)
			if divCount > 2 && isPrime(divCount) {
				flag++
			}
		}
	}

	fmt.Println(flag)
}
