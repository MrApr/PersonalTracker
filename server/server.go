package server

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

//define and declare package constants
const (
	DefaultHost         string        = "localhost"
	DefaultPort         int           = 8000
	DefaultReadTimeOut  time.Duration = 60
	DefaultWriteTimeOut time.Duration = 60
)

//config defines server configurations
type config struct {
	Port            int
	Host            string
	MaxReadTimeOut  time.Duration
	MaxWriteTimeOut time.Duration
	MaxCpuNums      int
}

//ConfigureServer StartServer makes configuration and starts server
func ConfigureServer(host string, port int) *config {
	cfg := new(config)

	if host == "" {
		host = DefaultHost
	}
	if port == 0 {
		port = DefaultPort
	}

	cfg.Port = port
	cfg.Host = host

	cfg.MaxCpuNums = runtime.NumCPU()
	cfg.MaxReadTimeOut = DefaultReadTimeOut * time.Second
	cfg.MaxWriteTimeOut = DefaultWriteTimeOut * time.Second
	return cfg
}

func (cfg *config) StartServer() {
	var address string = cfg.Host + ":" + strconv.Itoa(cfg.Port)

	httpSv := http.Server{
		Addr:         address,
		WriteTimeout: cfg.MaxWriteTimeOut,
		ReadTimeout:  cfg.MaxReadTimeOut,
	}

	fmt.Println("Starting new server instance")
	err := httpSv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
