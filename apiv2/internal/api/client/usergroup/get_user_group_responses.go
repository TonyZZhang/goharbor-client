// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// GetUserGroupReader is a Reader for the GetUserGroup structure.
type GetUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGroupOK creates a GetUserGroupOK with default headers values
func NewGetUserGroupOK() *GetUserGroupOK {
	return &GetUserGroupOK{}
}

/*GetUserGroupOK handles this case with default header values.

User group get successfully.
*/
type GetUserGroupOK struct {
	Payload *model.UserGroup
}

func (o *GetUserGroupOK) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUserGroupOK  %+v", 200, o.Payload)
}

func (o *GetUserGroupOK) GetPayload() *model.UserGroup {
	return o.Payload
}

func (o *GetUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.UserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupBadRequest creates a GetUserGroupBadRequest with default headers values
func NewGetUserGroupBadRequest() *GetUserGroupBadRequest {
	return &GetUserGroupBadRequest{}
}

/*GetUserGroupBadRequest handles this case with default header values.

Bad request
*/
type GetUserGroupBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUserGroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGroupBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupUnauthorized creates a GetUserGroupUnauthorized with default headers values
func NewGetUserGroupUnauthorized() *GetUserGroupUnauthorized {
	return &GetUserGroupUnauthorized{}
}

/*GetUserGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUserGroupUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUserGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGroupUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupForbidden creates a GetUserGroupForbidden with default headers values
func NewGetUserGroupForbidden() *GetUserGroupForbidden {
	return &GetUserGroupForbidden{}
}

/*GetUserGroupForbidden handles this case with default header values.

Forbidden
*/
type GetUserGroupForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetUserGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUserGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGroupForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupNotFound creates a GetUserGroupNotFound with default headers values
func NewGetUserGroupNotFound() *GetUserGroupNotFound {
	return &GetUserGroupNotFound{}
}

/*GetUserGroupNotFound handles this case with default header values.

Not found
*/
type GetUserGroupNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetUserGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUserGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGroupNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupInternalServerError creates a GetUserGroupInternalServerError with default headers values
func NewGetUserGroupInternalServerError() *GetUserGroupInternalServerError {
	return &GetUserGroupInternalServerError{}
}

/*GetUserGroupInternalServerError handles this case with default header values.

Internal server error
*/
type GetUserGroupInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetUserGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /usergroups/{group_id}][%d] getUserGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGroupInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetUserGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
