package stringutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

const STRING_EMPTY = ""

// check s is empty?
func IsEmpty(s string) bool {
	return strings.EqualFold(s, STRING_EMPTY)
}

// string reverse
func Reverse(s string) string {

	if IsEmpty(s) {
		return s
	}

	var r string
	var strLen int = len(s)

	for i := 0; i < strLen; i++ {
		r = r + fmt.Sprintf("%c", s[strLen-i])
	}

	return r
}

// char upper
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// convert interface to json string
func ToJson(d interface{}) string {
	r, err := json.Marshal(d)

	if err != nil {
		return STRING_EMPTY
	}

	return string(r)
}

// add `` Quotes
func AddQuote(str string) string {
	return "`" + str + "`"
}

// remove ` Quotes
func CleanQuote(str string) string {
	return strings.Replace(str, "`", STRING_EMPTY, -1)
}

// add comment
func AddToComment(s, suff string) string {
	if IsEmpty(s) {
		return STRING_EMPTY
	}

	return "//" + s + suff
}

// format field
func FormatField(field string, formats []string) string {
	if len(formats) <= 0 {
		return STRING_EMPTY
	}

	var buf bytes.Buffer

	for k := range formats {
		buf.WriteString(fmt.Sprintf(`%s:"%s"`, formats[k], field))
	}

	return "`" + strings.TrimRight(buf.String(), " ") + "`"
}
