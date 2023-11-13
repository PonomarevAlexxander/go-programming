package main

import (
	"embed"
)

//go:embed text.txt
var fileString string

//go:embed text.txt
var fileByte []byte

//go:embed testfiles/*.hash
var folder embed.FS

func main() {
	println(fileString)
	println(string(fileByte))

	file1, _ := folder.ReadFile("testfiles/file1.hash")
	println(string(file1))
	file2, _ := folder.ReadFile("testfiles/file2.hash")
	println(string(file2))
}
