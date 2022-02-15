package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

//TEMPLATE_DIR Contains path of tepmpate dir
const TEMPLATE_DIR string = "Templates"

//templatesCache contains cache of executed and parsed tmplates
var templatesCache map[string]*template.Template

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

//GenerateTemplate generates html Output
func (rq *Request) GenerateTemplate(name string, values interface{}) error {
	dir := TEMPLATE_DIR + "/" + name

	tmpl, exists := templatesCache[dir]
	if !exists {
		templatesCache[dir] = template.Must(template.ParseFiles(dir))
	}

	tmpl = templatesCache[dir]

	buffer := new(bytes.Buffer)

	err := tmpl.Execute(buffer, values)

	if err != nil {
		return fmt.Errorf("%s: %s", "Unable to execute and run template with err ", err)
	}

	_, err = fmt.Fprintf(rq.out, "%v", buffer)
	return err
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
