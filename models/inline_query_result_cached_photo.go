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

// InlineQueryResultCachedPhoto inline query result cached photo
// swagger:model InlineQueryResultCachedPhoto
type InlineQueryResultCachedPhoto struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// parse mode
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// photo file id
	// Required: true
	PhotoFileID *string `json:"photo_file_id"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Required: true
	Type InlineType `json:"type"`
}

// Validate validates this inline query result cached photo
func (m *InlineQueryResultCachedPhoto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParseMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhotoFileID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
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

func (m *InlineQueryResultCachedPhoto) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedPhoto) validateParseMode(formats strfmt.Registry) error {

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

func (m *InlineQueryResultCachedPhoto) validatePhotoFileID(formats strfmt.Registry) error {

	if err := validate.Required("photo_file_id", "body", m.PhotoFileID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedPhoto) validateReplyMarkup(formats strfmt.Registry) error {

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

func (m *InlineQueryResultCachedPhoto) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultCachedPhoto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultCachedPhoto) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultCachedPhoto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
