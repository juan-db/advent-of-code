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

type status int

const (
	running status = iota
	exited
	loop
	error
)

type virtualMachine struct {
	ip           int // instruction pointer
	acc          int
	instructions []*instruction
	status       status
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

func (vm *virtualMachine) executeNext() {
	if vm.status != running {
		panic(fmt.Errorf("attempted to execute instruction after vm terminated into state: %v", vm.status))
	}

	instr := vm.instructions[vm.ip]

	if instr.hits > 0 {
		vm.status = loop
		return
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

	if vm.ip == len(vm.instructions) {
		vm.status = exited
	} else if vm.ip > len(vm.instructions) {
		vm.status = error
	}
}

func (vm *virtualMachine) clone() *virtualMachine {
	n := virtualMachine{0, 0, []*instruction{}, running}
	for _, i := range vm.instructions {
		n.instructions = append(n.instructions, &instruction{i.op, i.arg, 0})
	}
	return &n
}

func (vm *virtualMachine) isValid() bool {
	for i, max := 0, len(vm.instructions); i <= max; i++ {
		vm.executeNext()
		if vm.status != running {
			break
		}
	}

	return vm.status == exited
}

func main() {
	vm := virtualMachine{0, 0, []*instruction{}, running}
	ReadInputFileByLine(func(line string) {
		vm.instructions = append(vm.instructions, parseInstruction(line))
	})

	for i, v := range vm.instructions {
		var newInstr instruction
		switch v.op {
		case acc:
			continue

		case jmp:
			newInstr = instruction{nop, v.arg, 0}

		case nop:
			newInstr = instruction{jmp, v.arg, 0}
		}

		n := vm.clone()
		n.instructions[i] = &newInstr
		if n.isValid() {
			fmt.Println(n.acc)
			return
		}
	}
}
