package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// GetChatReader is a Reader for the GetChat structure.
type GetChatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetChatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChatOK creates a GetChatOK with default headers values
func NewGetChatOK() *GetChatOK {
	return &GetChatOK{}
}

/*GetChatOK handles this case with default header values.

GetChatOK get chat o k
*/
type GetChatOK struct {
	Payload GetChatOKBody
}

func (o *GetChatOK) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChat][%d] getChatOK  %+v", 200, o.Payload)
}

func (o *GetChatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChatBadRequest creates a GetChatBadRequest with default headers values
func NewGetChatBadRequest() *GetChatBadRequest {
	return &GetChatBadRequest{}
}

/*GetChatBadRequest handles this case with default header values.

Error
*/
type GetChatBadRequest struct {
	Payload *models.Error
}

func (o *GetChatBadRequest) Error() string {
	return fmt.Sprintf("[GET /bot{token}/getChat][%d] getChatBadRequest  %+v", 400, o.Payload)
}

func (o *GetChatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetChatOKBody get chat o k body
swagger:model GetChatOKBody
*/
type GetChatOKBody struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// error code
	// Required: true
	ErrorCode *int64 `json:"error_code"`

	// ok
	// Required: true
	Ok *bool `json:"ok"`

	// result
	// Required: true
	Result *models.Chat `json:"result"`
}
