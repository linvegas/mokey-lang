package ast

import (
    token "monkey/token"
    "testing"
)

func TestString(t *testing.T) {
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENTIFIER, Literal: "muh_var"},
                    Value: "muh_var",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENTIFIER, Literal: "not_your_var"},
                    Value: "not_your_var",
                },
            },
        },
    }

    if program.String() != "let muh_var = not_your_var;" {
        t.Errorf("program.String() wrong. got=%q", program.String())
    }
}
