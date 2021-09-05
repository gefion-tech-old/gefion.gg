package model

var UTIL__ERROR string = "syntax error"
var GIT__ERROR string = "git error"
var UNDEFINED_UTIL__ERROR string = "unknown error"

type ErrorBody struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// В результате неудачи в любой команде возвращается объект структуры `Error`,
// `{}` пустой объект как признак удачной операции.
type Error struct {
	Error *ErrorBody `json:"error,omitempty"`
}
