package standard

type StandardReponse struct {
	Code         int     `json:"code"`
	Message      string  `json:"message"`
	ErrorMessage *string `json:"ErrorMessage,omitempty"`
}

func BadRequest(err string) StandardReponse {
	return StandardReponse{
		Code:         400,
		Message:      "Bad Request",
		ErrorMessage: &err,
	}
}

func NotFound(err string) StandardReponse {
	return StandardReponse{
		Code:         404,
		Message:      "Not Found",
		ErrorMessage: &err,
	}
}

func InternalServerError(err string) StandardReponse {
	return StandardReponse{
		Code:         500,
		Message:      "Not Found",
		ErrorMessage: &err,
	}
}

func OkStatus() StandardReponse {
	return StandardReponse{
		Code:    200,
		Message: "OK",
	}
}

func CreateOrUpdateSucccess() StandardReponse {
	return StandardReponse{
		Code:    201,
		Message: "success",
	}
}
