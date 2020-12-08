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

	fmt.Println(FixProgram(instructions))
}

func FixProgram(instructions []string) (accumulator int) {
	accumulator, executed, inLoop := RunProgram(instructions)

	for inLoop {
		fixedInstructions := make([]string, len(instructions))
		copy(fixedInstructions, instructions)

		pc := executed[0]
		executed = executed[1:]

		corrupted := fixedInstructions[pc]
		switch op := corrupted[:3]; op {
		case "nop":
			fixedInstructions[pc] = "jmp" + corrupted[3:]
		case "jmp":
			fixedInstructions[pc] = "nop" + corrupted[3:]
		}

		accumulator, _, inLoop = RunProgram(fixedInstructions)
	}

	return
}

func RunProgram(instructions []string) (accumulator int, executed []int, inLoop bool) {
	executedOp := make(map[int]struct{})

	for pc := 0; pc < len(instructions); {
		if _, ok := executedOp[pc]; ok {
			inLoop = true
			return
		}
		executedOp[pc] = struct{}{}
		val, _ := strconv.Atoi(instructions[pc][4:])

		switch op := instructions[pc][:3]; op {
		case "nop":
			executed = append(executed, pc)
			pc++
		case "acc":
			accumulator += val
			pc++
		case "jmp":
			executed = append(executed, pc)
			pc += val
		}
	}

	return
}
