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
 package commands




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/core/models"
)

/*WeaveCommandsGetQueueOK Successful response

swagger:response weaveCommandsGetQueueOK
*/
type WeaveCommandsGetQueueOK struct {

	// In: body
	Payload *models.CommandsQueueResponse `json:"body,omitempty"`
}

// NewWeaveCommandsGetQueueOK creates WeaveCommandsGetQueueOK with default headers values
func NewWeaveCommandsGetQueueOK() *WeaveCommandsGetQueueOK {
	return &WeaveCommandsGetQueueOK{}
}

// WithPayload adds the payload to the weave commands get queue o k response
func (o *WeaveCommandsGetQueueOK) WithPayload(payload *models.CommandsQueueResponse) *WeaveCommandsGetQueueOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave commands get queue o k response
func (o *WeaveCommandsGetQueueOK) SetPayload(payload *models.CommandsQueueResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveCommandsGetQueueOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}