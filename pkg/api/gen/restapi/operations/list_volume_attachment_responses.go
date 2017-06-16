package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rook/rook/pkg/api/gen/models"
)

/*ListVolumeAttachmentOK list of all attachments for the volume.

swagger:response listVolumeAttachmentOK
*/
type ListVolumeAttachmentOK struct {

	// In: body
	Payload []*models.Attachment `json:"body,omitempty"`
}

// NewListVolumeAttachmentOK creates ListVolumeAttachmentOK with default headers values
func NewListVolumeAttachmentOK() *ListVolumeAttachmentOK {
	return &ListVolumeAttachmentOK{}
}

// WithPayload adds the payload to the list volume attachment o k response
func (o *ListVolumeAttachmentOK) WithPayload(payload []*models.Attachment) *ListVolumeAttachmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list volume attachment o k response
func (o *ListVolumeAttachmentOK) SetPayload(payload []*models.Attachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVolumeAttachmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ListVolumeAttachmentDefault generic error response

swagger:response listVolumeAttachmentDefault
*/
type ListVolumeAttachmentDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewListVolumeAttachmentDefault creates ListVolumeAttachmentDefault with default headers values
func NewListVolumeAttachmentDefault(code int) *ListVolumeAttachmentDefault {
	if code <= 0 {
		code = 500
	}

	return &ListVolumeAttachmentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list volume attachment default response
func (o *ListVolumeAttachmentDefault) WithStatusCode(code int) *ListVolumeAttachmentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list volume attachment default response
func (o *ListVolumeAttachmentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list volume attachment default response
func (o *ListVolumeAttachmentDefault) WithPayload(payload *models.Error) *ListVolumeAttachmentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list volume attachment default response
func (o *ListVolumeAttachmentDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVolumeAttachmentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
