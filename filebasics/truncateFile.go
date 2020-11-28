package filebasics

import (
	"github.com/files/errorhandling"
	"os"
)

func TruncateFile(fileName string) {
	// Truncate file to 100 bytes
	// If the original file is less than 100 bytes, it will filled with null bytes
	// If the original file is larger than 100 bytes, additional bytes over 100 will be lost
	// If you pass 0, the original file will be a completely empty file.

	err := os.Truncate(fileName, 100)
	errorhandling.ReportError(err)
}
