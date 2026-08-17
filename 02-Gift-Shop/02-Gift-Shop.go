package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum_invalid_ids(low int, high int) int {
	total := 0
	for low <= high {
		id := strconv.Itoa(low)
		for i := 1; i <= len(id)/2; i++ {
			if len(id)%i != 0 {
				continue
			}
			sequence := id[0:i]
			valid := false
			for j := 1; j*i <= len(id); j++ {
				if id[(j-1)*i:j*i] != sequence {
					break
				}
				if j*i == len(id) {
					total += low
					valid = true
				}
			}
			if valid == true {
				break
			}
		}
		low++
	}
	return total
}

func parse_id(id string) (int, int) {
	var low int
	var high int
	var err error

	i := 0
	for id[i] != '-' {
		i++
	}
	low, err = strconv.Atoi(id[0:i])
	if err != nil {
		fmt.Printf("Atoi error on '%s'\n", id[0:i])
	}
	high, err = strconv.Atoi(id[i+1 : len(id)-1])
	if err != nil {
		fmt.Printf("Atoi error on '%s'\n", id[i+1:len(id)-1])
	}
	return low, high
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	total := 0
	eof_errors := 0
	for eof_errors == 0 {
		id, err := reader.ReadString(',')
		if err != nil {
			if err.Error() == "EOF" {
				eof_errors++
			} else {
				break
			}
		}
		low, high := parse_id(id)
		total += sum_invalid_ids(low, high)
	}
	fmt.Printf("total = %d\n", total)
}
