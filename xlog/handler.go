package xlog

import (
	"fmt"
	"io"
	"time"
)

type Handler interface {
	Handle() func(Msg) (Msg, error)
	Level() Level
	Name() string
}

type PrettyHandler struct {
	w      io.Writer
	json   bool
	color  bool
	layout string
}

func NewPrettyHandler(w io.Writer, json bool) Handler {
	return &PrettyHandler{
		w:      w,
		json:   json,
		color:  false,
		layout: "2006-01-02 15:04:05.000000",
	}
}

func (p *PrettyHandler) Handle() func(Msg) (Msg, error) {
	return func(msg Msg) (Msg, error) {
		t := time.Now().Format(p.layout)
		s := fmt.Sprintf("%v %s %s", t, p.Level(), msg)
		if p.json {
			s = fmt.Sprintf(`{"time":"%s","level":"%s","msg":"%s"}`, t, p.Level(), msg)
		}
		if _, err := io.WriteString(p.w, s); err != nil {
			return nilMsg, err
		}
		return nilMsg, nil
	}
}

func (p *PrettyHandler) Level() Level {
	return Error
}

func (p *PrettyHandler) Name() string {
	return "pretty_handler"
}

type MarkHandler struct {
	fields map[string]interface{}
}

func NewMarkHandler(fields map[string]interface{}) Handler {
	return &MarkHandler{
		fields: fields,
	}
}

func (m *MarkHandler) Handle() func(Msg) (Msg, error) {
	return func(msg Msg) (Msg, error) {
		for k, v := range m.fields {
			msg.fields[k] = v
		}
		return msg, nil
	}
}

func (m *MarkHandler) Level() Level {
	return Error
}

func (m *MarkHandler) Name() string {
	return "mark_handler"
}
