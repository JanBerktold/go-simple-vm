package vm

import (
	"testing"
	"fmt"
)

func TestAddSub(t *testing.T) {

	vm := &VM{}

	result := vm.run([]int{
		PUSH, 21,
		PUSH, 2,
		ADD,
		PUSH, 1,
		SUB,
		RET,
	})

	if result != 22 {
		t.Fail()
	}
}

func TestMulDiv(t *testing.T) {

	vm := &VM{}

	result := vm.run([]int{
		PUSH, 21,
		PUSH, 2,
		MUL,
		PUSH, 2,
		DIV,
		RET,
	})

	if result != 21 {
		t.Fail()
	}
}

func TestLoop(t *testing.T) {

	vm := &VM{}

	result := vm.run([]int{
		PUSH, 2,
		SLOOP,
		PUSH, 10,
		ADD,
		LOOP, 5,
		RET,
	})

	if result != 52 {
		t.Fail()
	}

}

func TestJNE(t *testing.T) {

	vm := &VM{}

	result := vm.run([]int{
		PUSH, 2,
		PUSH, 1,
		ADD,
		PUSH, 10,
		JNE, 2,
		RET,
	})

	if result != 10 {
		t.Fail()
	}

}

func TestJG(t *testing.T) {

	vm := &VM{}

	result := vm.run([]int{
		PUSH, 2,
		PUSH, 1,
		ADD,
		PUSH, 10,
		JG, 2,
		RET,
	})

	if result != 10 {
		t.Fail()
	}

}

func TestJL(t *testing.T) {

	vm := &VM{}

	result := vm.run([]int{
		PUSH, 40,
		PUSH, 2,
		SUB,
		PUSH, 20,
		JL, 2,
		RET,
	})

	fmt.Println(result)
	if result != 10 {
		t.Fail()
	}

}