package logger

type Msg struct {
	msg    string
	fields map[string]interface{}
}

var nilMsg = Msg{
	msg:    "",
	fields: nil,
}

func (m Msg) String() string {
	s := m.msg

	for k, v := range m.fields {
		s += " " + k + "=" + v.(string)
	}

	return s
}
