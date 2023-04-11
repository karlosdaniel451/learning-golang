package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SimpleAssembler(program []string) map[string]int {
	registers := make(map[string]int)

	for instructionCounter := 0; instructionCounter < len(program); instructionCounter++ {
		instruction := program[instructionCounter]
		tokens := strings.Split(instruction, " ")

		switch {
		case tokens[0] == "mov":
			paramX := tokens[1]
			paramY := tokens[2]
			ExecuteMov(paramX, paramY, registers)

		case tokens[0] == "inc":
			paramX := tokens[1]
			ExecuteInc(paramX, registers)

		case tokens[0] == "dec":
			paramX := tokens[1]
			ExecuteDec(paramX, registers)

		case tokens[0] == "jnz":
			paramX := tokens[1]
			paramY := tokens[2]
			ExecuteJnz(paramX, paramY, &instructionCounter, registers)
		}
	}

	return registers
}

func ExecuteMov(paramX string, paramY string, registers map[string]int) {
	// If paramX is a constant value.
	if paramYParsed, err := strconv.Atoi(paramY); err == nil {
		registers[paramX] = paramYParsed
		return
	}

	// If paramX is a register.
	registers[paramX] = registers[paramY]
}

func ExecuteInc(paramX string, registers map[string]int) {
	registers[paramX]++
}

func ExecuteDec(paramX string, registers map[string]int) {
	registers[paramX]--
}

func ExecuteJnz(paramX string, paramY string, instructionCounter *int, registers map[string]int) {
	var paramXValue int
	var paramYValue int

	// If paramX is a constant value.
	if paramXParsed, err := strconv.Atoi(paramX); err == nil {
		paramXValue = paramXParsed
	} else { // If paramX is a register.
		paramXValue = registers[paramX]
	}

	if paramXValue == 0 {
		return
	}

	// If paramY is a constant value.
	if paramYParsed, err := strconv.Atoi(paramY); err == nil {
		paramYValue = paramYParsed
	} else { // If paramY is a register.
		paramYValue = registers[paramY]
	}

	// Update instruction counter
	*instructionCounter += paramYValue - 1
}

func main() {
	program := []string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"}
	fmt.Println(SimpleAssembler(program))
}
