package server

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//route defines a struct for routing purpose
type route struct {
	Path    string
	Method  string
	Service AprHandler
}

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
	Routes          []route
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

	runtime.GOMAXPROCS(cfg.MaxCpuNums)

	httpSv := http.Server{
		Addr:         address,
		WriteTimeout: cfg.MaxWriteTimeOut,
		ReadTimeout:  cfg.MaxReadTimeOut,
		Handler:      cfg,
	}

	fmt.Println("Starting new server instance")
	err := httpSv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//Post registers a route with this method
func (cfg *config) Post(path string, handler AprHandler) {
	cfg.addRoute("post", path, handler)
}

//Get registers a route with this method
func (cfg *config) Get(path string, handler AprHandler) {
	cfg.addRoute("get", path, handler)
}

//addRoute internal interactor with routes container
func (cfg *config) addRoute(method string, path string, handler AprHandler) {
	method = strings.ToLower(method)
	path = strings.ToLower(path)

	if cfg.IsRouteExists(method, path) {
		panic(fmt.Errorf("%s", "Multiple routes with same signature"))
	}

	router := route{
		Path:    path,
		Method:  method,
		Service: handler,
	}

	cfg.Routes = append(cfg.Routes, router)
}

//IsRouteExists checks whether a route exists or not
func (cfg *config) IsRouteExists(method string, path string) bool {
	for _, val := range cfg.Routes {
		if val.Path == path && val.Method == method {
			return true
		}
	}
	return false
}

//Starting and getting http requests
func (cfg *config) ServeHTTP(outPut http.ResponseWriter, request *http.Request) {
	var method, path string
	method = strings.ToLower(request.Method)
	path = strings.ToLower(request.URL.Path)

	reqRoute := cfg.matchRoute(method, path)

	if reqRoute == nil {
		outPut.Header().Set("Content-Type", "application/json")
		outPut.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(outPut, "%s", "page not found")
		return
	}
	handleUserRequest(&outPut, reqRoute, request)
}

//Find matching route
func (cfg *config) matchRoute(method string, path string) *route {
	for _, val := range cfg.Routes {
		if val.Path == path && val.Method == method {
			return &val
		}
	}
	return nil
}
