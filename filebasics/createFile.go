package filebasics

import (
	"fmt"

	"os"

	"github.com/sunimalherath/go-files/errorhandling"
)

func CreateFile(fileName string) {
	newFile, err := os.Create(fileName)
	errorhandling.ReportError(err)

	defer newFile.Close()
	fmt.Println(newFile)
}
