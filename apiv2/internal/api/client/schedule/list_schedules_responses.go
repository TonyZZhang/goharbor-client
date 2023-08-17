// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// ListSchedulesReader is a Reader for the ListSchedules structure.
type ListSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSchedulesOK creates a ListSchedulesOK with default headers values
func NewListSchedulesOK() *ListSchedulesOK {
	return &ListSchedulesOK{}
}

/*ListSchedulesOK handles this case with default header values.

list schedule successfully.
*/
type ListSchedulesOK struct {
	/*Link to previous page and next page
	 */
	Link string
	/*The total count of available items
	 */
	XTotalCount int64

	Payload []*model.ScheduleTask
}

func (o *ListSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesOK  %+v", 200, o.Payload)
}

func (o *ListSchedulesOK) GetPayload() []*model.ScheduleTask {
	return o.Payload
}

func (o *ListSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSchedulesUnauthorized creates a ListSchedulesUnauthorized with default headers values
func NewListSchedulesUnauthorized() *ListSchedulesUnauthorized {
	return &ListSchedulesUnauthorized{}
}

/*ListSchedulesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListSchedulesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListSchedulesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSchedulesForbidden creates a ListSchedulesForbidden with default headers values
func NewListSchedulesForbidden() *ListSchedulesForbidden {
	return &ListSchedulesForbidden{}
}

/*ListSchedulesForbidden handles this case with default header values.

Forbidden
*/
type ListSchedulesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListSchedulesForbidden) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *ListSchedulesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSchedulesNotFound creates a ListSchedulesNotFound with default headers values
func NewListSchedulesNotFound() *ListSchedulesNotFound {
	return &ListSchedulesNotFound{}
}

/*ListSchedulesNotFound handles this case with default header values.

Not found
*/
type ListSchedulesNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListSchedulesNotFound) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *ListSchedulesNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSchedulesInternalServerError creates a ListSchedulesInternalServerError with default headers values
func NewListSchedulesInternalServerError() *ListSchedulesInternalServerError {
	return &ListSchedulesInternalServerError{}
}

/*ListSchedulesInternalServerError handles this case with default header values.

Internal server error
*/
type ListSchedulesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSchedulesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
