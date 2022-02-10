package helper

import "strings"

// response is used for static json shape
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"json"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// emppty object is used hen data doesnt want to be null on json
type EmptyObj struct{}

// BuildResponse menthod is to inject data value to dynamic succeess response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	return res

}

// BuildErrorResponse menthod is to inject data value to dynamic succeess response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}

	return res
}
