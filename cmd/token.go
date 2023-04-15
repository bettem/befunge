package main

import "fmt"

type Token struct {
	tokenType TokenType
	char      rune
	number    int
}

func (r Token) ASCII() int {
	if r.tokenType == TokenCharacter {
		return int(r.char)
	} else if r.tokenType == TokenNumber {
		return int([]rune(fmt.Sprintf("%d", r.number))[0])
	} else {
		return int(r.char)
	}
}

func (r Token) Print() {
	if r.tokenType == TokenCharacter {
		fmt.Printf("%c", r.char)
	} else if r.tokenType == TokenNumber {
		fmt.Printf("%d", r.number)
	} else if r.tokenType == TokenNewLine {
		fmt.Printf("\n")
	} else {
		fmt.Printf("%c", r.tokenType.print())
	}
}
