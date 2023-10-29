package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Handler interface {
	Handle() func(Msg) (Msg, error)
	Level() Level
	Name() string
}

type PrettyHandler struct {
	w      io.Writer
	level  Level
	json   bool
	color  bool
	layout string
}

func NewPrettyHandler() *PrettyHandler {
	return &PrettyHandler{
		w:      os.Stdout,
		json:   false,
		color:  false,
		layout: "2006-01-02 15:04:05",
	}
}

func (p *PrettyHandler) Layout(layout string) *PrettyHandler {
	p.layout = layout
	return p
}

func (p *PrettyHandler) W(w io.Writer) *PrettyHandler {
	p.w = w
	return p
}

func (p *PrettyHandler) Json() *PrettyHandler {
	p.json = true
	return p
}

func (p *PrettyHandler) Color() *PrettyHandler {
	p.color = true
	return p
}

func (p *PrettyHandler) SetLevel(level Level) *PrettyHandler {
	p.level = level
	return p
}

func (p *PrettyHandler) Handle() func(Msg) (Msg, error) {
	return func(msg Msg) (Msg, error) {
		t := time.Now().Format(p.layout)
		s := fmt.Sprintf("%v %s %s\n", t, p.Level(), msg)
		if p.json {
			s = fmt.Sprintf(`{"time":"%s","level":"%s","msg":"%s"}\n`, t, p.Level(), msg)
		}
		if _, err := io.WriteString(p.w, s); err != nil {
			return nilMsg, err
		}
		return nilMsg, nil
	}
}

func (p *PrettyHandler) Level() Level {
	return p.level
}

func (p *PrettyHandler) Name() string {
	return "pretty_handler"
}

type MarkHandler struct {
	fields map[string]interface{}
}

func NewMarkHandler(fields map[string]interface{}) *MarkHandler {
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
