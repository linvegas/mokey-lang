package main

import (
    "bufio"
    "fmt"
    "io"
    "monkey/lexer"
    "monkey/token"
)

func StartREPL(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(">> ")
        scanned := scanner.Scan()
        if !scanned {
            return
        }
        line := scanner.Text()
        l := lexer.New(line)
        for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
            fmt.Printf("{'%v', %v}\n", t.Literal, t.Type)
        }
    }
}
