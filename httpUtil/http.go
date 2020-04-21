package httpUtil

import (
	"bufio"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding"
)

// Get encode
func GetEncoding(r io.ReadCloser) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		fmt.Printf("GetEncoding err：%v", err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
