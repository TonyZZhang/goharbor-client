// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// DeleteImmuRuleReader is a Reader for the DeleteImmuRule structure.
type DeleteImmuRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImmuRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteImmuRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteImmuRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteImmuRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteImmuRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteImmuRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteImmuRuleOK creates a DeleteImmuRuleOK with default headers values
func NewDeleteImmuRuleOK() *DeleteImmuRuleOK {
	return &DeleteImmuRuleOK{}
}

/*DeleteImmuRuleOK handles this case with default header values.

Success
*/
type DeleteImmuRuleOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteImmuRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] deleteImmuRuleOK ", 200)
}

func (o *DeleteImmuRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteImmuRuleBadRequest creates a DeleteImmuRuleBadRequest with default headers values
func NewDeleteImmuRuleBadRequest() *DeleteImmuRuleBadRequest {
	return &DeleteImmuRuleBadRequest{}
}

/*DeleteImmuRuleBadRequest handles this case with default header values.

Bad request
*/
type DeleteImmuRuleBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteImmuRuleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] deleteImmuRuleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteImmuRuleBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteImmuRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImmuRuleUnauthorized creates a DeleteImmuRuleUnauthorized with default headers values
func NewDeleteImmuRuleUnauthorized() *DeleteImmuRuleUnauthorized {
	return &DeleteImmuRuleUnauthorized{}
}

/*DeleteImmuRuleUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteImmuRuleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteImmuRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] deleteImmuRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteImmuRuleUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteImmuRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImmuRuleForbidden creates a DeleteImmuRuleForbidden with default headers values
func NewDeleteImmuRuleForbidden() *DeleteImmuRuleForbidden {
	return &DeleteImmuRuleForbidden{}
}

/*DeleteImmuRuleForbidden handles this case with default header values.

Forbidden
*/
type DeleteImmuRuleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteImmuRuleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] deleteImmuRuleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteImmuRuleForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteImmuRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImmuRuleInternalServerError creates a DeleteImmuRuleInternalServerError with default headers values
func NewDeleteImmuRuleInternalServerError() *DeleteImmuRuleInternalServerError {
	return &DeleteImmuRuleInternalServerError{}
}

/*DeleteImmuRuleInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteImmuRuleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteImmuRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] deleteImmuRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteImmuRuleInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteImmuRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
