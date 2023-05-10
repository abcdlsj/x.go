package config

import "errors"

var (
	ErrNotFound = errors.New("key not found")
)

type EventType uint8

const (
	Create EventType = iota
	Update
	Delete

	Unknown EventType = 255
)
