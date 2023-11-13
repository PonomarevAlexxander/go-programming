package main

import (
	_ "embed"
)

//go:embed text.txt
var fileString string

func main() {
	println(fileString)
}
