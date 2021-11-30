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

// GetReplicationPolicyReader is a Reader for the GetReplicationPolicy structure.
type GetReplicationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReplicationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationPolicyOK creates a GetReplicationPolicyOK with default headers values
func NewGetReplicationPolicyOK() *GetReplicationPolicyOK {
	return &GetReplicationPolicyOK{}
}

/*GetReplicationPolicyOK handles this case with default header values.

Success
*/
type GetReplicationPolicyOK struct {
	Payload *model.ReplicationPolicy
}

func (o *GetReplicationPolicyOK) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyOK  %+v", 200, o.Payload)
}

func (o *GetReplicationPolicyOK) GetPayload() *model.ReplicationPolicy {
	return o.Payload
}

func (o *GetReplicationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ReplicationPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationPolicyUnauthorized creates a GetReplicationPolicyUnauthorized with default headers values
func NewGetReplicationPolicyUnauthorized() *GetReplicationPolicyUnauthorized {
	return &GetReplicationPolicyUnauthorized{}
}

/*GetReplicationPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type GetReplicationPolicyUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReplicationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReplicationPolicyUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReplicationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationPolicyForbidden creates a GetReplicationPolicyForbidden with default headers values
func NewGetReplicationPolicyForbidden() *GetReplicationPolicyForbidden {
	return &GetReplicationPolicyForbidden{}
}

/*GetReplicationPolicyForbidden handles this case with default header values.

Forbidden
*/
type GetReplicationPolicyForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReplicationPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetReplicationPolicyForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReplicationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationPolicyInternalServerError creates a GetReplicationPolicyInternalServerError with default headers values
func NewGetReplicationPolicyInternalServerError() *GetReplicationPolicyInternalServerError {
	return &GetReplicationPolicyInternalServerError{}
}

/*GetReplicationPolicyInternalServerError handles this case with default header values.

Internal server error
*/
type GetReplicationPolicyInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReplicationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReplicationPolicyInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReplicationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
