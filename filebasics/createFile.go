package filebasics

import (
	"fmt"
	"github.com/files/errorhandling"
	"os"
)

func CreateFile(fileName string) {
	newFile, err := os.Create(fileName)
	errorhandling.ReportError(err)

	defer newFile.Close()
	fmt.Println(newFile)
}
