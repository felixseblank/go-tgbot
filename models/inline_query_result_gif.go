// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineQueryResultGif inline query result gif
// swagger:model InlineQueryResultGif
type InlineQueryResultGif struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// gif duration
	GifDuration int64 `json:"gif_duration,omitempty"`

	// gif height
	GifHeight int64 `json:"gif_height,omitempty"`

	// gif url
	// Required: true
	GifURL *string `json:"gif_url"`

	// gif width
	GifWidth int64 `json:"gif_width,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// parse mode
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// thumb url
	// Required: true
	ThumbURL *string `json:"thumb_url"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Required: true
	Type InlineType `json:"type"`
}

// Validate validates this inline query result gif
func (m *InlineQueryResultGif) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGifURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParseMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThumbURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultGif) validateGifURL(formats strfmt.Registry) error {

	if err := validate.Required("gif_url", "body", m.GifURL); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultGif) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultGif) validateParseMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ParseMode) { // not required
		return nil
	}

	if err := m.ParseMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parse_mode")
		}
		return err
	}

	return nil
}

func (m *InlineQueryResultGif) validateReplyMarkup(formats strfmt.Registry) error {

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

func (m *InlineQueryResultGif) validateThumbURL(formats strfmt.Registry) error {

	if err := validate.Required("thumb_url", "body", m.ThumbURL); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultGif) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultGif) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultGif) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultGif
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
