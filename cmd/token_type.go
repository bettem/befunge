package main

type TokenType int64

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenStar
	TokenForwardSlash
	TokenPercent
	TokenBang
	TokenTick
	TokenRight
	TokenLeft
	TokenUp
	TokenDown
	TokenQuestion
	TokenUnderscore
	TokenPipe
	TokenQuote
	TokenColon
	TokenBackSlash
	TokenDollar
	TokenDot
	TokenComma
	TokenHash
	TokenPut
	TokenGet
	TokenAmpersand
	TokenTilda
	TokenAt
	TokenWhitespace
	TokenNewLine
	TokenCharacter
)

func (t TokenType) print() rune {
	switch t {
	case TokenPlus:
		return '+'
	case TokenMinus:
		return '-'
	case TokenStar:
		return '*'
	case TokenForwardSlash:
		return '/'
	case TokenPercent:
		return '%'
	case TokenBang:
		return '!'
	case TokenTick:
		return '`'
	case TokenRight:
		return '>'
	case TokenLeft:
		return '<'
	case TokenUp:
		return '^'
	case TokenDown:
		return 'v'
	case TokenQuestion:
		return '?'
	case TokenUnderscore:
		return '_'
	case TokenPipe:
		return '|'
	case TokenQuote:
		return '"'
	case TokenColon:
		return ':'
	case TokenBackSlash:
		return '\\'
	case TokenDollar:
		return '$'
	case TokenDot:
		return '.'
	case TokenComma:
		return ','
	case TokenHash:
		return '#'
	case TokenPut:
		return 'p'
	case TokenGet:
		return 'g'
	case TokenAmpersand:
		return '&'
	case TokenTilda:
		return '~'
	case TokenAt:
		return '@'
	case TokenNewLine:
		return '\n'
	case TokenWhitespace:
		return ' '
	}
	return ' '
}
