package model

var UTIL__ERROR string = "syntax error"
var GIT__ERROR string = "git error"
var UNDEFINED_UTIL__ERROR string = "unknown error"

type ErrorItem struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// В результате неудачи в любой команде возвращается объект структуры `Error`,
// `{}` пустой объект как признак удачной операции.
type Error struct {
	Error *ErrorItem `json:"error,omitempty"`
}

// Создание ответа
func CreateResponse(typeValue string, messageValue string) Error {
	var res Error
	if len(typeValue) == 0 && len(messageValue) == 0 {
		res = Error{
			Error: nil,
		}
	} else {
		res = Error{
			Error: &ErrorItem{
				Type:    typeValue,
				Message: messageValue,
			},
		}
	}
	return res
}
