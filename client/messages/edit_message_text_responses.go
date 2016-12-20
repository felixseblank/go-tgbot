package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// EditMessageTextReader is a Reader for the EditMessageText structure.
type EditMessageTextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditMessageTextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditMessageTextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditMessageTextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditMessageTextOK creates a EditMessageTextOK with default headers values
func NewEditMessageTextOK() *EditMessageTextOK {
	return &EditMessageTextOK{}
}

/*EditMessageTextOK handles this case with default header values.

EditMessageTextOK edit message text o k
*/
type EditMessageTextOK struct {
	Payload EditMessageTextOKBody
}

func (o *EditMessageTextOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextOK  %+v", 200, o.Payload)
}

func (o *EditMessageTextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageTextBadRequest creates a EditMessageTextBadRequest with default headers values
func NewEditMessageTextBadRequest() *EditMessageTextBadRequest {
	return &EditMessageTextBadRequest{}
}

/*EditMessageTextBadRequest handles this case with default header values.

Error
*/
type EditMessageTextBadRequest struct {
	Payload *models.Error
}

func (o *EditMessageTextBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextBadRequest  %+v", 400, o.Payload)
}

func (o *EditMessageTextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EditMessageTextBody edit message text body
swagger:model EditMessageTextBody
*/
type EditMessageTextBody struct {

	// chat id
	ChatID interface{} `json:"chat_id,omitempty"`

	// disable web page preview
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// inline message id
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// message id
	MessageID int64 `json:"message_id,omitempty"`

	// parse mode
	ParseMode string `json:"parse_mode,omitempty"`

	// reply markup
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// text
	// Required: true
	Text *string `json:"text"`
}

/*EditMessageTextOKBody edit message text o k body
swagger:model EditMessageTextOKBody
*/
type EditMessageTextOKBody struct {

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
	Result interface{} `json:"result"`
}
