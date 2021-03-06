// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EditMessageReplyMarkupBody edit message reply markup body
// swagger:model EditMessageReplyMarkupBody
type EditMessageReplyMarkupBody struct {

	// chat id
	ChatID interface{} `json:"chat_id,omitempty"`

	// inline message id
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// message id
	MessageID int64 `json:"message_id,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Validate validates this edit message reply markup body
func (m *EditMessageReplyMarkupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditMessageReplyMarkupBody) validateReplyMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyMarkup) { // not required
		return nil
	}

	if m.ReplyMarkup != nil {

		if err := m.ReplyMarkup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply_markup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditMessageReplyMarkupBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditMessageReplyMarkupBody) UnmarshalBinary(b []byte) error {
	var res EditMessageReplyMarkupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
