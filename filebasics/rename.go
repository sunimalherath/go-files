package filebasics

import (
	"os"

	"github.com/sunimalherath/go-files/errorhandling"
)

// RenameFile - Both rename and Move a file are the same
// To rename provide current path and name - or just the name if in same folder and provide new name and new path.
// Just rename => old name, new name
// Move without rename => old name, new path + name
// Move and Rename => old name, new path + new name
func RenameFile(oldName, newName string) {
	err := os.Rename(oldName, newName)
	errorhandling.ReportError(err)
}
