package common

import (
	"net/http"
)

type output struct {
	writer http.ResponseWriter
}

func NewOutput(w http.ResponseWriter) Output {
	return &output{w}
}

func (op *output) Push(data interface{}) error {
	_, err := op.writer.Write(data.([]byte))
	return err
}
