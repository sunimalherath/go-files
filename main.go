package main

import "github.com/sunimalherath/go-files/filebasics"

func main() {
	//fileName := "test.txt"
	// Create file
	// filebasics.CreateFile(fileName)

	// Truncate the file
	// filebasics.TruncateFile(fileName)

	// Get the file info
	//filebasics.GetFileInfo(fileName)

	// Rename | Move file
	//newName := "test2.txt"
	//filebasics.RenameFile(fileName, newName)

	// Delete a file
	fileToDelete := "test2.txt"
	filebasics.DeleteFile(fileToDelete)
}
