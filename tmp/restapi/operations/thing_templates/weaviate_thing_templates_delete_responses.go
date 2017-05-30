package thing_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateThingTemplatesDeleteNoContentCode is the HTTP code returned for type WeaviateThingTemplatesDeleteNoContent
const WeaviateThingTemplatesDeleteNoContentCode int = 204

/*WeaviateThingTemplatesDeleteNoContent Successful deleted.

swagger:response weaviateThingTemplatesDeleteNoContent
*/
type WeaviateThingTemplatesDeleteNoContent struct {
}

// NewWeaviateThingTemplatesDeleteNoContent creates WeaviateThingTemplatesDeleteNoContent with default headers values
func NewWeaviateThingTemplatesDeleteNoContent() *WeaviateThingTemplatesDeleteNoContent {
	return &WeaviateThingTemplatesDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateThingTemplatesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// WeaviateThingTemplatesDeleteUnauthorizedCode is the HTTP code returned for type WeaviateThingTemplatesDeleteUnauthorized
const WeaviateThingTemplatesDeleteUnauthorizedCode int = 401

/*WeaviateThingTemplatesDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingTemplatesDeleteUnauthorized
*/
type WeaviateThingTemplatesDeleteUnauthorized struct {
}

// NewWeaviateThingTemplatesDeleteUnauthorized creates WeaviateThingTemplatesDeleteUnauthorized with default headers values
func NewWeaviateThingTemplatesDeleteUnauthorized() *WeaviateThingTemplatesDeleteUnauthorized {
	return &WeaviateThingTemplatesDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingTemplatesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateThingTemplatesDeleteForbiddenCode is the HTTP code returned for type WeaviateThingTemplatesDeleteForbidden
const WeaviateThingTemplatesDeleteForbiddenCode int = 403

/*WeaviateThingTemplatesDeleteForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingTemplatesDeleteForbidden
*/
type WeaviateThingTemplatesDeleteForbidden struct {
}

// NewWeaviateThingTemplatesDeleteForbidden creates WeaviateThingTemplatesDeleteForbidden with default headers values
func NewWeaviateThingTemplatesDeleteForbidden() *WeaviateThingTemplatesDeleteForbidden {
	return &WeaviateThingTemplatesDeleteForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingTemplatesDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateThingTemplatesDeleteNotFoundCode is the HTTP code returned for type WeaviateThingTemplatesDeleteNotFound
const WeaviateThingTemplatesDeleteNotFoundCode int = 404

/*WeaviateThingTemplatesDeleteNotFound Successful query result but no resource was found.

swagger:response weaviateThingTemplatesDeleteNotFound
*/
type WeaviateThingTemplatesDeleteNotFound struct {
}

// NewWeaviateThingTemplatesDeleteNotFound creates WeaviateThingTemplatesDeleteNotFound with default headers values
func NewWeaviateThingTemplatesDeleteNotFound() *WeaviateThingTemplatesDeleteNotFound {
	return &WeaviateThingTemplatesDeleteNotFound{}
}

// WriteResponse to the client
func (o *WeaviateThingTemplatesDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateThingTemplatesDeleteNotImplementedCode is the HTTP code returned for type WeaviateThingTemplatesDeleteNotImplemented
const WeaviateThingTemplatesDeleteNotImplementedCode int = 501

/*WeaviateThingTemplatesDeleteNotImplemented Not (yet) implemented.

swagger:response weaviateThingTemplatesDeleteNotImplemented
*/
type WeaviateThingTemplatesDeleteNotImplemented struct {
}

// NewWeaviateThingTemplatesDeleteNotImplemented creates WeaviateThingTemplatesDeleteNotImplemented with default headers values
func NewWeaviateThingTemplatesDeleteNotImplemented() *WeaviateThingTemplatesDeleteNotImplemented {
	return &WeaviateThingTemplatesDeleteNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateThingTemplatesDeleteNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}