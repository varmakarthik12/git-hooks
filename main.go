package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	m := []byte(`#!/bin/bash
echo "Test Hook testing"`)
	os.WriteFile(".git/hooks/pre-commit", m, 0777) //TODO:
	os.Chmod(".git/hooks/pre-commit", 0777)
}
