package internal

import (
	"bufio"
	"github.com/MrApr/PersonalTracker/Error"
	"os"
	"strings"
)

//Env Defines path that contains environment variable
type Env string

//Load file that contains configuration
func (e *Env) Load(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		panic(Error.AdvanceError{
			Message: err.Error(),
			File:    "EnvReader",
			Type:    "Critical",
			Line:    17,
		})
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		text := reader.Text()
		if err = e.registerVars(text); err != nil {
			panic(Error.AdvanceError{
				Message: err.Error(),
				File:    "EnvReader",
				Type:    "Critical",
				Line:    30,
			})
		}
	}
}

//Get data from configuration files
func (e *Env) Get(key string) interface{} {
	return os.Getenv(key)
}

//registerVars extracts and registers passed values
func (e *Env) registerVars(text string) error {
	exploded := strings.Split(text, "=")
	return os.Setenv(exploded[0], exploded[1])
}
