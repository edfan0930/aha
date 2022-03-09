package response

type (
	// Response response struct
	Response struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
)

//Success
func Success() *Response {
	return &Response{
		Message: "success",
	}
}

//HasData
func HasData(data interface{}) *Response {
	return &Response{
		Data:    data,
		Message: "success",
	}
}

//Error
func Error(message string) *Response {
	return &Response{
		Message: message,
	}
}
