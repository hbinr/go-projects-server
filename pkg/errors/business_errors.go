package errors

type BusinessError struct {
	Code   int
	String string
}

func NewError(code int, msg ...string) *BusinessError {
	var m string
	if len(msg) > 0 {
		m = msg[0]
	}
	return &BusinessError{Code: code, String: m}
}

func (b *BusinessError) Error() string {
	return b.String
}
func (b *BusinessError) GetCode() int {
	return b.Code
}
