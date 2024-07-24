package main

import (
	"os"
	"testing"

	_ "github.com/christianhturner/go-interpreter/lexer"
	_ "github.com/christianhturner/go-interpreter/parser"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()

	os.Exit(exitCode)
}
