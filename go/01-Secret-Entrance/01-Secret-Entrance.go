package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func rotateDial(instruction string, pointing int) int {
	power, _ := strconv.Atoi(instruction[1 : len(instruction)-1])
	power = power % 100
	if instruction[0] == 'L' {
		if pointing-power < 0 {
			pointing = 100 - (power - pointing)
		} else {
			pointing = pointing - power
		}
	} else {
		if pointing+power > 99 {
			pointing = (pointing + power) % 100
		} else {
			pointing = pointing + power
		}
	}
	return pointing
}

func crackSafe() int {
	reader := bufio.NewReader(os.Stdin)
	pointing := 50
	total := 0
	for true {
		instruction, err := reader.ReadString('\n')
		pointing = rotateDial(instruction, pointing)
		if pointing == 0 {
			total++
		}
		if err != nil {
			break
		}
	}
	return (total)
}

func main() {
	fmt.Printf("%d\n", crackSafe())
}
