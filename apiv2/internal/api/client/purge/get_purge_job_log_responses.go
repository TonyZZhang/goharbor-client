// Code generated by go-swagger; DO NOT EDIT.

package purge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// GetPurgeJobLogReader is a Reader for the GetPurgeJobLog structure.
type GetPurgeJobLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPurgeJobLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPurgeJobLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPurgeJobLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPurgeJobLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPurgeJobLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPurgeJobLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPurgeJobLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPurgeJobLogOK creates a GetPurgeJobLogOK with default headers values
func NewGetPurgeJobLogOK() *GetPurgeJobLogOK {
	return &GetPurgeJobLogOK{}
}

/*GetPurgeJobLogOK handles this case with default header values.

Get successfully.
*/
type GetPurgeJobLogOK struct {
	Payload string
}

func (o *GetPurgeJobLogOK) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}/log][%d] getPurgeJobLogOK  %+v", 200, o.Payload)
}

func (o *GetPurgeJobLogOK) GetPayload() string {
	return o.Payload
}

func (o *GetPurgeJobLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobLogBadRequest creates a GetPurgeJobLogBadRequest with default headers values
func NewGetPurgeJobLogBadRequest() *GetPurgeJobLogBadRequest {
	return &GetPurgeJobLogBadRequest{}
}

/*GetPurgeJobLogBadRequest handles this case with default header values.

Bad request
*/
type GetPurgeJobLogBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobLogBadRequest) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}/log][%d] getPurgeJobLogBadRequest  %+v", 400, o.Payload)
}

func (o *GetPurgeJobLogBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobLogUnauthorized creates a GetPurgeJobLogUnauthorized with default headers values
func NewGetPurgeJobLogUnauthorized() *GetPurgeJobLogUnauthorized {
	return &GetPurgeJobLogUnauthorized{}
}

/*GetPurgeJobLogUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPurgeJobLogUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}/log][%d] getPurgeJobLogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPurgeJobLogUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobLogForbidden creates a GetPurgeJobLogForbidden with default headers values
func NewGetPurgeJobLogForbidden() *GetPurgeJobLogForbidden {
	return &GetPurgeJobLogForbidden{}
}

/*GetPurgeJobLogForbidden handles this case with default header values.

Forbidden
*/
type GetPurgeJobLogForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobLogForbidden) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}/log][%d] getPurgeJobLogForbidden  %+v", 403, o.Payload)
}

func (o *GetPurgeJobLogForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobLogNotFound creates a GetPurgeJobLogNotFound with default headers values
func NewGetPurgeJobLogNotFound() *GetPurgeJobLogNotFound {
	return &GetPurgeJobLogNotFound{}
}

/*GetPurgeJobLogNotFound handles this case with default header values.

Not found
*/
type GetPurgeJobLogNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobLogNotFound) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}/log][%d] getPurgeJobLogNotFound  %+v", 404, o.Payload)
}

func (o *GetPurgeJobLogNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeJobLogInternalServerError creates a GetPurgeJobLogInternalServerError with default headers values
func NewGetPurgeJobLogInternalServerError() *GetPurgeJobLogInternalServerError {
	return &GetPurgeJobLogInternalServerError{}
}

/*GetPurgeJobLogInternalServerError handles this case with default header values.

Internal server error
*/
type GetPurgeJobLogInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeJobLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/{purge_id}/log][%d] getPurgeJobLogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPurgeJobLogInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeJobLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
