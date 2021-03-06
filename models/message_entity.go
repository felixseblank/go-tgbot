// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MessageEntity message entity
// swagger:model MessageEntity
type MessageEntity struct {

	// length
	Length int64 `json:"length,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this message entity
func (m *MessageEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var messageEntityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mention","hashtag","bot_command","url","email","bold","italic","code","pre","text_link","text_mention"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageEntityTypeTypePropEnum = append(messageEntityTypeTypePropEnum, v)
	}
}

const (
	// MessageEntityTypeMention captures enum value "mention"
	MessageEntityTypeMention string = "mention"
	// MessageEntityTypeHashtag captures enum value "hashtag"
	MessageEntityTypeHashtag string = "hashtag"
	// MessageEntityTypeBotCommand captures enum value "bot_command"
	MessageEntityTypeBotCommand string = "bot_command"
	// MessageEntityTypeURL captures enum value "url"
	MessageEntityTypeURL string = "url"
	// MessageEntityTypeEmail captures enum value "email"
	MessageEntityTypeEmail string = "email"
	// MessageEntityTypeBold captures enum value "bold"
	MessageEntityTypeBold string = "bold"
	// MessageEntityTypeItalic captures enum value "italic"
	MessageEntityTypeItalic string = "italic"
	// MessageEntityTypeCode captures enum value "code"
	MessageEntityTypeCode string = "code"
	// MessageEntityTypePre captures enum value "pre"
	MessageEntityTypePre string = "pre"
	// MessageEntityTypeTextLink captures enum value "text_link"
	MessageEntityTypeTextLink string = "text_link"
	// MessageEntityTypeTextMention captures enum value "text_mention"
	MessageEntityTypeTextMention string = "text_mention"
)

// prop value enum
func (m *MessageEntity) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, messageEntityTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MessageEntity) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *MessageEntity) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {

		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageEntity) UnmarshalBinary(b []byte) error {
	var res MessageEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
