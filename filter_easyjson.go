// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package nostr

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson4d398eaaDecodeGithubComNbdWtfGoNostr(in *jlexer.Lexer, out *Filter) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "IDs":
			if in.IsNull() {
				in.Skip()
				out.IDs = nil
			} else {
				in.Delim('[')
				if out.IDs == nil {
					if !in.IsDelim(']') {
						out.IDs = make([]string, 0, 4)
					} else {
						out.IDs = []string{}
					}
				} else {
					out.IDs = (out.IDs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.IDs = append(out.IDs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Kinds":
			if in.IsNull() {
				in.Skip()
				out.Kinds = nil
			} else {
				in.Delim('[')
				if out.Kinds == nil {
					if !in.IsDelim(']') {
						out.Kinds = make([]int, 0, 8)
					} else {
						out.Kinds = []int{}
					}
				} else {
					out.Kinds = (out.Kinds)[:0]
				}
				for !in.IsDelim(']') {
					var v2 int
					v2 = int(in.Int())
					out.Kinds = append(out.Kinds, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Authors":
			if in.IsNull() {
				in.Skip()
				out.Authors = nil
			} else {
				in.Delim('[')
				if out.Authors == nil {
					if !in.IsDelim(']') {
						out.Authors = make([]string, 0, 4)
					} else {
						out.Authors = []string{}
					}
				} else {
					out.Authors = (out.Authors)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Authors = append(out.Authors, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Tags":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Tags = make(TagMap)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 []string
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						in.Delim('[')
						if v4 == nil {
							if !in.IsDelim(']') {
								v4 = make([]string, 0, 4)
							} else {
								v4 = []string{}
							}
						} else {
							v4 = (v4)[:0]
						}
						for !in.IsDelim(']') {
							var v5 string
							v5 = string(in.String())
							v4 = append(v4, v5)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Tags)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Since":
			out.Since = Timestamp(in.Int64())
		case "Until":
			out.Until = Timestamp(in.Int64())
		case "Limit":
			out.Limit = int(in.Int())
		case "Search":
			out.Search = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4d398eaaEncodeGithubComNbdWtfGoNostr(out *jwriter.Writer, in Filter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"IDs\":"
		out.RawString(prefix[1:])
		if in.IDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.IDs {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Kinds\":"
		out.RawString(prefix)
		if in.Kinds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Kinds {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Authors\":"
		out.RawString(prefix)
		if in.Authors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Authors {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v12First := true
			for v12Name, v12Value := range in.Tags {
				if v12First {
					v12First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v12Name))
				out.RawByte(':')
				if v12Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v13, v14 := range v12Value {
						if v13 > 0 {
							out.RawByte(',')
						}
						out.String(string(v14))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"Since\":"
		out.RawString(prefix)
		out.Int64(int64(in.Since))
	}
	{
		const prefix string = ",\"Until\":"
		out.RawString(prefix)
		out.Int64(int64(in.Until))
	}
	{
		const prefix string = ",\"Limit\":"
		out.RawString(prefix)
		out.Int(int(in.Limit))
	}
	{
		const prefix string = ",\"Search\":"
		out.RawString(prefix)
		out.String(string(in.Search))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Filter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4d398eaaEncodeGithubComNbdWtfGoNostr(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Filter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4d398eaaEncodeGithubComNbdWtfGoNostr(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Filter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4d398eaaDecodeGithubComNbdWtfGoNostr(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Filter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4d398eaaDecodeGithubComNbdWtfGoNostr(l, v)
}
