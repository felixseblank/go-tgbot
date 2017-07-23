package stickers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SetStickerPositionInSetReader is a Reader for the SetStickerPositionInSet structure.
type SetStickerPositionInSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetStickerPositionInSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetStickerPositionInSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetStickerPositionInSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSetStickerPositionInSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetStickerPositionInSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetStickerPositionInSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSetStickerPositionInSetEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSetStickerPositionInSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetStickerPositionInSetOK creates a SetStickerPositionInSetOK with default headers values
func NewSetStickerPositionInSetOK() *SetStickerPositionInSetOK {
	return &SetStickerPositionInSetOK{}
}

/*SetStickerPositionInSetOK handles this case with default header values.

SetStickerPositionInSetOK set sticker position in set o k
*/
type SetStickerPositionInSetOK struct {
	Payload bool
}

func (o *SetStickerPositionInSetOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setStickerPositionInSet][%d] setStickerPositionInSetOK  %+v", 200, o.Payload)
}

func (o *SetStickerPositionInSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetStickerPositionInSetBadRequest creates a SetStickerPositionInSetBadRequest with default headers values
func NewSetStickerPositionInSetBadRequest() *SetStickerPositionInSetBadRequest {
	return &SetStickerPositionInSetBadRequest{}
}

/*SetStickerPositionInSetBadRequest handles this case with default header values.

Bad Request
*/
type SetStickerPositionInSetBadRequest struct {
	Payload *models.Error
}

func (o *SetStickerPositionInSetBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setStickerPositionInSet][%d] setStickerPositionInSetBadRequest  %+v", 400, o.Payload)
}

func (o *SetStickerPositionInSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetStickerPositionInSetUnauthorized creates a SetStickerPositionInSetUnauthorized with default headers values
func NewSetStickerPositionInSetUnauthorized() *SetStickerPositionInSetUnauthorized {
	return &SetStickerPositionInSetUnauthorized{}
}

/*SetStickerPositionInSetUnauthorized handles this case with default header values.

Unauthorized
*/
type SetStickerPositionInSetUnauthorized struct {
	Payload *models.Error
}

func (o *SetStickerPositionInSetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setStickerPositionInSet][%d] setStickerPositionInSetUnauthorized  %+v", 401, o.Payload)
}

func (o *SetStickerPositionInSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetStickerPositionInSetForbidden creates a SetStickerPositionInSetForbidden with default headers values
func NewSetStickerPositionInSetForbidden() *SetStickerPositionInSetForbidden {
	return &SetStickerPositionInSetForbidden{}
}

/*SetStickerPositionInSetForbidden handles this case with default header values.

Forbidden
*/
type SetStickerPositionInSetForbidden struct {
	Payload *models.Error
}

func (o *SetStickerPositionInSetForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setStickerPositionInSet][%d] setStickerPositionInSetForbidden  %+v", 403, o.Payload)
}

func (o *SetStickerPositionInSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetStickerPositionInSetNotFound creates a SetStickerPositionInSetNotFound with default headers values
func NewSetStickerPositionInSetNotFound() *SetStickerPositionInSetNotFound {
	return &SetStickerPositionInSetNotFound{}
}

/*SetStickerPositionInSetNotFound handles this case with default header values.

Not Found
*/
type SetStickerPositionInSetNotFound struct {
	Payload *models.Error
}

func (o *SetStickerPositionInSetNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setStickerPositionInSet][%d] setStickerPositionInSetNotFound  %+v", 404, o.Payload)
}

func (o *SetStickerPositionInSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetStickerPositionInSetEnhanceYourCalm creates a SetStickerPositionInSetEnhanceYourCalm with default headers values
func NewSetStickerPositionInSetEnhanceYourCalm() *SetStickerPositionInSetEnhanceYourCalm {
	return &SetStickerPositionInSetEnhanceYourCalm{}
}

/*SetStickerPositionInSetEnhanceYourCalm handles this case with default header values.

Flood
*/
type SetStickerPositionInSetEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SetStickerPositionInSetEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setStickerPositionInSet][%d] setStickerPositionInSetEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SetStickerPositionInSetEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetStickerPositionInSetInternalServerError creates a SetStickerPositionInSetInternalServerError with default headers values
func NewSetStickerPositionInSetInternalServerError() *SetStickerPositionInSetInternalServerError {
	return &SetStickerPositionInSetInternalServerError{}
}

/*SetStickerPositionInSetInternalServerError handles this case with default header values.

Internal
*/
type SetStickerPositionInSetInternalServerError struct {
	Payload *models.Error
}

func (o *SetStickerPositionInSetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/setStickerPositionInSet][%d] setStickerPositionInSetInternalServerError  %+v", 500, o.Payload)
}

func (o *SetStickerPositionInSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
