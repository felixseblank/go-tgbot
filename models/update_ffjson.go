// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/update.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Update) MarshalJSON() ([]byte, error) {
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
func (mj *Update) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.CallbackQuery != nil {
		if true {
			buf.WriteString(`"callback_query":`)

			{

				err = mj.CallbackQuery.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.ChannelPost != nil {
		if true {
			buf.WriteString(`"channel_post":`)

			{

				err = mj.ChannelPost.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.ChosenInlineResult != nil {
		if true {
			buf.WriteString(`"chosen_inline_result":`)

			{

				err = mj.ChosenInlineResult.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.EditedChannelPost != nil {
		if true {
			buf.WriteString(`"edited_channel_post":`)

			{

				err = mj.EditedChannelPost.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.EditedMessage != nil {
		if true {
			buf.WriteString(`"edited_message":`)

			{

				err = mj.EditedMessage.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.InlineQuery != nil {
		if true {
			buf.WriteString(`"inline_query":`)

			{

				err = mj.InlineQuery.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.Message != nil {
		if true {
			buf.WriteString(`"message":`)

			{

				err = mj.Message.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.UpdateID != 0 {
		buf.WriteString(`"update_id":`)
		fflib.FormatBits2(buf, uint64(mj.UpdateID), 10, mj.UpdateID < 0)
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Updatebase = iota
	ffj_t_Updateno_such_key

	ffj_t_Update_CallbackQuery

	ffj_t_Update_ChannelPost

	ffj_t_Update_ChosenInlineResult

	ffj_t_Update_EditedChannelPost

	ffj_t_Update_EditedMessage

	ffj_t_Update_InlineQuery

	ffj_t_Update_Message

	ffj_t_Update_UpdateID
)

var ffj_key_Update_CallbackQuery = []byte("callback_query")

var ffj_key_Update_ChannelPost = []byte("channel_post")

var ffj_key_Update_ChosenInlineResult = []byte("chosen_inline_result")

var ffj_key_Update_EditedChannelPost = []byte("edited_channel_post")

var ffj_key_Update_EditedMessage = []byte("edited_message")

var ffj_key_Update_InlineQuery = []byte("inline_query")

var ffj_key_Update_Message = []byte("message")

var ffj_key_Update_UpdateID = []byte("update_id")

func (uj *Update) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Update) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Updatebase
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
				currentKey = ffj_t_Updateno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_Update_CallbackQuery, kn) {
						currentKey = ffj_t_Update_CallbackQuery
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Update_ChannelPost, kn) {
						currentKey = ffj_t_Update_ChannelPost
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Update_ChosenInlineResult, kn) {
						currentKey = ffj_t_Update_ChosenInlineResult
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Update_EditedChannelPost, kn) {
						currentKey = ffj_t_Update_EditedChannelPost
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Update_EditedMessage, kn) {
						currentKey = ffj_t_Update_EditedMessage
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Update_InlineQuery, kn) {
						currentKey = ffj_t_Update_InlineQuery
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_Update_Message, kn) {
						currentKey = ffj_t_Update_Message
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_Update_UpdateID, kn) {
						currentKey = ffj_t_Update_UpdateID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffj_key_Update_UpdateID, kn) {
					currentKey = ffj_t_Update_UpdateID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Update_Message, kn) {
					currentKey = ffj_t_Update_Message
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_Update_InlineQuery, kn) {
					currentKey = ffj_t_Update_InlineQuery
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Update_EditedMessage, kn) {
					currentKey = ffj_t_Update_EditedMessage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Update_EditedChannelPost, kn) {
					currentKey = ffj_t_Update_EditedChannelPost
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Update_ChosenInlineResult, kn) {
					currentKey = ffj_t_Update_ChosenInlineResult
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Update_ChannelPost, kn) {
					currentKey = ffj_t_Update_ChannelPost
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Update_CallbackQuery, kn) {
					currentKey = ffj_t_Update_CallbackQuery
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Updateno_such_key
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

				case ffj_t_Update_CallbackQuery:
					goto handle_CallbackQuery

				case ffj_t_Update_ChannelPost:
					goto handle_ChannelPost

				case ffj_t_Update_ChosenInlineResult:
					goto handle_ChosenInlineResult

				case ffj_t_Update_EditedChannelPost:
					goto handle_EditedChannelPost

				case ffj_t_Update_EditedMessage:
					goto handle_EditedMessage

				case ffj_t_Update_InlineQuery:
					goto handle_InlineQuery

				case ffj_t_Update_Message:
					goto handle_Message

				case ffj_t_Update_UpdateID:
					goto handle_UpdateID

				case ffj_t_Updateno_such_key:
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

handle_CallbackQuery:

	/* handler: uj.CallbackQuery type=models.CallbackQuery kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.CallbackQuery = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.CallbackQuery == nil {
			uj.CallbackQuery = new(CallbackQuery)
		}

		err = uj.CallbackQuery.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ChannelPost:

	/* handler: uj.ChannelPost type=models.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.ChannelPost = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.ChannelPost == nil {
			uj.ChannelPost = new(Message)
		}

		err = uj.ChannelPost.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ChosenInlineResult:

	/* handler: uj.ChosenInlineResult type=models.ChosenInlineResult kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.ChosenInlineResult = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.ChosenInlineResult == nil {
			uj.ChosenInlineResult = new(ChosenInlineResult)
		}

		err = uj.ChosenInlineResult.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_EditedChannelPost:

	/* handler: uj.EditedChannelPost type=models.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.EditedChannelPost = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.EditedChannelPost == nil {
			uj.EditedChannelPost = new(Message)
		}

		err = uj.EditedChannelPost.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_EditedMessage:

	/* handler: uj.EditedMessage type=models.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.EditedMessage = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.EditedMessage == nil {
			uj.EditedMessage = new(Message)
		}

		err = uj.EditedMessage.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InlineQuery:

	/* handler: uj.InlineQuery type=models.InlineQuery kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.InlineQuery = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.InlineQuery == nil {
			uj.InlineQuery = new(InlineQuery)
		}

		err = uj.InlineQuery.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Message:

	/* handler: uj.Message type=models.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Message = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Message == nil {
			uj.Message = new(Message)
		}

		err = uj.Message.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UpdateID:

	/* handler: uj.UpdateID type=int64 kind=int64 quoted=false*/

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

			uj.UpdateID = int64(tval)

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
