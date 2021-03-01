package errors

import (
	"strings"
)

const (
	MISSING_PARAM = "missing parameter"
	INVALID_PARAM = "invalid parameter"
	ERR_CONNECTION = "connection error"
	ERR_CONFIG = "configuration error"
	ERR_NO_METRIC = "no metrics"
	ERR_NO_INSTANCE = "no instances"
	ERR_NO_COLLECTOR = "no collectors"
	MATRIX_HASH = "matrix error"
	MATRIX_EMPTY = "empty cache"
	MATRIX_INV_PARAM = "matrix invalid parameter"
	MATRIX_PARSE_STR = "parse numeric value from string"
	API_RESPONSE = "error reading api response"
	API_REQ_REJECTED = "api request rejected"
	ERR_DLOAD = "dynamic load"
	ERR_IMPLEMENT = "implementation error"
	ERR_SCHEDULE = "schedule error"
)

type Error struct {
	class string
	msg string
}

func (e Error) Error() string {
	return e.class + " => " + e.msg
}

func New(class, msg string) Error {
	return Error{class:class, msg:msg}
}

func IsErr(err error, class string) bool {
	// dirty solution, temporarily
	e := strings.Split(err.Error(), " => ")
	if len(e) > 1 {
		return strings.Contains(e[0], class)
	}
	return false
}