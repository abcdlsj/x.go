package xlog

type Level uint8

const (
	Unknown Level = iota
	Debug
	Info
	Warn
	Error
)

func (l Level) String() string {
	switch l {
	case Info:
		return "INFO"
	case Error:
		return "ERROR"
	case Debug:
		return "DEBUG"
	case Warn:
		return "WARN"
	default:
		return "UNKNOWN"
	}
}
