// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// GetInternalconfigReader is a Reader for the GetInternalconfig structure.
type GetInternalconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternalconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInternalconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetInternalconfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInternalconfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInternalconfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInternalconfigOK creates a GetInternalconfigOK with default headers values
func NewGetInternalconfigOK() *GetInternalconfigOK {
	return &GetInternalconfigOK{}
}

/*GetInternalconfigOK handles this case with default header values.

Get system configurations successfully. The response body is a map.
*/
type GetInternalconfigOK struct {
	Payload model.InternalConfigurationsResponse
}

func (o *GetInternalconfigOK) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigOK  %+v", 200, o.Payload)
}

func (o *GetInternalconfigOK) GetPayload() model.InternalConfigurationsResponse {
	return o.Payload
}

func (o *GetInternalconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternalconfigUnauthorized creates a GetInternalconfigUnauthorized with default headers values
func NewGetInternalconfigUnauthorized() *GetInternalconfigUnauthorized {
	return &GetInternalconfigUnauthorized{}
}

/*GetInternalconfigUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetInternalconfigUnauthorized struct {
}

func (o *GetInternalconfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigUnauthorized ", 401)
}

func (o *GetInternalconfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInternalconfigForbidden creates a GetInternalconfigForbidden with default headers values
func NewGetInternalconfigForbidden() *GetInternalconfigForbidden {
	return &GetInternalconfigForbidden{}
}

/*GetInternalconfigForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type GetInternalconfigForbidden struct {
}

func (o *GetInternalconfigForbidden) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigForbidden ", 403)
}

func (o *GetInternalconfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInternalconfigInternalServerError creates a GetInternalconfigInternalServerError with default headers values
func NewGetInternalconfigInternalServerError() *GetInternalconfigInternalServerError {
	return &GetInternalconfigInternalServerError{}
}

/*GetInternalconfigInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetInternalconfigInternalServerError struct {
}

func (o *GetInternalconfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigInternalServerError ", 500)
}

func (o *GetInternalconfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
