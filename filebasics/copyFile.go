package filebasics

import (
	"github.com/sunimalherath/go-files/errorhandling"
	"io"
	"log"
	"os"
)

func CopyFile(existingFile, newFile string) {
	// Open the existing file
	ef, err := os.Open(existingFile)
	errorhandling.ReportError(err)
	defer ef.Close()

	// Create new file
	nf, err2 := os.Create(newFile)
	errorhandling.ReportError(err2)
	defer nf.Close()

	// Copy from the existing file to the new file
	bytesWritten, err3 := io.Copy(nf, ef)
	errorhandling.ReportError(err3)
	log.Printf("Bytes Written: %v", bytesWritten)

	// Commit changes
	err4 := nf.Sync()
	errorhandling.ReportError(err4)
}
