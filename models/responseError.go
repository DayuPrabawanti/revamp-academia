package models

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (re *ResponseError) Error() string {
	return re.Message
}
