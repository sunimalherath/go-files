package filebasics

import (
	"github.com/sunimalherath/go-files/errorhandling"
	"log"
	"os"
	"time"
)

func ChangePermission(fileName string) {
	err := os.Chmod(fileName, 0777)
	errorhandling.ReportError(err)
	log.Println("File permission changed to - read, write, & execute for owner, group and others")
}

func ChangeOwnership(fileName string) {
	err := os.Chown(fileName, os.Getuid(), os.Getgid())
	errorhandling.ReportError(err)
	log.Println("File ownership changed")
}

func ChangeTimestamp(fileName string) {
	GetFileInfo(fileName)
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow

	err := os.Chtimes(fileName, lastAccessTime, lastModifyTime)
	errorhandling.ReportError(err)
	GetFileInfo(fileName)
}
