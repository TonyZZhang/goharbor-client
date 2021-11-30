// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// CreateReplicationPolicyReader is a Reader for the CreateReplicationPolicy structure.
type CreateReplicationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReplicationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateReplicationPolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReplicationPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateReplicationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateReplicationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateReplicationPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateReplicationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateReplicationPolicyCreated creates a CreateReplicationPolicyCreated with default headers values
func NewCreateReplicationPolicyCreated() *CreateReplicationPolicyCreated {
	return &CreateReplicationPolicyCreated{}
}

/*CreateReplicationPolicyCreated handles this case with default header values.

Created
*/
type CreateReplicationPolicyCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateReplicationPolicyCreated) Error() string {
	return fmt.Sprintf("[POST /replication/policies][%d] createReplicationPolicyCreated ", 201)
}

func (o *CreateReplicationPolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCreateReplicationPolicyBadRequest creates a CreateReplicationPolicyBadRequest with default headers values
func NewCreateReplicationPolicyBadRequest() *CreateReplicationPolicyBadRequest {
	return &CreateReplicationPolicyBadRequest{}
}

/*CreateReplicationPolicyBadRequest handles this case with default header values.

Bad request
*/
type CreateReplicationPolicyBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateReplicationPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /replication/policies][%d] createReplicationPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateReplicationPolicyBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateReplicationPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReplicationPolicyUnauthorized creates a CreateReplicationPolicyUnauthorized with default headers values
func NewCreateReplicationPolicyUnauthorized() *CreateReplicationPolicyUnauthorized {
	return &CreateReplicationPolicyUnauthorized{}
}

/*CreateReplicationPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateReplicationPolicyUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateReplicationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /replication/policies][%d] createReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateReplicationPolicyUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateReplicationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReplicationPolicyForbidden creates a CreateReplicationPolicyForbidden with default headers values
func NewCreateReplicationPolicyForbidden() *CreateReplicationPolicyForbidden {
	return &CreateReplicationPolicyForbidden{}
}

/*CreateReplicationPolicyForbidden handles this case with default header values.

Forbidden
*/
type CreateReplicationPolicyForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateReplicationPolicyForbidden) Error() string {
	return fmt.Sprintf("[POST /replication/policies][%d] createReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *CreateReplicationPolicyForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateReplicationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReplicationPolicyConflict creates a CreateReplicationPolicyConflict with default headers values
func NewCreateReplicationPolicyConflict() *CreateReplicationPolicyConflict {
	return &CreateReplicationPolicyConflict{}
}

/*CreateReplicationPolicyConflict handles this case with default header values.

Conflict
*/
type CreateReplicationPolicyConflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateReplicationPolicyConflict) Error() string {
	return fmt.Sprintf("[POST /replication/policies][%d] createReplicationPolicyConflict  %+v", 409, o.Payload)
}

func (o *CreateReplicationPolicyConflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateReplicationPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReplicationPolicyInternalServerError creates a CreateReplicationPolicyInternalServerError with default headers values
func NewCreateReplicationPolicyInternalServerError() *CreateReplicationPolicyInternalServerError {
	return &CreateReplicationPolicyInternalServerError{}
}

/*CreateReplicationPolicyInternalServerError handles this case with default header values.

Internal server error
*/
type CreateReplicationPolicyInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateReplicationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /replication/policies][%d] createReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateReplicationPolicyInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateReplicationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
