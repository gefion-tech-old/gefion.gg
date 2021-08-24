package model

// Повторяющиеся ошибки
var UTIL__ERROR string = "syntax error"
var UNDEFINED_UTIL__ERROR string = "Unknown error"

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// В результате неудачи в любой команде возвращается объект структуры `Error`,
// `{}` пустой объект как признак удачной операции.
type Response struct {
	Response *Error `json:"error,omitempty"`
}

// Создание ответа
func CreateResponse(typeValue string, messageValue string) Response {
	var res Response
	if len(typeValue) == 0 && len(messageValue) == 0 {
		res = Response{
			Response: nil,
		}
	} else {
		res = Response{
			Response: &Error{
				Type:    typeValue,
				Message: messageValue,
			},
		}
	}
	return res
}
