package xconfig

import "errors"

var (
	ErrNotFound = errors.New("key not found")
)

type EventType uint8

const (
	Unknown EventType = iota
	Create
	Update
	Delete
)
