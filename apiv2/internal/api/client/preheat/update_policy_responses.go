// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// UpdatePolicyReader is a Reader for the UpdatePolicy structure.
type UpdatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdatePolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePolicyOK creates a UpdatePolicyOK with default headers values
func NewUpdatePolicyOK() *UpdatePolicyOK {
	return &UpdatePolicyOK{}
}

/*UpdatePolicyOK handles this case with default header values.

Success
*/
type UpdatePolicyOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdatePolicyOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] updatePolicyOK ", 200)
}

func (o *UpdatePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdatePolicyBadRequest creates a UpdatePolicyBadRequest with default headers values
func NewUpdatePolicyBadRequest() *UpdatePolicyBadRequest {
	return &UpdatePolicyBadRequest{}
}

/*UpdatePolicyBadRequest handles this case with default header values.

Bad request
*/
type UpdatePolicyBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdatePolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] updatePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePolicyBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdatePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyUnauthorized creates a UpdatePolicyUnauthorized with default headers values
func NewUpdatePolicyUnauthorized() *UpdatePolicyUnauthorized {
	return &UpdatePolicyUnauthorized{}
}

/*UpdatePolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdatePolicyUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdatePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] updatePolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdatePolicyUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdatePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyForbidden creates a UpdatePolicyForbidden with default headers values
func NewUpdatePolicyForbidden() *UpdatePolicyForbidden {
	return &UpdatePolicyForbidden{}
}

/*UpdatePolicyForbidden handles this case with default header values.

Forbidden
*/
type UpdatePolicyForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdatePolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] updatePolicyForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePolicyForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdatePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyNotFound creates a UpdatePolicyNotFound with default headers values
func NewUpdatePolicyNotFound() *UpdatePolicyNotFound {
	return &UpdatePolicyNotFound{}
}

/*UpdatePolicyNotFound handles this case with default header values.

Not found
*/
type UpdatePolicyNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdatePolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] updatePolicyNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePolicyNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdatePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyConflict creates a UpdatePolicyConflict with default headers values
func NewUpdatePolicyConflict() *UpdatePolicyConflict {
	return &UpdatePolicyConflict{}
}

/*UpdatePolicyConflict handles this case with default header values.

Conflict
*/
type UpdatePolicyConflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdatePolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] updatePolicyConflict  %+v", 409, o.Payload)
}

func (o *UpdatePolicyConflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdatePolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyInternalServerError creates a UpdatePolicyInternalServerError with default headers values
func NewUpdatePolicyInternalServerError() *UpdatePolicyInternalServerError {
	return &UpdatePolicyInternalServerError{}
}

/*UpdatePolicyInternalServerError handles this case with default header values.

Internal server error
*/
type UpdatePolicyInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdatePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] updatePolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdatePolicyInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdatePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
