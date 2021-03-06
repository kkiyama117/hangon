package framework_drivers

import (
	"net/http"

	"pumpkin/codes/interface_adapters/html/user/driver_ports"
)

type output struct {
	writer http.ResponseWriter
}

func NewAPIOutput(w http.ResponseWriter) driver_ports.APIOutput {
	return &output{w}
}

func (op *output) ShowUser(data []byte) error {
	_, err := op.writer.Write(data)
	return err
}

func (op *output) ShowUsers(data []byte) error {
	_, err := op.writer.Write(data)
	return err
}
