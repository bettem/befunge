package main

import (
	"fmt"
	"math/rand"
)

type VM struct {
	Direction        Direction
	row              int
	col              int
	running          bool
	stack            []Token
	stackSize        int
	stringMode       bool
	encounteredError bool
	program          *Program
}

func NewVM(program *Program) *VM {
	return &VM{
		Direction:        DirectionRight,
		row:              0,
		col:              0,
		running:          true,
		stack:            make([]Token, 0),
		stackSize:        0,
		stringMode:       false,
		encounteredError: false,
		program:          program,
	}
}

func (vm *VM) directionRight() {
	vm.Direction = DirectionRight
}

func (vm *VM) directionDown() {
	vm.Direction = DirectionDown
}

func (vm *VM) directionLeft() {
	vm.Direction = DirectionLeft
}

func (vm *VM) directionUp() {
	vm.Direction = DirectionUp
}

func (vm *VM) directionRandom() {
	switch rand.Intn(4) {
	case 0:
		vm.directionRight()
	case 1:
		vm.directionLeft()
	case 2:
		vm.directionUp()
	case 3:
		vm.directionDown()
	}
}

func (vm *VM) push(t Token) {
	vm.stackSize++
	vm.stack = append(make([]Token, 0, vm.stackSize), vm.stack...)
	vm.stack = append(vm.stack, t)
}

func (vm *VM) pop() Token {
	// As a 'compliant' interpreter, we return zero when popping on an empty stack.
	if vm.stackSize == 0 {
		return Token{
			tokenType: TokenNumber,
			number:    0,
		}
	}
	vm.stackSize--
	token := vm.stack[vm.stackSize]
	vm.stack = append([]Token(nil), vm.stack[:vm.stackSize]...)
	return token
}

func (vm *VM) Advance() {
	switch vm.Direction {
	case DirectionDown:
		vm.row++
		break
	case DirectionRight:
		vm.col++
		break
	case DirectionLeft:
		vm.col--
		break
	case DirectionUp:
		vm.row--
		break
	}
}

func (vm *VM) Execute() {
	token := vm.program.lines[vm.row].source[vm.col]

	if vm.stringMode && token.tokenType != TokenQuote {
		vm.push(Token{
			tokenType: TokenNumber,
			number:    token.ASCII(),
		})
		return
	}

	switch token.tokenType {
	case TokenNumber:
		// Push this number on the stack
		vm.push(token)
		break
	case TokenPlus:
		// Addition: Pop a and b, then push a+b
		a := vm.pop()
		b := vm.pop()
		vm.push(Token{
			tokenType: TokenNumber,
			number:    a.number + b.number,
		})
		break
	case TokenMinus:
		// Subtraction: Pop a and b, then push b-a
		a := vm.pop()
		b := vm.pop()
		vm.push(Token{
			tokenType: TokenNumber,
			number:    b.number - a.number,
		})
	case TokenStar:
		// Multiplication: Pop a and b, then push a*b
		a := vm.pop()
		b := vm.pop()
		vm.push(Token{
			tokenType: TokenNumber,
			number:    a.number * b.number,
		})
		break
	case TokenForwardSlash:
		// Integer division: Pop a and b, then push b/a, rounded towards 0.
		a := vm.pop()
		b := vm.pop()
		vm.push(Token{
			tokenType: TokenNumber,
			number:    a.number / b.number,
		})
		break
	case TokenPercent:
		// Modulo: Pop a and b, then push the remainder of the integer division of b/a.
		a := vm.pop()
		b := vm.pop()
		vm.push(Token{
			tokenType: TokenNumber,
			number:    a.number % b.number,
		})
		break
	case TokenBang:
		// Logical NOT: Pop a value. If the value is zero, push 1; otherwise, push zero.
		a := vm.pop()
		if a.number == 0 {
			vm.push(Token{
				tokenType: TokenNumber,
				number:    1,
			})
		} else {
			vm.push(Token{
				tokenType: TokenNumber,
				number:    0,
			})
		}
		break
	case TokenTick:
		// Greater than: Pop a and b, then push 1 if b>a, otherwise zero.
		a := vm.pop()
		b := vm.pop()
		if b.number > a.number {
			vm.push(Token{
				tokenType: TokenNumber,
				number:    1,
			})
		} else {
			vm.push(Token{
				tokenType: TokenNumber,
				number:    0,
			})
		}
		break
	case TokenRight:
		// Start moving right
		vm.directionRight()
		break
	case TokenLeft:
		// Start moving left
		vm.directionLeft()
		break
	case TokenUp:
		// Start moving up
		vm.directionUp()
		break
	case TokenDown:
		//  Start moving down
		vm.directionDown()
		break
	case TokenQuestion:
		// Start moving in a random cardinal direction
		vm.directionRandom()
		break
	case TokenUnderscore:
		//  Pop a value; move right if value=0, left otherwise
		a := vm.pop()
		if a.number == 0 {
			vm.directionRight()
		} else {
			vm.directionLeft()
		}
		break
	case TokenPipe:
		// Pop a value; move down if value=0, up otherwise
		a := vm.pop()
		if a.number == 0 {
			vm.directionDown()
		} else {
			vm.directionUp()
		}
		break
	case TokenQuote:
		// Start string mode: push each character's ASCII value all the way up to the next "
		vm.stringMode = !vm.stringMode
		break
	case TokenColon:
		// Duplicate value on top of the stack
		a := vm.pop()
		vm.push(a)
		vm.push(a)
		break
	case TokenBackSlash:
		// Swap two values on top of the stack
		a := vm.pop()
		b := vm.pop()
		vm.push(a)
		vm.push(b)
		break
	case TokenDollar:
		// Pop value from the stack and discard it
		vm.pop()
		break
	case TokenDot:
		// Pop value and output as an integer followed by a space
		a := vm.pop()
		fmt.Printf("%d ", a.number)
		break
	case TokenComma:
		//  Pop value and output as ASCII character
		a := vm.pop()
		fmt.Printf("%c", rune(a.number))
		break
	case TokenHash:
		//  Bridge: Skip next cell
		vm.Advance()
		break
	case TokenPut:
		// A "put" call (a way to store a value for later use). Pop y, x, and v, then change the character at (x,y)
		// in the program to the character with ASCII value v
		y := vm.pop()
		x := vm.pop()
		v := vm.pop()
		vm.program.lines[x.number].source[y.number] = MakeTokenFromRune(v.char)
		break
	case TokenGet:
		// A "get" call (a way to retrieve data in storage). Pop y and x, then push ASCII value of the character at
		// that position in the program
		y := vm.pop()
		x := vm.pop()
		vm.push(Token{
			tokenType: TokenCharacter,
			char:      vm.program.lines[x.number].source[y.number].char,
		})
		break
	case TokenAmpersand:
		// Ask user for a number and push it
		var input int
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			fmt.Printf("\nUnable to read input at [%d,%d], %e", vm.row, vm.col, err)
			vm.encounteredError = true
			break
		}
		vm.push(Token{
			tokenType: TokenNumber,
			number:    input,
		})
		break
	case TokenTilda:
		// Ask user for a character and push its ASCII value
		var input rune
		_, err := fmt.Scanf("%c", &input)
		if err != nil {
			fmt.Printf("Unable to read input [%d,%d], %e", vm.row, vm.col, err)
			vm.encounteredError = true
			break
		}
		vm.push(Token{
			tokenType: TokenCharacter,
			char:      input,
		})
		break
	case TokenAt:
		// End program (space) No-op. Does nothing
		vm.running = false
		break
	}
}

func Interpret(p *Program) {
	vm := NewVM(p)
	for vm.running && !vm.encounteredError {
		vm.Execute()
		vm.Advance()
	}
}
