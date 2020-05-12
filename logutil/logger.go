package logutil

import (
	"fmt"
	"os"
	"time"

	"github.com/therhub/goutils/stringutil"
)

var logpath string

// write log to file
// m is module Name,
// s is content
// to get log
func write(m, lt, s string) error {

	// get current time
	t := time.Now()

	if logpath == "" {
		logpath = "./"
	}

	// get file name
	f := fmt.Sprintf("%slog%s.log", logpath, t.Format("2006010215"))

	// get log format
	out := fmt.Sprintf("\r\n#id:%v\r\n#timestamp:%v;\r\n#module:%s;\r\n#logtype:%s;\r\n#time:%v;\r\n#content:%v;\r\n", stringutil.NewID(), t.Unix(), m, lt, t.Format("2006-01-02 15:04:05"), s)

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
	logpath = p
}

// system log
func SystemLog(lt, s string) error {
	return write("sys", lt, s)
}
