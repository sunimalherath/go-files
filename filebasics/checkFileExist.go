package filebasics

import (
	"log"
	"os"
)

func CheckFileExistance(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
	log.Println("File found: ")
	log.Println(fileInfo.Name())
}
