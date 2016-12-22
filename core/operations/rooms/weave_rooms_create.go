/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package rooms


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveRoomsCreateHandlerFunc turns a function with the right signature into a weave rooms create handler
type WeaveRoomsCreateHandlerFunc func(WeaveRoomsCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveRoomsCreateHandlerFunc) Handle(params WeaveRoomsCreateParams) middleware.Responder {
	return fn(params)
}

// WeaveRoomsCreateHandler interface for that can handle valid weave rooms create params
type WeaveRoomsCreateHandler interface {
	Handle(WeaveRoomsCreateParams) middleware.Responder
}

// NewWeaveRoomsCreate creates a new http.Handler for the weave rooms create operation
func NewWeaveRoomsCreate(ctx *middleware.Context, handler WeaveRoomsCreateHandler) *WeaveRoomsCreate {
	return &WeaveRoomsCreate{Context: ctx, Handler: handler}
}

/*WeaveRoomsCreate swagger:route POST /places/{placeId}/rooms/create rooms weaveRoomsCreate

Creates a new room.

*/
type WeaveRoomsCreate struct {
	Context *middleware.Context
	Handler WeaveRoomsCreateHandler
}

func (o *WeaveRoomsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveRoomsCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}