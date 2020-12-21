package filebasics

import (
	"log"
	"os"
)

func CheckFileRWPermission(fileName string) {
	// Check write permission
	log.Printf("Opening file with write only permission: %v", fileName)
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("File does not exists...")
		} else if os.IsPermission(err) {
			log.Fatalln("Write permission denied for the file")
		}
	}
	log.Println("File has write only permissions")
	file.Close()

	// Check read permission
	log.Printf("Opening file with read only permission: %v", fileName)
	file2, err2 := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err2 != nil {
		if os.IsNotExist(err2) {
			log.Fatalln("File does not exists...")
		} else if os.IsPermission(err2) {
			log.Fatalln("Read permission denied for the file")
		}
	}
	log.Println("File has read only permissions")
	file2.Close()
}
