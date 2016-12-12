// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/response_message.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *ResponseMessage) MarshalJSON() ([]byte, error) {
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
func (mj *ResponseMessage) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.Description) != 0 {
		buf.WriteString(`"description":`)
		fflib.WriteJsonString(buf, string(mj.Description))
		buf.WriteByte(',')
	}
	if mj.ErrorCode != 0 {
		buf.WriteString(`"error_code":`)
		fflib.FormatBits2(buf, uint64(mj.ErrorCode), 10, mj.ErrorCode < 0)
		buf.WriteByte(',')
	}
	if mj.Ok != false {
		if mj.Ok {
			buf.WriteString(`"ok":true`)
		} else {
			buf.WriteString(`"ok":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Result != nil {
		if true {
			buf.WriteString(`"result":`)

			{

				err = mj.Result.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_ResponseMessagebase = iota
	ffj_t_ResponseMessageno_such_key

	ffj_t_ResponseMessage_Description

	ffj_t_ResponseMessage_ErrorCode

	ffj_t_ResponseMessage_Ok

	ffj_t_ResponseMessage_Result
)

var ffj_key_ResponseMessage_Description = []byte("description")

var ffj_key_ResponseMessage_ErrorCode = []byte("error_code")

var ffj_key_ResponseMessage_Ok = []byte("ok")

var ffj_key_ResponseMessage_Result = []byte("result")

func (uj *ResponseMessage) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *ResponseMessage) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_ResponseMessagebase
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
				currentKey = ffj_t_ResponseMessageno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffj_key_ResponseMessage_Description, kn) {
						currentKey = ffj_t_ResponseMessage_Description
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_ResponseMessage_ErrorCode, kn) {
						currentKey = ffj_t_ResponseMessage_ErrorCode
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffj_key_ResponseMessage_Ok, kn) {
						currentKey = ffj_t_ResponseMessage_Ok
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_ResponseMessage_Result, kn) {
						currentKey = ffj_t_ResponseMessage_Result
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_ResponseMessage_Result, kn) {
					currentKey = ffj_t_ResponseMessage_Result
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_ResponseMessage_Ok, kn) {
					currentKey = ffj_t_ResponseMessage_Ok
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_ResponseMessage_ErrorCode, kn) {
					currentKey = ffj_t_ResponseMessage_ErrorCode
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_ResponseMessage_Description, kn) {
					currentKey = ffj_t_ResponseMessage_Description
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_ResponseMessageno_such_key
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

				case ffj_t_ResponseMessage_Description:
					goto handle_Description

				case ffj_t_ResponseMessage_ErrorCode:
					goto handle_ErrorCode

				case ffj_t_ResponseMessage_Ok:
					goto handle_Ok

				case ffj_t_ResponseMessage_Result:
					goto handle_Result

				case ffj_t_ResponseMessageno_such_key:
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

handle_Description:

	/* handler: uj.Description type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Description = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ErrorCode:

	/* handler: uj.ErrorCode type=int64 kind=int64 quoted=false*/

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

			uj.ErrorCode = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ok:

	/* handler: uj.Ok type=bool kind=bool quoted=false*/

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

				uj.Ok = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.Ok = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Result:

	/* handler: uj.Result type=models.Message kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Result = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Result == nil {
			uj.Result = new(Message)
		}

		err = uj.Result.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
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
