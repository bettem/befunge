## Befunge Interpreter

Implements a Befunge Interpreter in Go. 

## Language Description

The following definition is taken from the [esolange wiki](https://esolangs.org/wiki/Befunge).

Befunge is a two-dimensional esoteric programming language invented in 1993 by Chris Pressey with the goal of being 
as difficult to compile as possible. Code is laid out on a two-dimensional grid of instructions, and execution can 
proceed in any direction of that grid.

### Operations
The following are the support operations of the language:
 - `0-9`	Push this number on the stack
 - `+`	Addition: Pop a and b, then push a+b
 - `-`	Subtraction: Pop a and b, then push b-a
 - `*`	Multiplication: Pop a and b, then push a*b
 - `/`	Integer division: Pop a and b, then push b/a, rounded towards 0.
 - `%`	Modulo: Pop a and b, then push the remainder of the integer division of b/a.
 - `!`	Logical NOT: Pop a value. If the value is zero, push 1; otherwise, push zero.
 - `\``  Greater than: Pop a and b, then push 1 if b>a, otherwise zero.
 - `>`	Start moving right
 - `<`	Start moving left
 - `^`	Start moving up
 - `v`	Start moving down
 - `?`	Start moving in a random cardinal direction
 - `_`	Pop a value; move right if value=0, left otherwise
 - `|`	Pop a value; move down if value=0, up otherwise
 - `"`	Start string mode: push each character's ASCII value all the way up to the next "
 - `:`	Duplicate value on top of the stack
 - `\`	Swap two values on top of the stack
 - `$`	Pop value from the stack and discard it
 - `.`	Pop value and output as an integer followed by a space
 - `,`	Pop value and output as ASCII character
 - `#`	Bridge: Skip next cell
 - `p`	A "put" call (a way to store a value for later use). Pop y, x, and v, then change the character at (x,y) in the program to the character with ASCII value v
 - `g`	A "get" call (a way to retrieve data in storage). Pop y and x, then push ASCII value of the character at that position in the program
 - `&`	Ask user for a number and push it
 - `~`	Ask user for a character and push its ASCII value
 - `@`	End program
(space)	No-op. Does nothing

### Sample Programs

Some sample programs are include in the "./test" directory. As per the wikipedia entry, a simple hello world program
might read:
```befunge
>              v
v  ,,,,,"Hello"<
>48*,          v
v,,,,,,"World!"<
>25*,@
```

To calculate the factorial of a number, a program might look like:
```befunge
&>:1-:v v *_$.@ 
 ^    _$>\:^
```