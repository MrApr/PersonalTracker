package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/MrApr/PersonalTracker/Error"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

//TEMPLATE_DIR Contains path of tepmpate dir
const TEMPLATE_DIR string = "Templates"

//templatesCache contains cache of executed and parsed templates
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

//init initializes function
func init() {
	templatesCache = make(map[string]*template.Template)
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
	dir := makeTmplateDir(name)

	tmpl, exists := templatesCache[dir]
	if !exists {
		templatesCache[dir] = template.Must(template.ParseFiles(dir))
	}

	tmpl = templatesCache[dir]

	buffer := new(bytes.Buffer)

	err := tmpl.Execute(buffer, values)

	if err != nil {
		return &Error.AdvanceError{
			File:    "request",
			Type:    "Warning",
			Message: fmt.Sprintf("%s: %s", "Unable to execute and run template with err ", err.Error()),
			Line:    72,
		}
	}

	_, err = fmt.Fprintf(rq.out, "%v", buffer)
	return err
}

//ParseBody parses body and returns it
func (rq *Request) ParseBody(values interface{}) error {
	body, err := ioutil.ReadAll(rq.Request.Body)

	if len(body) == 0 {
		return nil
	}

	if err != nil {
		return &Error.AdvanceError{
			File:    "request",
			Type:    "Warning",
			Message: fmt.Sprintf("%s: %s", "Unable to parse json body with error", err.Error()),
			Line:    89,
		}
	}
	err = json.Unmarshal(body, values)

	if err != nil {
		return &Error.AdvanceError{
			File:    "request",
			Type:    "Warning",
			Message: fmt.Sprintf("%s: %s", "Unable to parse json body with error", err.Error()),
			Line:    98,
		}
	}
	return nil
}

//Status for executed user request
func (rq *Request) Status(status int) *Request {
	rq.status = status
	return rq
}

func makeTmplateDir(tmplName string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(&Error.AdvanceError{
			File:    "request",
			Type:    "Warning",
			Message: fmt.Sprintf("%s: %s", "cannot obtain current directory with error", err.Error()),
			Line:    118,
		})
	}
	return currentDir + "/" + TEMPLATE_DIR + "/" + tmplName
}
