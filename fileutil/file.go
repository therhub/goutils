package fileutil

import (
	"errors"
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

// build dir
func (f *FileMgr) CreateDir(path string) bool {
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
