package response

type (
	Response struct {
		Data  interface{} `json:"data"`
		Error string      `json:"error"`
	}
)

//Success
func Success() *Response {
	return &Response{
		Error: "nil",
	}
}

//HasData
func HasData(data interface{}) *Response {
	return &Response{
		Data:  data,
		Error: "nil",
	}
}

//Error
func Error(message string) *Response {
	return &Response{
		Error: message,
	}
}
