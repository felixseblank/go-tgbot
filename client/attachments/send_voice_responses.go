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

// SendVoiceReader is a Reader for the SendVoice structure.
type SendVoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVoiceOK creates a SendVoiceOK with default headers values
func NewSendVoiceOK() *SendVoiceOK {
	return &SendVoiceOK{}
}

/*SendVoiceOK handles this case with default header values.

SendVoiceOK send voice o k
*/
type SendVoiceOK struct {
	Payload *models.ResponseMessage
}

func (o *SendVoiceOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVoice][%d] sendVoiceOK  %+v", 200, o.Payload)
}

func (o *SendVoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}