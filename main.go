package main

import (
	"github.com/sunimalherath/go-files/filebasics"
)

func main() {
	fileName := "test.txt"
	// Create file
	//filebasics.CreateFile(fileName)

	// Truncate the file
	// filebasics.TruncateFile(fileName)

	// Get the file info
	//filebasics.GetFileInfo(fileName)

	// Rename | Move file
	//newName := "test2.txt"
	//filebasics.RenameFile(fileName, newName)

	// Delete a file
	//fileToDelete := "test2.txt"
	//filebasics.DeleteFile(fileToDelete)

	// Open and Close file with os.Open() and os.OpenFile()
	//filebasics2.OpenAndCloseFile(fileName)

	// If file found
	//filebasics.CheckFileExistance(fileName)

	// If file not found
	//fileName2 := "test2.txt"
	//filebasics.CheckFileExistance(fileName2)

	// Check file's read & write permissions
	filebasics.CheckFileRWPermission(fileName)
}
