// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// GetWebhookPolicyOfProjectReader is a Reader for the GetWebhookPolicyOfProject structure.
type GetWebhookPolicyOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebhookPolicyOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebhookPolicyOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebhookPolicyOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWebhookPolicyOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebhookPolicyOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebhookPolicyOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebhookPolicyOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebhookPolicyOfProjectOK creates a GetWebhookPolicyOfProjectOK with default headers values
func NewGetWebhookPolicyOfProjectOK() *GetWebhookPolicyOfProjectOK {
	return &GetWebhookPolicyOfProjectOK{}
}

/*GetWebhookPolicyOfProjectOK handles this case with default header values.

Get webhook policy successfully.
*/
type GetWebhookPolicyOfProjectOK struct {
	Payload *model.WebhookPolicy
}

func (o *GetWebhookPolicyOfProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectOK  %+v", 200, o.Payload)
}

func (o *GetWebhookPolicyOfProjectOK) GetPayload() *model.WebhookPolicy {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.WebhookPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookPolicyOfProjectBadRequest creates a GetWebhookPolicyOfProjectBadRequest with default headers values
func NewGetWebhookPolicyOfProjectBadRequest() *GetWebhookPolicyOfProjectBadRequest {
	return &GetWebhookPolicyOfProjectBadRequest{}
}

/*GetWebhookPolicyOfProjectBadRequest handles this case with default header values.

Bad request
*/
type GetWebhookPolicyOfProjectBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWebhookPolicyOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebhookPolicyOfProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookPolicyOfProjectUnauthorized creates a GetWebhookPolicyOfProjectUnauthorized with default headers values
func NewGetWebhookPolicyOfProjectUnauthorized() *GetWebhookPolicyOfProjectUnauthorized {
	return &GetWebhookPolicyOfProjectUnauthorized{}
}

/*GetWebhookPolicyOfProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type GetWebhookPolicyOfProjectUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWebhookPolicyOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebhookPolicyOfProjectUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookPolicyOfProjectForbidden creates a GetWebhookPolicyOfProjectForbidden with default headers values
func NewGetWebhookPolicyOfProjectForbidden() *GetWebhookPolicyOfProjectForbidden {
	return &GetWebhookPolicyOfProjectForbidden{}
}

/*GetWebhookPolicyOfProjectForbidden handles this case with default header values.

Forbidden
*/
type GetWebhookPolicyOfProjectForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWebhookPolicyOfProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetWebhookPolicyOfProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookPolicyOfProjectNotFound creates a GetWebhookPolicyOfProjectNotFound with default headers values
func NewGetWebhookPolicyOfProjectNotFound() *GetWebhookPolicyOfProjectNotFound {
	return &GetWebhookPolicyOfProjectNotFound{}
}

/*GetWebhookPolicyOfProjectNotFound handles this case with default header values.

Not found
*/
type GetWebhookPolicyOfProjectNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWebhookPolicyOfProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetWebhookPolicyOfProjectNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookPolicyOfProjectInternalServerError creates a GetWebhookPolicyOfProjectInternalServerError with default headers values
func NewGetWebhookPolicyOfProjectInternalServerError() *GetWebhookPolicyOfProjectInternalServerError {
	return &GetWebhookPolicyOfProjectInternalServerError{}
}

/*GetWebhookPolicyOfProjectInternalServerError handles this case with default header values.

Internal server error
*/
type GetWebhookPolicyOfProjectInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWebhookPolicyOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebhookPolicyOfProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
