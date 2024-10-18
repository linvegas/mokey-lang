package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("Monkey programming language REPL!\n")
    StartREPL(os.Stdin, os.Stdout)
}
