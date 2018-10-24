// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

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

func easyjson6601e8cdDecodeGithubComWindghoulShellwithgo(in *jlexer.Lexer, out *Headcommit) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "timestamp":
			out.Timestamp = string(in.String())
		case "added":
			if in.IsNull() {
				in.Skip()
				out.Added = nil
			} else {
				in.Delim('[')
				if out.Added == nil {
					if !in.IsDelim(']') {
						out.Added = make([]string, 0, 4)
					} else {
						out.Added = []string{}
					}
				} else {
					out.Added = (out.Added)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Added = append(out.Added, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "removed":
			if in.IsNull() {
				in.Skip()
				out.Removed = nil
			} else {
				in.Delim('[')
				if out.Removed == nil {
					if !in.IsDelim(']') {
						out.Removed = make([]string, 0, 4)
					} else {
						out.Removed = []string{}
					}
				} else {
					out.Removed = (out.Removed)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Removed = append(out.Removed, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "modified":
			if in.IsNull() {
				in.Skip()
				out.Modified = nil
			} else {
				in.Delim('[')
				if out.Modified == nil {
					if !in.IsDelim(']') {
						out.Modified = make([]string, 0, 4)
					} else {
						out.Modified = []string{}
					}
				} else {
					out.Modified = (out.Modified)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Modified = append(out.Modified, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6601e8cdEncodeGithubComWindghoulShellwithgo(out *jwriter.Writer, in Headcommit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Timestamp))
	}
	{
		const prefix string = ",\"added\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Added == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Added {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"removed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Removed == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Removed {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"modified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Modified == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Modified {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Headcommit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComWindghoulShellwithgo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Headcommit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComWindghoulShellwithgo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Headcommit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComWindghoulShellwithgo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Headcommit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComWindghoulShellwithgo(l, v)
}
func easyjson6601e8cdDecodeGithubComWindghoulShellwithgo1(in *jlexer.Lexer, out *GitJSON) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ref":
			out.Ref = string(in.String())
		case "head_commit":
			(out.Headcommit).UnmarshalEasyJSON(in)
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
func easyjson6601e8cdEncodeGithubComWindghoulShellwithgo1(out *jwriter.Writer, in GitJSON) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ref\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ref))
	}
	{
		const prefix string = ",\"head_commit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Headcommit).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GitJSON) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComWindghoulShellwithgo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GitJSON) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComWindghoulShellwithgo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GitJSON) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComWindghoulShellwithgo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GitJSON) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComWindghoulShellwithgo1(l, v)
}
