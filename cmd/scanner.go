package main

import (
	"strconv"
)

type scanner struct {
	line       int
	current    int
	input      []rune
	stringMode bool
}

func (receiver scanner) isAtEnd() bool {
	return receiver.current >= len(receiver.input)
}

func advance(scanner *scanner) rune {
	scanner.current++
	return scanner.input[scanner.current-1]
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func makeASCII(r rune) Token {
	return Token{
		tokenType: TokenCharacter,
		char:      r,
	}
}

func makeNumber(r rune) Token {
	num, _ := strconv.ParseInt(string(r), 10, 8)
	return Token{
		tokenType: TokenNumber,
		number:    int(num),
	}
}

func makeToken(tokenType TokenType) Token {
	return Token{tokenType: tokenType}
}

func MakeTokenFromRune(c rune) Token {
	switch c {
	case '+':
		return makeToken(TokenPlus)
	case '-':
		return makeToken(TokenMinus)
	case '*':
		return makeToken(TokenStar)
	case '/':
		return makeToken(TokenForwardSlash)
	case '%':
		return makeToken(TokenPercent)
	case '!':
		return makeToken(TokenBang)
	case '`':
		return makeToken(TokenTick)
	case '>':
		return makeToken(TokenRight)
	case '<':
		return makeToken(TokenLeft)
	case '^':
		return makeToken(TokenUp)
	case 'v':
		return makeToken(TokenDown)
	case '?':
		return makeToken(TokenQuestion)
	case '_':
		return makeToken(TokenUnderscore)
	case '|':
		return makeToken(TokenPipe)
	case '"':
		return makeToken(TokenQuote)
	case ':':
		return makeToken(TokenColon)
	case '\\':
		return makeToken(TokenBackSlash)
	case '$':
		return makeToken(TokenDollar)
	case '.':
		return makeToken(TokenDot)
	case ',':
		return makeToken(TokenComma)
	case '#':
		return makeToken(TokenHash)
	case 'p':
		return makeToken(TokenPut)
	case 'g':
		return makeToken(TokenGet)
	case '&':
		return makeToken(TokenAmpersand)
	case '~':
		return makeToken(TokenTilda)
	case '@':
		return makeToken(TokenAt)
	case 10:
		return makeToken(TokenNewLine)
	default:
		if isDigit(c) {
			return makeNumber(c)
		}
		return makeToken(TokenWhitespace)
	}
}

func scanToken(scanner *scanner) Token {
	c := advance(scanner)
	if scanner.stringMode && c != '"' {
		return makeASCII(c)
	}
	if c == '"' {
		scanner.stringMode = !scanner.stringMode
	}
	return MakeTokenFromRune(c)
}

func NewScanner(input string) *scanner {
	return &scanner{
		line:    0,
		current: 0,
		input:   []rune(input),
	}
}

func addToken(program *Program, t Token) {
	line := program.lines[program.rows-1]
	line.length = line.length + 1
	line.source = append(make([]Token, 0, len(line.source)+1), line.source...)
	line.source = append(line.source, t)
	program.lines[program.rows-1] = line
}

func newLine(program *Program, scanner *scanner) {
	scanner.line++
	program.rows = program.rows + 1
	program.lines = append(make([]Line, 0, len(program.lines)+1), program.lines...)
	program.lines = append(program.lines, Line{
		source: []Token{},
		length: 0,
	})
}

func ScanInput(input string) *Program {
	scanner := NewScanner(input)
	program := NewProgram()
	for !scanner.isAtEnd() {
		token := scanToken(scanner)
		if token.tokenType == TokenNewLine {
			newLine(program, scanner)
			continue
		} else {
			addToken(program, token)
		}
	}

	// Adjust the maximum width of each row to be padded with whitespace to match the longest line
	// of the provided program. Firstly by finding the maximum length. This isn't the most tremendously efficient
	// approach, but solves the immediate problem.
	max := 0
	for _, line := range program.lines {
		if line.length > max {
			max = line.length
		}
	}
	// Now pad all lines to the maximum length
	for idx, line := range program.lines {
		if line.length < max {
			for i := 0; i < max-line.length; i++ {
				line.source = append(make([]Token, 0, len(line.source)+1), line.source...)
				line.source = append(line.source, makeToken(TokenWhitespace))
				program.lines[idx] = line
			}
		}
	}
	return program
}
