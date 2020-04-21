package httputil

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"

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

// Compress use by zlib
func ZlibCompress(byteArray []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)

	w.Write(byteArray)
	w.Close()

	return b.Bytes()
}

// Decompress use by zlib
func ZlibUnCompress(byteArray []byte) []byte {

	// read byteArray
	b := bytes.NewReader(byteArray)

	// read into zlib
	z, err := zlib.NewReader(b)

	if err != nil {
		panic(err)
	}

	var r bytes.Buffer

	io.Copy(&r, z)

	return r.Bytes()
}

// Decompress use by zlib from io.ReadCloser
func ZlibUnCompressReaderCloser(r io.ReadCloser) []byte {

	byteArray, err := ioutil.ReadAll(r)

	if err != nil {
		panic(err)
	}

	return ZlibUnCompress(byteArray)
}
