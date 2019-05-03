package web

import (
	"net/http"

	"pumpkin/external_interfaces/common"
)

type output struct {
	writer http.ResponseWriter
}

func NewOutput(w http.ResponseWriter) common.Output {
	return &output{w}
}

func (op *output) Push(data interface{}) error {
	_, err := op.writer.Write(data.([]byte))
	return err
}
