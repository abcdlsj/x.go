package logger

type Level uint8

const (
	Unknown Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

var (
	DEBUG_PREFIX = "[DEBUG]"
	INFO_PREFIX  = blue("[INFO]")
	WARN_PREFIX  = red("[WARN]")
	ERROR_PREFIX = red("[ERROR]")
	FATAL_PREFIX = red("[FATAL]")
)

func (l Level) String() string {
	switch l {
	case Debug:
		return DEBUG_PREFIX
	case Info:
		return INFO_PREFIX
	case Error:
		return ERROR_PREFIX
	case Warn:
		return WARN_PREFIX
	case Fatal:
		return FATAL_PREFIX
	default:
		return ""
	}
}

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func blue(s string) string {
	return "\033[34m" + s + "\033[0m"
}
