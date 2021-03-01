// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHealthOKCode is the HTTP code returned for type GetHealthOK
const GetHealthOKCode int = 200

/*GetHealthOK Health check

swagger:response getHealthOK
*/
type GetHealthOK struct {
}

// NewGetHealthOK creates GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {

	return &GetHealthOK{}
}

// WriteResponse to the client
func (o *GetHealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
