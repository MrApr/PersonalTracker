package config

import "github.com/MrApr/PersonalTracker/config/internal"

//Reader and interface for config readers
type Reader interface {
	//Load file that contains configuration
	Load(dir string)
	//Get data from configuration files
	Get(key string) interface{}
}

//CreateNewReader for reading operations
func CreateNewReader(name string) Reader {
	switch name {
	case "env":
		return new(internal.Memory)
	default:
		return new(internal.Env)
	}
}
