// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CheckHealthOKCode is the HTTP code returned for type CheckHealthOK
const CheckHealthOKCode int = 200

/*
CheckHealthOK OK message.

swagger:response checkHealthOK
*/
type CheckHealthOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCheckHealthOK creates CheckHealthOK with default headers values
func NewCheckHealthOK() *CheckHealthOK {

	return &CheckHealthOK{}
}

// WithPayload adds the payload to the check health o k response
func (o *CheckHealthOK) WithPayload(payload string) *CheckHealthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check health o k response
func (o *CheckHealthOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckHealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
