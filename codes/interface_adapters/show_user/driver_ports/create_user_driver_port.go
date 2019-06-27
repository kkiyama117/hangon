package driver_ports

type APIOutput interface {
	ShowUser(user []byte) error
}
