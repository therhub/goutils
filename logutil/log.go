package logutil

type LogType int32

const (
	Debug LogType = iota
	Info
	Warning
	Error
)
