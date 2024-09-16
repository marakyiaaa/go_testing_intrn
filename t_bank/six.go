package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	times := make([]int, n)
	dependencies := make([][]int, n)
	degree := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		times[i], _ = strconv.Atoi(parts[0])
		for j := 1; j < len(parts); j++ {
			dep, _ := strconv.Atoi(parts[j])
			dependencies[i] = append(dependencies[i], dep-1)
			degree[dep-1]++
		}
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	minTime := make([]int, n)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, dep := range dependencies[node] {
			degree[dep]--
			if degree[dep] == 0 {
				queue = append(queue, dep)
			}
			minTime[dep] = max(minTime[dep], minTime[node]+times[node])
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		result = max(result, minTime[i]+times[i])
	}
	fmt.Println(result)
}
