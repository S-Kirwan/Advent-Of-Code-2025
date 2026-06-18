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
		if len(id)%2 != 0 {
			low++
			continue
		}
		if id[0:len(id)/2] == id[len(id)/2:] {
			total += low
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
	eof_erros := 0
	for {
		id, err := reader.ReadString(',')
		if err != nil && (err.Error() != "EOF" || eof_erros != 0) {
			break
		} else if err != nil && err.Error() == "EOF" {
			eof_erros++
		}
		low, high := parse_id(id)
		total += sum_invalid_ids(low, high)
	}
	fmt.Printf("total = %d\n", total)
}
