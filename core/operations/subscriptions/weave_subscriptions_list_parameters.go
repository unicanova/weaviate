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
 package subscriptions




import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaveSubscriptionsListParams creates a new WeaveSubscriptionsListParams object
// with the default values initialized.
func NewWeaveSubscriptionsListParams() WeaveSubscriptionsListParams {
	var (
		altDefault         = string("json")
		prettyPrintDefault = bool(true)
	)
	return WeaveSubscriptionsListParams{
		Alt: &altDefault,

		PrettyPrint: &prettyPrintDefault,
	}
}

// WeaveSubscriptionsListParams contains all the bound params for the weave subscriptions list operation
// typically these are obtained from a http.Request
//
// swagger:parameters weave.subscriptions.list
type WeaveSubscriptionsListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Data format for the response.
	  In: query
	  Default: "json"
	*/
	Alt *string
	/*Channel type.
	  In: query
	*/
	ChannelType *string
	/*Selector specifying which fields to include in a partial response.
	  In: query
	*/
	Fields *string
	/*GCM registration id.
	  In: query
	*/
	GcmRegistrationID *string
	/*Specifies the language code that should be used for text values in the API response.
	  In: query
	*/
	Hl *string
	/*API key. Your API key identifies your project and provides you with API access, quota, and reports. Required unless you provide an OAuth 2.0 token.
	  In: query
	*/
	Key *string
	/*
	  In: query
	*/
	MaxResults *int64
	/*OAuth 2.0 token for the current user.
	  In: query
	*/
	OauthToken *string
	/*Returns response with indentations and line breaks.
	  In: query
	  Default: true
	*/
	PrettyPrint *bool
	/*Available to use for quota purposes for server-side applications. Can be any arbitrary string assigned to a user, but should not exceed 40 characters. Overrides userIp if both are provided.
	  In: query
	*/
	QuotaUser *string
	/*
	  In: query
	*/
	StartIndex *int64
	/*
	  In: query
	*/
	Token *string
	/*IP address of the site where the request originates. Use this if you want to enforce per-user limits.
	  In: query
	*/
	UserIP *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaveSubscriptionsListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAlt, qhkAlt, _ := qs.GetOK("alt")
	if err := o.bindAlt(qAlt, qhkAlt, route.Formats); err != nil {
		res = append(res, err)
	}

	qChannelType, qhkChannelType, _ := qs.GetOK("channelType")
	if err := o.bindChannelType(qChannelType, qhkChannelType, route.Formats); err != nil {
		res = append(res, err)
	}

	qFields, qhkFields, _ := qs.GetOK("fields")
	if err := o.bindFields(qFields, qhkFields, route.Formats); err != nil {
		res = append(res, err)
	}

	qGcmRegistrationID, qhkGcmRegistrationID, _ := qs.GetOK("gcmRegistrationId")
	if err := o.bindGcmRegistrationID(qGcmRegistrationID, qhkGcmRegistrationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qHl, qhkHl, _ := qs.GetOK("hl")
	if err := o.bindHl(qHl, qhkHl, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qMaxResults, qhkMaxResults, _ := qs.GetOK("maxResults")
	if err := o.bindMaxResults(qMaxResults, qhkMaxResults, route.Formats); err != nil {
		res = append(res, err)
	}

	qOauthToken, qhkOauthToken, _ := qs.GetOK("oauth_token")
	if err := o.bindOauthToken(qOauthToken, qhkOauthToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qPrettyPrint, qhkPrettyPrint, _ := qs.GetOK("prettyPrint")
	if err := o.bindPrettyPrint(qPrettyPrint, qhkPrettyPrint, route.Formats); err != nil {
		res = append(res, err)
	}

	qQuotaUser, qhkQuotaUser, _ := qs.GetOK("quotaUser")
	if err := o.bindQuotaUser(qQuotaUser, qhkQuotaUser, route.Formats); err != nil {
		res = append(res, err)
	}

	qStartIndex, qhkStartIndex, _ := qs.GetOK("startIndex")
	if err := o.bindStartIndex(qStartIndex, qhkStartIndex, route.Formats); err != nil {
		res = append(res, err)
	}

	qToken, qhkToken, _ := qs.GetOK("token")
	if err := o.bindToken(qToken, qhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserIP, qhkUserIP, _ := qs.GetOK("userIp")
	if err := o.bindUserIP(qUserIP, qhkUserIP, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaveSubscriptionsListParams) bindAlt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var altDefault string = string("json")
		o.Alt = &altDefault
		return nil
	}

	o.Alt = &raw

	if err := o.validateAlt(formats); err != nil {
		return err
	}

	return nil
}

func (o *WeaveSubscriptionsListParams) validateAlt(formats strfmt.Registry) error {

	if err := validate.Enum("alt", "query", *o.Alt, []interface{}{"json"}); err != nil {
		return err
	}

	return nil
}

func (o *WeaveSubscriptionsListParams) bindChannelType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ChannelType = &raw

	if err := o.validateChannelType(formats); err != nil {
		return err
	}

	return nil
}

func (o *WeaveSubscriptionsListParams) validateChannelType(formats strfmt.Registry) error {

	if err := validate.Enum("channelType", "query", *o.ChannelType, []interface{}{"gcm", "unknownChannelType"}); err != nil {
		return err
	}

	return nil
}

func (o *WeaveSubscriptionsListParams) bindFields(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Fields = &raw

	return nil
}

func (o *WeaveSubscriptionsListParams) bindGcmRegistrationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.GcmRegistrationID = &raw

	return nil
}

func (o *WeaveSubscriptionsListParams) bindHl(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Hl = &raw

	return nil
}

func (o *WeaveSubscriptionsListParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Key = &raw

	return nil
}

func (o *WeaveSubscriptionsListParams) bindMaxResults(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("maxResults", "query", "int64", raw)
	}
	o.MaxResults = &value

	return nil
}

func (o *WeaveSubscriptionsListParams) bindOauthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.OauthToken = &raw

	return nil
}

func (o *WeaveSubscriptionsListParams) bindPrettyPrint(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var prettyPrintDefault bool = bool(true)
		o.PrettyPrint = &prettyPrintDefault
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("prettyPrint", "query", "bool", raw)
	}
	o.PrettyPrint = &value

	return nil
}

func (o *WeaveSubscriptionsListParams) bindQuotaUser(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.QuotaUser = &raw

	return nil
}

func (o *WeaveSubscriptionsListParams) bindStartIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("startIndex", "query", "int64", raw)
	}
	o.StartIndex = &value

	return nil
}

func (o *WeaveSubscriptionsListParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Token = &raw

	return nil
}

func (o *WeaveSubscriptionsListParams) bindUserIP(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UserIP = &raw

	return nil
}