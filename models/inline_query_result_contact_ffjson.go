// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/inline_query_result_contact.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *InlineQueryResultContact) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *InlineQueryResultContact) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	if mj.FirstName != nil {
		buf.WriteString(`{"first_name":`)
		fflib.WriteJsonString(buf, string(*mj.FirstName))
	} else {
		buf.WriteString(`{"first_name":null`)
	}
	if mj.ID != nil {
		buf.WriteString(`,"id":`)
		fflib.WriteJsonString(buf, string(*mj.ID))
	} else {
		buf.WriteString(`,"id":null`)
	}
	buf.WriteByte(',')
	if mj.InputMessageContent != nil {
		buf.WriteString(`"input_message_content":`)
		/* Interface types must use runtime reflection. type=interface {} kind=interface */
		err = buf.Encode(mj.InputMessageContent)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.LastName) != 0 {
		buf.WriteString(`"last_name":`)
		fflib.WriteJsonString(buf, string(mj.LastName))
		buf.WriteByte(',')
	}
	if mj.PhoneNumber != nil {
		buf.WriteString(`"phone_number":`)
		fflib.WriteJsonString(buf, string(*mj.PhoneNumber))
	} else {
		buf.WriteString(`"phone_number":null`)
	}
	buf.WriteByte(',')
	if mj.ReplyMarkup != nil {
		if true {
			buf.WriteString(`"reply_markup":`)

			{

				err = mj.ReplyMarkup.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.ThumbHeight != 0 {
		buf.WriteString(`"thumb_height":`)
		fflib.FormatBits2(buf, uint64(mj.ThumbHeight), 10, mj.ThumbHeight < 0)
		buf.WriteByte(',')
	}
	if len(mj.ThumbURL) != 0 {
		buf.WriteString(`"thumb_url":`)
		fflib.WriteJsonString(buf, string(mj.ThumbURL))
		buf.WriteByte(',')
	}
	if mj.ThumbWidth != 0 {
		buf.WriteString(`"thumb_width":`)
		fflib.FormatBits2(buf, uint64(mj.ThumbWidth), 10, mj.ThumbWidth < 0)
		buf.WriteByte(',')
	}
	buf.WriteString(`"type":`)
	fflib.WriteJsonString(buf, string(mj.Type))
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_InlineQueryResultContactbase = iota
	ffj_t_InlineQueryResultContactno_such_key

	ffj_t_InlineQueryResultContact_FirstName

	ffj_t_InlineQueryResultContact_ID

	ffj_t_InlineQueryResultContact_InputMessageContent

	ffj_t_InlineQueryResultContact_LastName

	ffj_t_InlineQueryResultContact_PhoneNumber

	ffj_t_InlineQueryResultContact_ReplyMarkup

	ffj_t_InlineQueryResultContact_ThumbHeight

	ffj_t_InlineQueryResultContact_ThumbURL

	ffj_t_InlineQueryResultContact_ThumbWidth

	ffj_t_InlineQueryResultContact_Type
)

var ffj_key_InlineQueryResultContact_FirstName = []byte("first_name")

var ffj_key_InlineQueryResultContact_ID = []byte("id")

var ffj_key_InlineQueryResultContact_InputMessageContent = []byte("input_message_content")

var ffj_key_InlineQueryResultContact_LastName = []byte("last_name")

var ffj_key_InlineQueryResultContact_PhoneNumber = []byte("phone_number")

var ffj_key_InlineQueryResultContact_ReplyMarkup = []byte("reply_markup")

var ffj_key_InlineQueryResultContact_ThumbHeight = []byte("thumb_height")

var ffj_key_InlineQueryResultContact_ThumbURL = []byte("thumb_url")

var ffj_key_InlineQueryResultContact_ThumbWidth = []byte("thumb_width")

var ffj_key_InlineQueryResultContact_Type = []byte("type")

func (uj *InlineQueryResultContact) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *InlineQueryResultContact) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_InlineQueryResultContactbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_InlineQueryResultContactno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'f':

					if bytes.Equal(ffj_key_InlineQueryResultContact_FirstName, kn) {
						currentKey = ffj_t_InlineQueryResultContact_FirstName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_InlineQueryResultContact_ID, kn) {
						currentKey = ffj_t_InlineQueryResultContact_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultContact_InputMessageContent, kn) {
						currentKey = ffj_t_InlineQueryResultContact_InputMessageContent
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_InlineQueryResultContact_LastName, kn) {
						currentKey = ffj_t_InlineQueryResultContact_LastName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_InlineQueryResultContact_PhoneNumber, kn) {
						currentKey = ffj_t_InlineQueryResultContact_PhoneNumber
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_InlineQueryResultContact_ReplyMarkup, kn) {
						currentKey = ffj_t_InlineQueryResultContact_ReplyMarkup
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_InlineQueryResultContact_ThumbHeight, kn) {
						currentKey = ffj_t_InlineQueryResultContact_ThumbHeight
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultContact_ThumbURL, kn) {
						currentKey = ffj_t_InlineQueryResultContact_ThumbURL
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultContact_ThumbWidth, kn) {
						currentKey = ffj_t_InlineQueryResultContact_ThumbWidth
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_InlineQueryResultContact_Type, kn) {
						currentKey = ffj_t_InlineQueryResultContact_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_InlineQueryResultContact_Type, kn) {
					currentKey = ffj_t_InlineQueryResultContact_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultContact_ThumbWidth, kn) {
					currentKey = ffj_t_InlineQueryResultContact_ThumbWidth
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultContact_ThumbURL, kn) {
					currentKey = ffj_t_InlineQueryResultContact_ThumbURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultContact_ThumbHeight, kn) {
					currentKey = ffj_t_InlineQueryResultContact_ThumbHeight
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InlineQueryResultContact_ReplyMarkup, kn) {
					currentKey = ffj_t_InlineQueryResultContact_ReplyMarkup
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_InlineQueryResultContact_PhoneNumber, kn) {
					currentKey = ffj_t_InlineQueryResultContact_PhoneNumber
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InlineQueryResultContact_LastName, kn) {
					currentKey = ffj_t_InlineQueryResultContact_LastName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InlineQueryResultContact_InputMessageContent, kn) {
					currentKey = ffj_t_InlineQueryResultContact_InputMessageContent
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_InlineQueryResultContact_ID, kn) {
					currentKey = ffj_t_InlineQueryResultContact_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InlineQueryResultContact_FirstName, kn) {
					currentKey = ffj_t_InlineQueryResultContact_FirstName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_InlineQueryResultContactno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_InlineQueryResultContact_FirstName:
					goto handle_FirstName

				case ffj_t_InlineQueryResultContact_ID:
					goto handle_ID

				case ffj_t_InlineQueryResultContact_InputMessageContent:
					goto handle_InputMessageContent

				case ffj_t_InlineQueryResultContact_LastName:
					goto handle_LastName

				case ffj_t_InlineQueryResultContact_PhoneNumber:
					goto handle_PhoneNumber

				case ffj_t_InlineQueryResultContact_ReplyMarkup:
					goto handle_ReplyMarkup

				case ffj_t_InlineQueryResultContact_ThumbHeight:
					goto handle_ThumbHeight

				case ffj_t_InlineQueryResultContact_ThumbURL:
					goto handle_ThumbURL

				case ffj_t_InlineQueryResultContact_ThumbWidth:
					goto handle_ThumbWidth

				case ffj_t_InlineQueryResultContact_Type:
					goto handle_Type

				case ffj_t_InlineQueryResultContactno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_FirstName:

	/* handler: uj.FirstName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.FirstName = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.FirstName = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ID:

	/* handler: uj.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.ID = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.ID = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InputMessageContent:

	/* handler: uj.InputMessageContent type=interface {} kind=interface quoted=false*/

	{
		/* Falling back. type=interface {} kind=interface */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.InputMessageContent)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastName:

	/* handler: uj.LastName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.LastName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PhoneNumber:

	/* handler: uj.PhoneNumber type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.PhoneNumber = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.PhoneNumber = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReplyMarkup:

	/* handler: uj.ReplyMarkup type=models.InlineKeyboardMarkup kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.ReplyMarkup = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.ReplyMarkup == nil {
			uj.ReplyMarkup = new(InlineKeyboardMarkup)
		}

		err = uj.ReplyMarkup.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ThumbHeight:

	/* handler: uj.ThumbHeight type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.ThumbHeight = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ThumbURL:

	/* handler: uj.ThumbURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ThumbURL = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ThumbWidth:

	/* handler: uj.ThumbWidth type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.ThumbWidth = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Type:

	/* handler: uj.Type type=models.InlineType kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for InlineType", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Type = InlineType(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}