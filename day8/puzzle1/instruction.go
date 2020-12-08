package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var instructions []string

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		instructions = append(instructions, scanner.Text())
	}

	fmt.Println(RunProgram(instructions))
}

func RunProgram(instructions []string) (accumulator int) {

	executedOp := make(map[int]struct{})
	for pc := 0; pc < len(instructions); {
		if _, ok := executedOp[pc]; ok {
			return
		}
		executedOp[pc] = struct{}{}
		val, _ := strconv.Atoi(instructions[pc][4:])

		switch op := instructions[pc][:3]; op {
		case "nop":
			pc++
		case "acc":
			accumulator += val
			pc++
		case "jmp":
			pc += val
		}
	}
	return
}
