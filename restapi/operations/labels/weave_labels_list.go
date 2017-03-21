package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveLabelsListHandlerFunc turns a function with the right signature into a weave labels list handler
type WeaveLabelsListHandlerFunc func(WeaveLabelsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveLabelsListHandlerFunc) Handle(params WeaveLabelsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveLabelsListHandler interface for that can handle valid weave labels list params
type WeaveLabelsListHandler interface {
	Handle(WeaveLabelsListParams, interface{}) middleware.Responder
}

// NewWeaveLabelsList creates a new http.Handler for the weave labels list operation
func NewWeaveLabelsList(ctx *middleware.Context, handler WeaveLabelsListHandler) *WeaveLabelsList {
	return &WeaveLabelsList{Context: ctx, Handler: handler}
}

/*WeaveLabelsList swagger:route GET /places/{placeId}/labels labels weaveLabelsList

Lists labels for a place.

*/
type WeaveLabelsList struct {
	Context *middleware.Context
	Handler WeaveLabelsListHandler
}

func (o *WeaveLabelsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveLabelsListParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}