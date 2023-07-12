package codes

import (
	"fmt"
)

type Code int

const (
	CodeUnknown Code = iota
	CodeOK
	CodeBadRequest
	CodeNotFound
	CodeForbidden
	CodeDatabase
	CodeInternal
)

func (c Code) String() string {
	switch c {
	case CodeUnknown:
		return "Unknown"
	case CodeBadRequest:
		return "Bad request"
	case CodeNotFound:
		return "Not found"
	case CodeForbidden:
		return "Forbidden"
	case CodeDatabase:
		return "Database error"
	case CodeInternal:
		return "Internal server error"
	case CodeOK:
		return "OK"
	}
	return fmt.Sprintf("Unknown: %d", c)
}

func (c Code) GoString() string {
	return "errcode.Code[" + c.String() + "]"
}

func (c Code) DetailString(s string) string {
	return c.String() + "[" + s + "]"
}
