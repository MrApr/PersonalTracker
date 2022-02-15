package logger

import (
	"log"
	"os"
)

//Create a new custom logger
func Create(lgType string) *log.Logger {
	switch lgType {
	default:
		return log.New(os.Stdout, "custom logger", log.Lmicroseconds|log.Llongfile|log.Ltime)
	}
}
