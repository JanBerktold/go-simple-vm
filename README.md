# go-simple-vm
A very simple VM writen in Golang which can execute programs described as integer arrays.

Look into vm_test.go for examples of usage.

## Commands

| Command | Description | Arguments |
| :------: | ----------- | :-----------: |
| PUSH | Pushes the integer into the stack. | The number to be added to the stack |
| PRINT | Takes the last inserted integer out of the stack and prints it to the console. | - |
| RET | See PRINT but returns the integer as the result of the code execution. Halts the VM afterwards. | - |
| ADD | Adds up the last two inserted integers in the stack and writes its result back to the stack. | - |
| SUB | See ADD, but instead of adding it subtracts the two integers. | - |
| MUL | See ADD, but instead of adding them it multiplies the two integers. | - |
| DIV | See MUL, but instead of multiplying it divides the two integers. | - |
| JMP | Jumps to the supplied position. | The position |
| JNE | Jumps to the supplied position, if the last two elements in the stack are equal. Removes the last integer out of the stack. | The position |
| JG  | Jumps to the supplied position, if the last element in the stack is greater than the second-last. Removes the last integer out of the stack. | The position |
| JGE | Jumps to the supplied position, if the last element in the stack is greater than or equal to the second-last. Removes the last integer out of the stack. | The position |
| JL | Jumps to the supplied position, if the last element in the stack is less than the second-last. Removes the last integer out of the stack. | The position |
| JLE | Jumps to the supplied position, if the last element in the stack is less than or equal to the second-last. Removes the last integer out of the stack. | The position |
| SLOOP | Marks the beginning of a loop. Instruction itself is not doing anything. | - |
| LOOP | Repeates the code between the last SLOOP call and this statement for the specified times. | Amount of loop cycles |
