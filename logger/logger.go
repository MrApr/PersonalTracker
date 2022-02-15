package logger

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

//Create a new custom logger
func Create(lgType string) *log.Logger {
	switch lgType {
	case "file":
		file := createLocalFile()
		return log.New(file, "custom logger", log.Lmicroseconds|log.Llongfile|log.Ltime)
	default:
		return log.New(os.Stdout, "custom logger", log.Lmicroseconds|log.Llongfile|log.Ltime)
	}
}

//createLocalFile and return it
func createLocalFile() *os.File {
	fileName := "Logger" + strconv.Itoa(time.Now().Day())

	if !exists(fileName) {
		createFileOnDisk(fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func createFileOnDisk(fileName string) {
	_, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
}
