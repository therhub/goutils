package logutil

import (
	"fmt"
	"os"
	"time"
)

// write log to file
// m is module Name,
// s is content
// to get log
func write(m, lt, s string) error {

	// get current time
	t := time.Now()

	// get file name
	f := fmt.Sprintf("./log%s.log", t.Format("2006010215"))

	// get log format
	out := fmt.Sprintf("\r\nid:%s\r\ntimestamp:%v;\r\nmodule:%s;\r\nlogtype:%s;\r\ntime:%v;\r\ncontent:%v;\r\n", t.Unix(), m, lt, t.Format("2006-01-02 15:04:05"), s)

	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	// close file
	defer file.Close()

	file.WriteString(out)

	return nil
}

// set log path
func SetLogPath(p string) {

}

// system log
func SystemLog(lt, s string) error {
	return write("sys", lt, s)
}
