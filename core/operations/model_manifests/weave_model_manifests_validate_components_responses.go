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
 package model_manifests




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/core/models"
)

/*WeaveModelManifestsValidateComponentsOK Successful response

swagger:response weaveModelManifestsValidateComponentsOK
*/
type WeaveModelManifestsValidateComponentsOK struct {

	// In: body
	Payload *models.ModelManifestsValidateComponentsResponse `json:"body,omitempty"`
}

// NewWeaveModelManifestsValidateComponentsOK creates WeaveModelManifestsValidateComponentsOK with default headers values
func NewWeaveModelManifestsValidateComponentsOK() *WeaveModelManifestsValidateComponentsOK {
	return &WeaveModelManifestsValidateComponentsOK{}
}

// WithPayload adds the payload to the weave model manifests validate components o k response
func (o *WeaveModelManifestsValidateComponentsOK) WithPayload(payload *models.ModelManifestsValidateComponentsResponse) *WeaveModelManifestsValidateComponentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave model manifests validate components o k response
func (o *WeaveModelManifestsValidateComponentsOK) SetPayload(payload *models.ModelManifestsValidateComponentsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveModelManifestsValidateComponentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}