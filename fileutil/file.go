package fileutil

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// operate file should be only one way
type FileMgr struct {
	l sync.RWMutex
}

// make object
var mFileMgr = new(FileMgr)

// check file path is file or dir
func (f *FileMgr) isFileOrDir(fileName string, decideDir bool) bool {

	f.l.RLock()
	defer f.l.RUnlock()

	fileInfo, err := os.Stat(fileName)

	if err != nil {
		return false
	}

	isDir := fileInfo.IsDir()

	if decideDir == true {
		return isDir
	}

	// is dir should be false
	// not is dir should be true
	return !isDir
}

// write date to file ,by not same auth
func (f *FileMgr) write(fileName, d string, auth int) (count int, err error) {

	f.l.Lock()
	defer f.l.Unlock()

	var file *os.File

	if f.IsDirOrFileExist(fileName) == false {
		file, err = os.Create(fileName)

		if err != nil {
			return
		}
	} else {
		file, err = os.OpenFile(fileName, auth, 0666)
	}

	defer file.Close()

	return io.WriteString(file, d)
}

// check char exist file
func (f *FileMgr) CheckFileContainsChar(filename, s string) bool {
	data := f.ReadFile(filename)
	if len(data) > 0 {
		return strings.LastIndex(data, s) > 0
	}

	return false
}

func (f *FileMgr) ReadFile(fileName string) string {
	r, err := ioutil.ReadFile(fileName)
	if err != nil {
		return ""
	}

	return string(r)
}

// write data to file
func (f *FileMgr) WriteFile(fileName, d string) (int, error) {
	return f.write(fileName, d, os.O_CREATE|os.O_WRONLY)
}

// write data to file by append modern
func (f *FileMgr) WriteFileAppend(fileName, d string) (count int, err error) {

	return f.write(fileName, d, os.O_APPEND|os.O_WRONLY)
}

// create file
func (f *FileMgr) CreateFile(path string) bool {

	f.l.Lock()
	defer f.l.Unlock()

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	defer file.Close()

	if err != nil {
		return false
	}

	return true
}

// build dir
func (f *FileMgr) CreateDir(path string) bool {

	f.l.Lock()
	defer f.l.Unlock()

	if f.IsDirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)

		if err != nil {
			return false
		}
	}

	return true
}

// create dir
func (f *FileMgr) GenerateDir(path string) (string, error) {

	if len(path) == 0 {
		return "", errors.New("create dir is fail")
	}

	last := path[len(path)-1:]

	if !strings.EqualFold(last, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}

	if f.IsDir(path) {
		if f.CreateDir(path) {
			return path, nil
		}

		return "", errors.New(path + ",create failed")
	}

	return path, nil
}

// check exist path
func (f *FileMgr) IsDirOrFileExist(path string) bool {

	f.l.RLock()
	defer f.l.RUnlock()

	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// check dir path is right?
func (f *FileMgr) IsDir(path string) bool {
	return f.isFileOrDir(path, true)
}

// check file name  is file?
func (f *FileMgr) IsFile(fileName string) bool {
	return f.isFileOrDir(fileName, false)
}
