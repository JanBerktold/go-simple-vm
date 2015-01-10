package vm

import (
	"fmt"
	"strconv"
)

type VM struct {
	log bool

	code []int 
	pc int

	stack []int
	sp int

	loop int
}

const (
	PUSH = iota
	PRINT
	RET // return the last stack member as the result of the executed code
	ADD
	SUB
	MUL
	DIV
	JMP // jump
	JNE // jump on inequality
	JG // jump if greater
	JGE // jump if greater or equal
	JL // jump if less
	JLE // jump if less or equal
	SLOOP // marks the start of a loop
	LOOP // loops the code from here to the last SLOOP for X times
)

var (
	incName = map[int]string{
		PUSH: "PUSH",
		PRINT: "PRINT",
		RET: "RET",
		ADD: "ADD",
		SUB: "SUB",
		MUL: "MUL",
		DIV: "DIV",
		JMP: "JMP",
		JNE: "JNE",
		JG: "JG",
		JGE: "JGE",
		JL: "JL",
		JLE: "JLE",
		SLOOP: "SLOOP",
		LOOP: "LOOP",
	}
)

// prints the current state of the VM to the console
func (v *VM) trace() {
	fmt.Printf("%04d - INSTRUCTION \"%v\" \n", v.pc, incName[v.code[v.pc]])
	fmt.Println(v.stack[0:v.sp])
}

func (v *VM) run(c []int) int {

	v.stack = make([]int, 100)
	v.sp = 0

	v.code = c
	v.pc = 0

	for {
		if v.log {
			v.trace()
		}

		op := v.code[v.pc]
		v.pc++

		switch op {
		case PUSH:
			if v.log {
				fmt.Println("PUSH: " + strconv.Itoa(v.code[v.pc]))
			}
			v.stack[v.sp] = v.code[v.pc]
			v.sp++
			v.pc++
		case PRINT:
			v.sp--
		case RET:
			return v.stack[v.sp - 1]
		case ADD:
			v.sp--
			v.stack[v.sp - 1] = v.stack[v.sp] + v.stack[v.sp - 1]
		case SUB:
			v.sp--
			v.stack[v.sp - 1] = v.stack[v.sp - 1] - v.stack[v.sp]
		case MUL:
			v.sp--
			v.stack[v.sp - 1] = v.stack[v.sp - 1] * v.stack[v.sp]
		case DIV:
			v.sp--
			v.stack[v.sp - 1] = v.stack[v.sp - 1] / v.stack[v.sp]
		case JMP:
			v.pc = v.code[v.pc]
		case JNE: 
			v.sp--
			if (v.stack[v.sp] != v.stack[v.sp - 1]) {
				v.pc = v.code[v.pc]
			} else {
				v.pc++
			}
		case JG:
			v.sp--
			if (v.stack[v.sp] > v.stack[v.sp - 1]) {
				v.pc = v.code[v.pc]
			} else {
				v.pc++
			}
		case JGE:
			v.sp--
			if (v.stack[v.sp] >= v.stack[v.sp - 1]) {
				v.pc = v.code[v.pc]
			} else {
				v.pc++
			}
		case JL:
			v.sp--
			if (v.stack[v.sp] < v.stack[v.sp - 1]) {
				v.pc = v.code[v.pc]
			} else {
				v.pc++
			}	
		case JLE:
			v.sp--
			if (v.stack[v.sp] <= v.stack[v.sp - 1]) {
				v.pc = v.code[v.pc]
			} else {
				v.pc++
			}		
		case SLOOP:
			v.loop = v.pc - 1
		case LOOP:
			// bigger than 1, as we already ran trough the loop once
			if v.code[v.pc] > 1 {
				v.code[v.pc]--
				v.pc = v.loop
			} else {
				v.pc++
			}
		}

		if v.pc >= len(v.code) {
			return 0
		}
	}

}