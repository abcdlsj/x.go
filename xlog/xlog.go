package xlog

type Logger struct {
	handlers []Handler
}

func New(options ...Option) *Logger {
	l := &Logger{}
	for _, opt := range options {
		opt(l)
	}
	return l
}

type Option func(*Logger)

func WithHandler(hs ...Handler) Option {
	return func(l *Logger) {
		l.handlers = append(l.handlers, hs...)
	}
}

func (l *Logger) Log(level Level, msg Msg) {
	for _, h := range l.handlers {
		if h.Level() < level {
			continue
		}
		if nmsg, err := h.Handle()(msg); err != nil {
			panic(err)
		} else {
			msg = nmsg
		}
	}
}

func builder(str string, fields ...interface{}) Msg {
	msg := Msg{
		msg:    str,
		fields: make(map[string]interface{}),
	}
	for i := 0; i < len(fields); i += 2 {
		msg.fields[fields[i].(string)] = fields[i+1]
	}
	return msg
}

func (l *Logger) Error(str string, fields ...interface{}) {
	l.Log(Error, builder(str, fields...))
}

func (l *Logger) Info(str string, fields ...interface{}) {
	l.Log(Info, builder(str, fields...))
}

func (l *Logger) Debug(str string, fields ...interface{}) {
	l.Log(Debug, builder(str, fields...))
}

func (l *Logger) Warn(str string, fields ...interface{}) {
	l.Log(Warn, builder(str, fields...))
}
