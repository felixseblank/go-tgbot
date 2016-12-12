// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/input_text_message_content.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *InputTextMessageContent) MarshalJSON() ([]byte, error) {
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
func (mj *InputTextMessageContent) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.DisableWebPagePreview != false {
		if mj.DisableWebPagePreview {
			buf.WriteString(`"disable_web_page_preview":true`)
		} else {
			buf.WriteString(`"disable_web_page_preview":false`)
		}
		buf.WriteByte(',')
	}
	if mj.MessageText != nil {
		buf.WriteString(`"message_text":`)
		fflib.WriteJsonString(buf, string(*mj.MessageText))
	} else {
		buf.WriteString(`"message_text":null`)
	}
	buf.WriteByte(',')
	if len(mj.ParseMode) != 0 {
		buf.WriteString(`"parse_mode":`)
		fflib.WriteJsonString(buf, string(mj.ParseMode))
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_InputTextMessageContentbase = iota
	ffj_t_InputTextMessageContentno_such_key

	ffj_t_InputTextMessageContent_DisableWebPagePreview

	ffj_t_InputTextMessageContent_MessageText

	ffj_t_InputTextMessageContent_ParseMode
)

var ffj_key_InputTextMessageContent_DisableWebPagePreview = []byte("disable_web_page_preview")

var ffj_key_InputTextMessageContent_MessageText = []byte("message_text")

var ffj_key_InputTextMessageContent_ParseMode = []byte("parse_mode")

func (uj *InputTextMessageContent) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *InputTextMessageContent) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_InputTextMessageContentbase
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
				currentKey = ffj_t_InputTextMessageContentno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffj_key_InputTextMessageContent_DisableWebPagePreview, kn) {
						currentKey = ffj_t_InputTextMessageContent_DisableWebPagePreview
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_InputTextMessageContent_MessageText, kn) {
						currentKey = ffj_t_InputTextMessageContent_MessageText
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_InputTextMessageContent_ParseMode, kn) {
						currentKey = ffj_t_InputTextMessageContent_ParseMode
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_InputTextMessageContent_ParseMode, kn) {
					currentKey = ffj_t_InputTextMessageContent_ParseMode
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InputTextMessageContent_MessageText, kn) {
					currentKey = ffj_t_InputTextMessageContent_MessageText
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InputTextMessageContent_DisableWebPagePreview, kn) {
					currentKey = ffj_t_InputTextMessageContent_DisableWebPagePreview
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_InputTextMessageContentno_such_key
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

				case ffj_t_InputTextMessageContent_DisableWebPagePreview:
					goto handle_DisableWebPagePreview

				case ffj_t_InputTextMessageContent_MessageText:
					goto handle_MessageText

				case ffj_t_InputTextMessageContent_ParseMode:
					goto handle_ParseMode

				case ffj_t_InputTextMessageContentno_such_key:
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

handle_DisableWebPagePreview:

	/* handler: uj.DisableWebPagePreview type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.DisableWebPagePreview = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.DisableWebPagePreview = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MessageText:

	/* handler: uj.MessageText type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.MessageText = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.MessageText = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ParseMode:

	/* handler: uj.ParseMode type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ParseMode = string(string(outBuf))

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
