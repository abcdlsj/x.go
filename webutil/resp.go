package webutil

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Msg(w http.ResponseWriter, code int, data []byte) {
	w.WriteHeader(code)
	if data != nil {
		w.Write(data)
	}
}

func Ok(w http.ResponseWriter, data []byte) {
	Msg(w, http.StatusOK, data)
}

func OkStr(w http.ResponseWriter, data string) {
	Ok(w, []byte(data))
}

func OkJSON(w http.ResponseWriter, data any) {
	d, err := json.Marshal(data)
	if err != nil {
		Err500(w, err)
		return
	}

	Ok(w, d)
}

type H map[string]any

func OkH(w http.ResponseWriter, data H) {
	OkJSON(w, data)
}

func Err(w http.ResponseWriter, code int, err error) {
	Msg(w, code, []byte(err.Error()))
}

func ErrStr(w http.ResponseWriter, code int, err string) {
	Err(w, code, wrapStr2Err(err))
}

func Err500(w http.ResponseWriter, err error) {
	Err(w, http.StatusInternalServerError, err)
}

func Err400(w http.ResponseWriter, err error) {
	Err(w, http.StatusBadRequest, err)
}

func wrapStr2Err(err string) error {
	return errors.New(err)
}
