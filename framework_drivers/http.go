package framework_drivers

import (
	"net/http"

	"pumpkin/interface_adapters/create_user/driver_ports"
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
