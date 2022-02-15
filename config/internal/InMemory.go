package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Memory custom type
type Memory map[string]interface{}

//Load values from memory
func (e *Memory) Load(dir string) {
	if dir == "" {
		panic(fmt.Errorf("%s", "Invalid directory for .memory"))
	}
	e.readValues(dir)
}

//Get and return values inside a map
func (e Memory) Get(key string) interface{} {
	val, ok := e[key]

	if !ok {
		return nil
	}

	return val
}

//readValues reads and
func (e *Memory) readValues(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		panic(fmt.Errorf("%s: %s", "Config cannot get opened with err", err))
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		err = e.extractVars(reader.Text())
		if err != nil {
			panic(fmt.Errorf("%s %s", "Unable to extract string with error: ", err))
		}
	}
}

//extractVars inside an string and puts them inside array
func (e Memory) extractVars(key string) error {
	if !strings.Contains(key, "=") {
		return fmt.Errorf("%s", "Invalid key format")
	}
	exploded := strings.Split(key, "=")
	e[exploded[0]] = exploded[1]
	return nil
}
