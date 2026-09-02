package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find_biggest_voltage(bank string) int {
	biggest_voltage := 0
	for j := 12; j > 0; j-- {
		var highest byte = '0'
		for i := 0; i <= len(bank)-j; i++ {
			if bank[i] > highest {
				highest = bank[i]
			}
		}
		value, err := strconv.Atoi(string(highest))
		if err != nil {
			fmt.Printf("Atoi error = %s\n", err)
			return (-1)
		}
		multiplier := 1
		for k := j; k > 1; k-- {
			multiplier *= 10
		}
		biggest_voltage += value * multiplier
		for i := 0; i < len(bank); i++ {
			if bank[i] == highest {
				bank = bank[i+1:]
				break
			}
		}
	}
	// fmt.Printf("biggest_voltage = %d\n", biggest_voltage)
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
