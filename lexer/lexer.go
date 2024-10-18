package lexer

import (
    t "monkey/token"
)

type Lexer struct {
    Input        string
    Position     int
    ReadPosition int
    Char         byte
}

func New(input string) *Lexer {
    l := &Lexer{Input: input}
    l.ReadChar()
    return l
}

func (l *Lexer) ReadChar() {
    if l.ReadPosition >= len(l.Input) {
        l.Char = 0 // NULL or EOF
    } else {
        l.Char = l.Input[l.ReadPosition]
    }

    l.Position = l.ReadPosition
    l.ReadPosition += 1
}

func (l *Lexer) ReadIndentifier() string {
    position := l.Position
    for isLetter(l.Char) {
        l.ReadChar()
    }
    return l.Input[position:l.Position]
}

func (l *Lexer) ReadNumber() string {
    position := l.Position
    for isDigit(l.Char) {
        l.ReadChar()
    }
    return l.Input[position:l.Position]
}

func (l *Lexer) PeekChar() byte {
    if l.ReadPosition >= len(l.Input) {
        return 0 // NULL or EOF
    } else {
        return l.Input[l.ReadPosition]
    }
}

func (l *Lexer) NextToken() t.Token {
    var token t.Token

    // Skipping whitespaces
    for l.Char == ' ' || l.Char == '\t' || l.Char == '\n' || l.Char == '\r' {
        l.ReadChar()
    }

    switch l.Char {
    case '=':
        if l.PeekChar() == '=' {
            c := l.Char
            l.ReadChar()
            literal := string(c) + string(l.Char)
            token = t.Token{Type: t.EQUAL, Literal: literal}
        } else {
            token = t.Token{Type: t.ASSIGN, Literal: string(l.Char)}
        }
    case '+':
        token = t.Token{Type: t.PLUS, Literal: string(l.Char)}
    case '-':
        token = t.Token{Type: t.MINUS, Literal: string(l.Char)}
    case '*':
        token = t.Token{Type: t.ASTERISK, Literal: string(l.Char)}
    case '/':
        token = t.Token{Type: t.SLASH, Literal: string(l.Char)}
    case '!':
        if l.PeekChar() == '=' {
            c := l.Char
            l.ReadChar()
            literal := string(c) + string(l.Char)
            token = t.Token{Type: t.NOT_EQUAL, Literal: literal}
        } else {
            token = t.Token{Type: t.BANG, Literal: string(l.Char)}
        }
    case '<':
        token = t.Token{Type: t.LESS_T, Literal: string(l.Char)}
    case '>':
        token = t.Token{Type: t.GREATER_T, Literal: string(l.Char)}
    case ';':
        token = t.Token{Type: t.SEMICOLON, Literal: string(l.Char)}
    case ',':
        token = t.Token{Type: t.COMMA, Literal: string(l.Char)}
    case '(':
        token = t.Token{Type: t.L_PAREN, Literal: string(l.Char)}
    case ')':
        token = t.Token{Type: t.R_PAREN, Literal: string(l.Char)}
    case '{':
        token = t.Token{Type: t.L_BRACE, Literal: string(l.Char)}
    case '}':
        token = t.Token{Type: t.R_BRACE, Literal: string(l.Char)}
    case 0:
        token = t.Token{Type: t.EOF, Literal: "\\0"}
    default:
        if isLetter(l.Char) {
            token.Literal = l.ReadIndentifier()
            token.Type = t.LookupIdentifier(token.Literal)
            return token
        } else if isDigit(l.Char) {
            token.Type = t.INT
            token.Literal = l.ReadNumber()
            return token
        } else {
            token =  t.Token{Type: t.ILLEGAL, Literal: string(l.Char)}
        }
    }

    l.ReadChar()
    return token
}

func isLetter(char byte) bool {
    return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
    return '0' <= char && char <= '9'
}

