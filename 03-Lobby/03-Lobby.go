package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find_biggest_voltage(bank string) int {
	biggest_voltage := 0
	var highest byte = '0'
	for i := 0; i < len(bank)-1; i++ {
		if bank[i] > highest {
			highest = bank[i]
		}
	}
	value, err := strconv.Atoi(string(highest))
	if err != nil {
		fmt.Printf("Atoi error = %s\n", err)
		return (-1)
	}
	biggest_voltage += value * 10
	for i := 0; i < len(bank); i++ {
		if bank[i] == highest {
			bank = bank[i+1:]
			break
		}
	}
	highest = '0'
	for i := 0; i < len(bank); i++ {
		if bank[i] > highest {
			highest = bank[i]
		}
	}
	value, err = strconv.Atoi(string(highest))
	if err != nil {
		fmt.Printf("Atoi error = %s\n", err)
		return (-1)
	}
	biggest_voltage += value
	return biggest_voltage
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		bank := scanner.Text()
		total += find_biggest_voltage(bank)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error = %s\n", err)
	}
	fmt.Printf("total = %d\n", total)
}
