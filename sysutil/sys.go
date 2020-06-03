package sysutil

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	Unknown = iota
	Darwin
	Windows
	Linux
)

// get exe excaute dir path
func GetExeRootDir() string {
	r, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		r = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		r = fmt.Sprintf("%s%s", r, string(os.PathSeparator))
	}

	return r
}

// get root path
func GetRootPath(path string) string {

	length := len(path)
	if path[length-1:] == string(os.PathSeparator) {
		path = path[:length-1]
	}

	// path can contain chinese etc.
	// separate path until last path separator
	return subStr(path, 0, strings.LastIndex(path, string(os.PathSeparator)))
}

// substr unicode/utf-8
func subStr(s string, pos, length int) string {
	runes := []rune(s)

	l := pos + length

	runesLength := len(runes)

	if l > runesLength {
		l = runesLength
	}

	return string(runes[pos:l])
}

// clean
func Clean() {
	switch GetOs() {
	case Darwin, Linux:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case Windows:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// get os
func GetOs() int {
	switch runtime.GOOS {
	case "darwin":
		return Darwin
	case "windows":
		return Windows
	case "linux":
		return Linux
	default:
		return Unknown
	}
}
