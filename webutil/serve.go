package webutil

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Serve struct {
	mu      sync.RWMutex
	handles map[string]Handle
}

type Handler func(w http.ResponseWriter, r *http.Request)

type Handle struct {
	fn     Handler
	method string
}

var (
	ErrNotFound = `{"error": "not found"}`
)

func NewServe() *Serve {
	return &Serve{
		handles: make(map[string]Handle),
	}
}

func (s *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if h, ok := s.handles[r.URL.Path]; ok {
		if h.method == r.Method {
			h.fn(w, r)
			return
		}
	}

	ErrStr(w, http.StatusNotFound, ErrNotFound)
}

func (s *Serve) POST(path string, handler Handler) *Serve {
	s.addH(path, http.MethodPost, handler)
	return s
}

func (s *Serve) GET(path string, handler Handler) *Serve {
	s.addH(path, http.MethodGet, handler)
	return s
}

func (s *Serve) addH(path, method string, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handles[path] = Handle{
		fn:     h,
		method: method,
	}
}

func (s *Serve) Start(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), s)
}

func Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
