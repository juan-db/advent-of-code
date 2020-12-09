package main

import (
	"fmt"
	"strconv"
	"strings"
)

type operation int

const (
	acc operation = iota
	jmp
	nop
)

type instruction struct {
	op   operation
	arg  int
	hits int // amount of times this instruction has been executed
}

type virtualMachine struct {
	ip           int // instruction pointer
	acc          int
	instructions []*instruction
}

func parseInstruction(line string) *instruction {
	tokens := strings.Split(line, " ")

	var op operation
	switch tokens[0] {
	case "acc":
		op = acc

	case "jmp":
		op = jmp

	case "nop":
		op = nop
	}

	arg, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(fmt.Errorf("could not parse argument: %v", tokens[1]))
	}

	return &instruction{op, arg, 0}
}

func (vm *virtualMachine) executeNext() (done bool) {
	instr := vm.instructions[vm.ip]

	if instr.hits > 0 {
		fmt.Println(vm.acc)
		return true
	}

	switch instr.op {
	case acc:
		vm.acc += instr.arg
		vm.ip += 1

	case jmp:
		vm.ip += instr.arg

	case nop:
		vm.ip += 1
	}

	instr.hits += 1
	return false
}

func main() {
	vm := virtualMachine{0, 0, []*instruction{}}
	ReadInputFileByLine(func(line string) {
		vm.instructions = append(vm.instructions, parseInstruction(line))
	})

	// Don't want an infinite loop
	const max = 1000
	for i := 0; i < max; i++ {
		if vm.executeNext() {
			return
		}
	}
	panic(fmt.Errorf("executed %v instructions without hitting any instruction twice", max))
}
