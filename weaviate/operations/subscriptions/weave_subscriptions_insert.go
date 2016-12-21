package subscriptions


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveSubscriptionsInsertHandlerFunc turns a function with the right signature into a weave subscriptions insert handler
type WeaveSubscriptionsInsertHandlerFunc func(WeaveSubscriptionsInsertParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveSubscriptionsInsertHandlerFunc) Handle(params WeaveSubscriptionsInsertParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveSubscriptionsInsertHandler interface for that can handle valid weave subscriptions insert params
type WeaveSubscriptionsInsertHandler interface {
	Handle(WeaveSubscriptionsInsertParams, interface{}) middleware.Responder
}

// NewWeaveSubscriptionsInsert creates a new http.Handler for the weave subscriptions insert operation
func NewWeaveSubscriptionsInsert(ctx *middleware.Context, handler WeaveSubscriptionsInsertHandler) *WeaveSubscriptionsInsert {
	return &WeaveSubscriptionsInsert{Context: ctx, Handler: handler}
}

/*WeaveSubscriptionsInsert swagger:route POST /subscriptions subscriptions weaveSubscriptionsInsert

Inserts a new subscription.

*/
type WeaveSubscriptionsInsert struct {
	Context *middleware.Context
	Handler WeaveSubscriptionsInsertHandler
}

func (o *WeaveSubscriptionsInsert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveSubscriptionsInsertParams()

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