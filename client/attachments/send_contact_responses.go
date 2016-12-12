package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SendContactReader is a Reader for the SendContact structure.
type SendContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendContactOK creates a SendContactOK with default headers values
func NewSendContactOK() *SendContactOK {
	return &SendContactOK{}
}

/*SendContactOK handles this case with default header values.

SendContactOK send contact o k
*/
type SendContactOK struct {
	Payload *models.ResponseMessage
}

func (o *SendContactOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendContact][%d] sendContactOK  %+v", 200, o.Payload)
}

func (o *SendContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SendContactBody send contact body
swagger:model SendContactBody
*/
type SendContactBody struct {

	// chat id
	// Required: true
	ChatID interface{} `json:"chat_id"`

	// disable notification
	DisableNotification bool `json:"disable_notification,omitempty"`

	// first name
	// Required: true
	FirstName *string `json:"first_name"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// phone number
	// Required: true
	PhoneNumber *string `json:"phone_number"`

	// reply markup
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`

	// reply to message id
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`
}
