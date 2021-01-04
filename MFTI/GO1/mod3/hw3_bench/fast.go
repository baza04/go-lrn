package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

const filePath string = "./data/users.txt"

type User struct {
	Browsers []string `json:"browsers"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
}

var userPool = sync.Pool{
	New: func() interface{} {
		return &User{}
	},
}

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := bufio.NewScanner(file)

	alreadySeen := make(map[string]bool, 200)
	var count int
	var isAndroid, isMSIE bool

	fmt.Fprintln(out, "found users:")
	for buffer.Scan() {
		count++

		line := buffer.Bytes()
		user := userPool.Get().(*User)
		err := user.UnmarshalJSON(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		isAndroid = false
		isMSIE = false
		for _, browser := range user.Browsers {

			if strings.Contains(browser, "Android") {
				isAndroid = true
				if _, ok := alreadySeen[browser]; !ok {
					alreadySeen[browser] = true
				}
			}
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				if _, ok := alreadySeen[browser]; !ok {
					alreadySeen[browser] = true
				}
			}
		}
		userPool.Put(user)

		if !(isAndroid && isMSIE) {
			continue
		}

		strB := strings.Builder{}
		strB.Grow(12 + len(user.Name) + len(user.Email))
		strB.WriteByte('[')
		strB.WriteString(strconv.Itoa(count - 1))
		strB.WriteByte(']')
		strB.WriteByte(' ')
		strB.WriteString(user.Name)
		strB.WriteByte(' ')
		strB.WriteByte('<')
		strB.WriteString(strings.Replace(user.Email, "@", " [at] ", 1))
		strB.WriteByte('>')
		fmt.Fprintln(out, strB.String())
		strB.Reset()
	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(alreadySeen))
}

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3486653aDecodeMFTIGO1Mod3Hw3Bench(in *jlexer.Lexer, out *User) {
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
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
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

func easyjson3486653aEncodeMFTIGO1Mod3Hw3Bench(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3486653aEncodeMFTIGO1Mod3Hw3Bench(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3486653aEncodeMFTIGO1Mod3Hw3Bench(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3486653aDecodeMFTIGO1Mod3Hw3Bench(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3486653aDecodeMFTIGO1Mod3Hw3Bench(l, v)
}
