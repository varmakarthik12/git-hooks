package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello")
	m := []byte(`#!/bin/bash
echo "Test Hook"`)
	ioutil.WriteFile(".git/hooks/pre-commit", m, 0644) //TODO:
}
