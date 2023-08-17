// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// DeleteArtifactReader is a Reader for the DeleteArtifact structure.
type DeleteArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteArtifactOK creates a DeleteArtifactOK with default headers values
func NewDeleteArtifactOK() *DeleteArtifactOK {
	return &DeleteArtifactOK{}
}

/*DeleteArtifactOK handles this case with default header values.

Success
*/
type DeleteArtifactOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteArtifactOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] deleteArtifactOK ", 200)
}

func (o *DeleteArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteArtifactUnauthorized creates a DeleteArtifactUnauthorized with default headers values
func NewDeleteArtifactUnauthorized() *DeleteArtifactUnauthorized {
	return &DeleteArtifactUnauthorized{}
}

/*DeleteArtifactUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteArtifactUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] deleteArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteArtifactUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactForbidden creates a DeleteArtifactForbidden with default headers values
func NewDeleteArtifactForbidden() *DeleteArtifactForbidden {
	return &DeleteArtifactForbidden{}
}

/*DeleteArtifactForbidden handles this case with default header values.

Forbidden
*/
type DeleteArtifactForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteArtifactForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] deleteArtifactForbidden  %+v", 403, o.Payload)
}

func (o *DeleteArtifactForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactNotFound creates a DeleteArtifactNotFound with default headers values
func NewDeleteArtifactNotFound() *DeleteArtifactNotFound {
	return &DeleteArtifactNotFound{}
}

/*DeleteArtifactNotFound handles this case with default header values.

Not found
*/
type DeleteArtifactNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteArtifactNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] deleteArtifactNotFound  %+v", 404, o.Payload)
}

func (o *DeleteArtifactNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactInternalServerError creates a DeleteArtifactInternalServerError with default headers values
func NewDeleteArtifactInternalServerError() *DeleteArtifactInternalServerError {
	return &DeleteArtifactInternalServerError{}
}

/*DeleteArtifactInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteArtifactInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] deleteArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteArtifactInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
