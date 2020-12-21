package filebasics

import (
	"fmt"
	"os"

	"github.com/sunimalherath/go-files/errorhandling"
)

// GetFileInfo - using the os.Stat to obtain file metadata
func GetFileInfo(fileName string) {
	fileInfo, err := os.Stat(fileName)
	errorhandling.ReportError(err)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
}
