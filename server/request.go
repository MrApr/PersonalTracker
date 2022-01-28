package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Response defines response of app
type Response map[string]interface{}

//Request Defines request struct for handlers
type Request struct {
	*route
	out     http.ResponseWriter
	ctx     context.Context
	Request *http.Request
	status  int
}

//Json generates json output
func (rq *Request) Json(resp *Response) error {
	body, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	if rq.status == 0 {
		rq.status = 200
	}

	rq.out.Header().Set("Content-Type", "application/json")
	rq.out.WriteHeader(rq.status)

	_, err = fmt.Fprintf(rq.out, "%s", body)
	if err != nil {
		return err
	}
	return nil
}

//ParseBody parses body and returns it
func (rq *Request) ParseBody(values interface{}) error {
	body, err := ioutil.ReadAll(rq.Request.Body)
	if err != nil {
		return fmt.Errorf("%s, %s", "Unable to parse json body with error", err)
	}
	err = json.Unmarshal(body, values)

	if err != nil {
		return fmt.Errorf("%s, %s", "Unable to parse json body with error", err)
	}
	return nil
}

//Status for executed user request
func (rq *Request) Status(status int) *Request {
	rq.status = status
	return rq
}
