package token

const (
    ILLEGAL     = "ILLEGAL"
    EOF         = "EOF"
    // Literals
    IDENTIFIER  = "IDENTIFIER"
    INT         = "INT"
    // Operators
    ASSIGN      = "ASSIGN"
    PLUS        = "PLUS"
    MINUS       = "MINUS"
    ASTERISK    = "ASTERISK"
    SLASH       = "SLASH"
    BANG        = "BANG"
    LESS_T      = "LESS_T"
    GREATER_T   = "GREATER_T"
    EQUAL       = "EQUAL"
    NOT_EQUAL   = "NOT_EQUAL"
    // Delimiters
    COMMA       = "COMMA"
    SEMICOLON   = "SEMICOLON"
    L_PAREN     = "L_PAREN"
    R_PAREN     = "R_PAREN"
    L_BRACE     = "L_BRACE"
    R_BRACE     = "R_BRACE"
    // Keywords
    FUNCTION    = "FUNCTION"
    LET         = "LET"
    TRUE        = "TRUE"
    FALSE       = "FALSE"
    IF          = "IF"
    ELSE        = "ELSE"
    RETURN      = "RETURN"
)

type TokenType string

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdentifier(id string) TokenType {
    if token, ok := keywords[id]; ok {
        return token
    }
    return IDENTIFIER
}

type Token struct {
    Type    TokenType
    Literal string
}
