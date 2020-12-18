package filebasics

import (
	"github.com/sunimalherath/go-files/errorhandling"
	"os"
)

func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	errorhandling.ReportError(err)
}
