package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rook/rook/pkg/api/gen/models"
)

/*GetVolumeAttachmentOK OK

swagger:response getVolumeAttachmentOK
*/
type GetVolumeAttachmentOK struct {

	// In: body
	Payload *models.Attachment `json:"body,omitempty"`
}

// NewGetVolumeAttachmentOK creates GetVolumeAttachmentOK with default headers values
func NewGetVolumeAttachmentOK() *GetVolumeAttachmentOK {
	return &GetVolumeAttachmentOK{}
}

// WithPayload adds the payload to the get volume attachment o k response
func (o *GetVolumeAttachmentOK) WithPayload(payload *models.Attachment) *GetVolumeAttachmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get volume attachment o k response
func (o *GetVolumeAttachmentOK) SetPayload(payload *models.Attachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVolumeAttachmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetVolumeAttachmentDefault error

swagger:response getVolumeAttachmentDefault
*/
type GetVolumeAttachmentDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVolumeAttachmentDefault creates GetVolumeAttachmentDefault with default headers values
func NewGetVolumeAttachmentDefault(code int) *GetVolumeAttachmentDefault {
	if code <= 0 {
		code = 500
	}

	return &GetVolumeAttachmentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get volume attachment default response
func (o *GetVolumeAttachmentDefault) WithStatusCode(code int) *GetVolumeAttachmentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get volume attachment default response
func (o *GetVolumeAttachmentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get volume attachment default response
func (o *GetVolumeAttachmentDefault) WithPayload(payload *models.Error) *GetVolumeAttachmentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get volume attachment default response
func (o *GetVolumeAttachmentDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVolumeAttachmentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
