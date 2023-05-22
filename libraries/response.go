package libraries

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// functioon to return http response ok
func ResponseOK(data interface{}) Response {
	return Response{
		Status: "success",
		Data:   data,
	}
}

// function to return http response fail
func ResponseFail(data interface{}) Response {
	return Response{
		Status: "fail",
		Data:   data,
	}
}

// function to return http response error
func ResponseError(data interface{}) Response {
	errorMessage := "Internal Server Error"

	if data != nil {
		errorMessage = data.(string)
	}
	return Response{
		Status: "error",
		Data: map[string]interface{}{
			"error": errorMessage,
		},
	}
}
