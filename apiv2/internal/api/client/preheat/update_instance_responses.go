// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// UpdateInstanceReader is a Reader for the UpdateInstance structure.
type UpdateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInstanceOK creates a UpdateInstanceOK with default headers values
func NewUpdateInstanceOK() *UpdateInstanceOK {
	return &UpdateInstanceOK{}
}

/*UpdateInstanceOK handles this case with default header values.

Success
*/
type UpdateInstanceOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /p2p/preheat/instances/{preheat_instance_name}][%d] updateInstanceOK ", 200)
}

func (o *UpdateInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateInstanceBadRequest creates a UpdateInstanceBadRequest with default headers values
func NewUpdateInstanceBadRequest() *UpdateInstanceBadRequest {
	return &UpdateInstanceBadRequest{}
}

/*UpdateInstanceBadRequest handles this case with default header values.

Bad request
*/
type UpdateInstanceBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateInstanceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /p2p/preheat/instances/{preheat_instance_name}][%d] updateInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInstanceBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceUnauthorized creates a UpdateInstanceUnauthorized with default headers values
func NewUpdateInstanceUnauthorized() *UpdateInstanceUnauthorized {
	return &UpdateInstanceUnauthorized{}
}

/*UpdateInstanceUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateInstanceUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /p2p/preheat/instances/{preheat_instance_name}][%d] updateInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateInstanceUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceForbidden creates a UpdateInstanceForbidden with default headers values
func NewUpdateInstanceForbidden() *UpdateInstanceForbidden {
	return &UpdateInstanceForbidden{}
}

/*UpdateInstanceForbidden handles this case with default header values.

Forbidden
*/
type UpdateInstanceForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateInstanceForbidden) Error() string {
	return fmt.Sprintf("[PUT /p2p/preheat/instances/{preheat_instance_name}][%d] updateInstanceForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInstanceForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceNotFound creates a UpdateInstanceNotFound with default headers values
func NewUpdateInstanceNotFound() *UpdateInstanceNotFound {
	return &UpdateInstanceNotFound{}
}

/*UpdateInstanceNotFound handles this case with default header values.

Not found
*/
type UpdateInstanceNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateInstanceNotFound) Error() string {
	return fmt.Sprintf("[PUT /p2p/preheat/instances/{preheat_instance_name}][%d] updateInstanceNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInstanceNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceInternalServerError creates a UpdateInstanceInternalServerError with default headers values
func NewUpdateInstanceInternalServerError() *UpdateInstanceInternalServerError {
	return &UpdateInstanceInternalServerError{}
}

/*UpdateInstanceInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateInstanceInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /p2p/preheat/instances/{preheat_instance_name}][%d] updateInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInstanceInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
