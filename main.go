package main

import "github.com/files/filebasics"

func main() {
	fileName := "test.txt"
	filebasics.CreateFile(fileName)
	filebasics.TruncateFile(fileName)
}
