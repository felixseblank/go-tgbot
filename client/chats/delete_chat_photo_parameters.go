// Code generated by go-swagger; DO NOT EDIT.

package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteChatPhotoParams creates a new DeleteChatPhotoParams object
// with the default values initialized.
func NewDeleteChatPhotoParams() *DeleteChatPhotoParams {
	var ()
	return &DeleteChatPhotoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteChatPhotoParamsWithTimeout creates a new DeleteChatPhotoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteChatPhotoParamsWithTimeout(timeout time.Duration) *DeleteChatPhotoParams {
	var ()
	return &DeleteChatPhotoParams{

		timeout: timeout,
	}
}

// NewDeleteChatPhotoParamsWithContext creates a new DeleteChatPhotoParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteChatPhotoParamsWithContext(ctx context.Context) *DeleteChatPhotoParams {
	var ()
	return &DeleteChatPhotoParams{

		Context: ctx,
	}
}

// NewDeleteChatPhotoParamsWithHTTPClient creates a new DeleteChatPhotoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteChatPhotoParamsWithHTTPClient(client *http.Client) *DeleteChatPhotoParams {
	var ()
	return &DeleteChatPhotoParams{
		HTTPClient: client,
	}
}

/*DeleteChatPhotoParams contains all the parameters to send to the API endpoint
for the delete chat photo operation typically these are written to a http.Request
*/
type DeleteChatPhotoParams struct {

	/*ChatID*/
	ChatID string
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete chat photo params
func (o *DeleteChatPhotoParams) WithTimeout(timeout time.Duration) *DeleteChatPhotoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete chat photo params
func (o *DeleteChatPhotoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete chat photo params
func (o *DeleteChatPhotoParams) WithContext(ctx context.Context) *DeleteChatPhotoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete chat photo params
func (o *DeleteChatPhotoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete chat photo params
func (o *DeleteChatPhotoParams) WithHTTPClient(client *http.Client) *DeleteChatPhotoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete chat photo params
func (o *DeleteChatPhotoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChatID adds the chatID to the delete chat photo params
func (o *DeleteChatPhotoParams) WithChatID(chatID string) *DeleteChatPhotoParams {
	o.SetChatID(chatID)
	return o
}

// SetChatID adds the chatId to the delete chat photo params
func (o *DeleteChatPhotoParams) SetChatID(chatID string) {
	o.ChatID = chatID
}

// WithToken adds the token to the delete chat photo params
func (o *DeleteChatPhotoParams) WithToken(token *string) *DeleteChatPhotoParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete chat photo params
func (o *DeleteChatPhotoParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteChatPhotoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param chat_id
	qrChatID := o.ChatID
	qChatID := qrChatID
	if qChatID != "" {
		if err := r.SetQueryParam("chat_id", qChatID); err != nil {
			return err
		}
	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
