package main

import (
	"bufio"
	"fmt"
	"os"
)

func count_current_row(i int, current string) int {
	adjacent := 0
	if i != 0 && current[i-1] == '@' {
		adjacent++
	}
	if i < len(current)-1 && current[i+1] == '@' {
		adjacent++
	}
	return adjacent
}

func count_prev_next_row(i int, prev_next string) int {
	if len(prev_next) == 0 {
		return 0
	}
	adjacent := 0
	if i == 0 {
		for j := range 2 {
			if prev_next[j] == '@' {
				adjacent++
			}
		}
	} else if i == len(prev_next)-1 {
		for j := i - 1; j <= i; j++ {
			if prev_next[j] == '@' {
				adjacent++
			}
		}
	} else {
		for j := i - 1; j <= i+1; j++ {
			if prev_next[j] == '@' {
				adjacent++
			}
		}
	}
	return adjacent
}

func count_accessible_rolls(prev string, current string, next string) int {
	accessible := 0
	for i := 0; i < len(current); i++ {
		adjacent := 0
		if current[i] == '@' {
			adjacent += count_prev_next_row(i, prev)
			adjacent += count_current_row(i, current)
			adjacent += count_prev_next_row(i, next)
			if adjacent < 4 {
				accessible++
			}
		}
	}
	return accessible
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	var prev string = ""
	var current string = ""
	var next string = ""
	if scanner.Scan() {
		current = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error = %s\n", err)
		return
	}
	for scanner.Scan() {
		next = scanner.Text()
		total += count_accessible_rolls(prev, current, next)
		prev = current
		current = next
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error = %s\n", err)
		return
	}
	next = ""
	total += count_accessible_rolls(prev, current, next)
	fmt.Printf("total = %d\n", total)
}
