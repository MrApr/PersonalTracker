package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Env Defines path that contains environment variable
type Env string

//Load file that contains configuration
func (e *Env) Load(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		panic(fmt.Errorf("%s: %s", "Unable to load .env file with error", err))
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		text := reader.Text()
		if err = e.registerVars(text); err != nil {
			panic(fmt.Errorf("%s: %s", "Unable to register .env variable with error", err))
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
