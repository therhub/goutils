package stringutil

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/therhub/goutils/logutil"
)

// string reverse
func Reverse(s string) string {

	if s == "" {
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
		logutil.ErrorLog(err)
	}

	return string(r)
}

// add `` Quotes
func AddQuote(str string) string {
	return "`" + str + "`"
}

// remove ` Quotes
func CleanQuote(str string) string {
	return strings.Replace(str, "`", "", -1)
}
