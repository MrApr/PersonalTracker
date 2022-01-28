package server

import (
	"context"
	"net/http"
)

//AprHandler defines custom handler for endpoints
type AprHandler func(request *Request) error

//handleUserRequest executes user request
func handleUserRequest(outPut *http.ResponseWriter, reqRoute *route, request *http.Request) {
	userRequest := new(Request)
	userRequest.out = *outPut
	userRequest.route = reqRoute
	userRequest.Request = request
	userRequest.ctx = context.Background()

	err := reqRoute.Service(userRequest)
	if err != nil {
		userRequest.ctx.Done()
		panic(err)
	}
}
