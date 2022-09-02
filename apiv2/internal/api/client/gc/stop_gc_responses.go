// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// StopGCReader is a Reader for the StopGC structure.
type StopGCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopGCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopGCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopGCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopGCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopGCNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopGCInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopGCOK creates a StopGCOK with default headers values
func NewStopGCOK() *StopGCOK {
	return &StopGCOK{}
}

/*StopGCOK handles this case with default header values.

Success
*/
type StopGCOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *StopGCOK) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcOK ", 200)
}

func (o *StopGCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewStopGCUnauthorized creates a StopGCUnauthorized with default headers values
func NewStopGCUnauthorized() *StopGCUnauthorized {
	return &StopGCUnauthorized{}
}

/*StopGCUnauthorized handles this case with default header values.

Unauthorized
*/
type StopGCUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopGCUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcUnauthorized  %+v", 401, o.Payload)
}

func (o *StopGCUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopGCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopGCForbidden creates a StopGCForbidden with default headers values
func NewStopGCForbidden() *StopGCForbidden {
	return &StopGCForbidden{}
}

/*StopGCForbidden handles this case with default header values.

Forbidden
*/
type StopGCForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopGCForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcForbidden  %+v", 403, o.Payload)
}

func (o *StopGCForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopGCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopGCNotFound creates a StopGCNotFound with default headers values
func NewStopGCNotFound() *StopGCNotFound {
	return &StopGCNotFound{}
}

/*StopGCNotFound handles this case with default header values.

Not found
*/
type StopGCNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopGCNotFound) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcNotFound  %+v", 404, o.Payload)
}

func (o *StopGCNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopGCNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopGCInternalServerError creates a StopGCInternalServerError with default headers values
func NewStopGCInternalServerError() *StopGCInternalServerError {
	return &StopGCInternalServerError{}
}

/*StopGCInternalServerError handles this case with default header values.

Internal server error
*/
type StopGCInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopGCInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcInternalServerError  %+v", 500, o.Payload)
}

func (o *StopGCInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopGCInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
