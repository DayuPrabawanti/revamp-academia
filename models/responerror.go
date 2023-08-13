package models

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

// type ResponseError struct {
// 	Message string `json:"message"`
// 	Status  int    `json:"status"`
// }

// func (re *ResponseError) Error() string {
// 	return re.Message
// }
