# go-simple-vm
A very simple VM writen in Golang which can execute programs described as integer arrays.

Look into vm_test.go for examples of usage.

## Commands

| Command | Description |
| :------ | -----------:|
| PUSH |  av |
| PRINT | a |
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
