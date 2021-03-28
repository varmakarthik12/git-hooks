package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello")
	m := []byte(`#!/bin/bash
echo "Test Hook testing"`)
	ioutil.WriteFile(".git/hooks/pre-commit", m, 0777) //TODO:
}
