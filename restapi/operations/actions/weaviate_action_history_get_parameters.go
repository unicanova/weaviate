/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

package actions

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateActionHistoryGetParams creates a new WeaviateActionHistoryGetParams object
// with the default values initialized.
func NewWeaviateActionHistoryGetParams() WeaviateActionHistoryGetParams {
	var ()
	return WeaviateActionHistoryGetParams{}
}

// WeaviateActionHistoryGetParams contains all the bound params for the weaviate action history get operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.action.history.get
type WeaviateActionHistoryGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Unique ID of the action.
	  Required: true
	  In: path
	*/
	ActionID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaviateActionHistoryGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rActionID, rhkActionID, _ := route.Params.GetOK("actionId")
	if err := o.bindActionID(rActionID, rhkActionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateActionHistoryGetParams) bindActionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("actionId", "path", "strfmt.UUID", raw)
	}
	o.ActionID = *(value.(*strfmt.UUID))

	return nil
}
