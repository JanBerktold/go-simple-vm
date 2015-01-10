# go-simple-vm
A very simple VM writen in Golang which can execute programs described as integer arrays.

Look into vm_test.go for examples of usage.

## Commands

| Command | Description |
| ------ | -----------|
| PUSH | Pushes the integer which it is followed by in the programm code into the stack. |
| PRINT | Takes the last inserted integer out of the stack and prints it to the console. |
| RET | See PRINT but returns the integer as the result of the code execution. Halts the VM afterwards. |
| ADD | Adds up the last two inserted integers in the stack and writes its result back to the stack. |
| SUB | See ADD, but instead of adding it subtracts the two integers. |
| MUL | See ADD, but instead of adding them it multiplies the two integers. |
| DIV | See MUL, but instead of multiplying it divides the two integers. |
| JMP | 
| JNE |
| JG  |
| JGE |
| JL |
| JLE |
| SLOOP |
| LOOP |
