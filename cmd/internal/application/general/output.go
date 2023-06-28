package general

import "net/http"

type Output struct {
	Errors     []string    `json:"errors"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"-"`
}

func NewOutput() *Output {
	return &Output{
		Errors:     []string{},
		Data:       nil,
		StatusCode: http.StatusOK,
	}
}

func (r *Output) AddError(err string) {
	r.Errors = append(r.Errors, err)
}

func (r *Output) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

func (r *Output) SetData(data interface{}) {
	r.Data = data
}

func (r *Output) HasErrors() bool {
	return len(r.Errors) > 0
}
