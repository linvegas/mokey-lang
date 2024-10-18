package parser

import (
    "fmt"
    "monkey/ast"
    l "monkey/lexer"
    t "monkey/token"
)

const (
    _ int = iota
    LOWEST
    EQUALS
    LESSGREATER
    SUM
    PRODUCT
    PREFIX
    CALL
)

type Parser struct {
    lex *l.Lexer
    errors []string

    CurrentToken t.Token
    PeekToken t.Token

    PrefixParseFns map[t.TokenType]PrefixParseFn
    InfixParseFns  map[t.TokenType]InfixParseFn
}

func New(l *l.Lexer) *Parser {
    p := &Parser{
        lex: l,
        errors: []string{},
    }

    p.NextToken()
    p.NextToken()

    p.PrefixParseFns = make(map[t.TokenType]PrefixParseFn)
    p.RegisterPrefix(t.IDENTIFIER, p.ParseIdentifier)

    return p
}

func (p *Parser) ParseIdentifier() ast.Expression {
    return &ast.Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
}

func (p *Parser) Errors() []string {
    return p.errors
}

func (p *Parser) PeekErrors(t t.TokenType) {
    msg := fmt.Sprintf("Expected next token to be %s, but got %s instead",
        t, p.PeekToken.Type)
    p.errors = append(p.errors, msg)
}

func (p *Parser) NextToken() {
    p.CurrentToken = p.PeekToken
    p.PeekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.CurrentToken.Type != t.EOF {
        statement := p.ParseStatement()
        if statement != nil {
            program.Statements = append(program.Statements, statement)
        }
        p.NextToken()
    }

    return program
}

func (p *Parser) ParseStatement() ast.Statement {
    switch p.CurrentToken.Type {
    case t.LET:
        return p.ParseLetStatement()
    case t.RETURN:
        return p.ParseReturnStatement()
    default:
        return p.ParseExpressionStatement()
    }
}

func (p *Parser) ParseLetStatement() *ast.LetStatement {
    statement := &ast.LetStatement{Token: p.CurrentToken}

    if !p.ExpectPeek(t.IDENTIFIER) {
        return nil
    }

    statement.Name = &ast.Identifier{
        Token: p.CurrentToken,
        Value: p.CurrentToken.Literal,
    }

    if !p.ExpectPeek(t.ASSIGN) {
        return nil
    }

    for !p.CurrentTokenIs(t.SEMICOLON) {
        p.NextToken()
    }

    return statement
}

func (p *Parser) CurrentTokenIs(t t.TokenType) bool {
    return p.CurrentToken.Type == t
}

func (p *Parser) PeekTokenIs(t t.TokenType) bool {
    return p.PeekToken.Type == t
}

func (p *Parser) ExpectPeek(t t.TokenType) bool {
    if p.PeekTokenIs(t) {
        p.NextToken()
        return true
    } else {
        return false
    }
}

func (p *Parser) ParseReturnStatement() *ast.ReturnStatement {
    statement := &ast.ReturnStatement{Token: p.CurrentToken}

    p.NextToken()

    for !p.CurrentTokenIs(t.SEMICOLON) {
        p.NextToken()
    }

    return statement
}

type (
    PrefixParseFn func() ast.Expression
    InfixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) RegisterPrefix(tokeType t.TokenType, fn PrefixParseFn) {
    p.PrefixParseFns[tokeType] = fn
}

func (p *Parser) RegisterInfix(tokeType t.TokenType, fn InfixParseFn) {
    p.InfixParseFns[tokeType] = fn
}

func (p *Parser) ParseExpressionStatement() *ast.ExpressionStatement {
    statement := &ast.ExpressionStatement{Token: p.CurrentToken}

    statement.Expression = p.ParseExpression(LOWEST)

    for !p.CurrentTokenIs(t.SEMICOLON) {
        p.NextToken()
    }

    return statement
}

func (p *Parser) ParseExpression(precendence int) ast.Expression {
    prefix := p.PrefixParseFns[p.CurrentToken.Type]

    if prefix == nil {
        return nil
    }

    leftExp := prefix()

    return leftExp
}
