package filebasics

import (
	"github.com/sunimalherath/go-files/errorhandling"
	"os"
)

func OpenAndCloseFile(fileName string) {
	// Open file with os.Open()
	file, err := os.Open(fileName)
	errorhandling.ReportError(err)
	GetFileInfo(file.Name())
	file.Close()

	// Open file with os.OpenFile()
	file2, err2 := os.OpenFile(fileName, os.O_APPEND, 0666)
	errorhandling.ReportError(err2)
	GetFileInfo(file2.Name())
	file2.Close()

	//	os.O_RDONLY		=> READ ONLY
	//	os.O_WRONLY		=> WRITE ONLY
	//	os.O_RDWR		=> READ WRITE
	//	os.O_APPEND		=> APPEND TO END OF THE FILE
	//	os.O_CREATE		=> CREATE IF DON'T EXIST
	//	os.O_TRUNC		=> TRUNCATE WHEN OPENING THE FILE
}
