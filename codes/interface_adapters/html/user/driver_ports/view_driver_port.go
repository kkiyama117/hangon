package driver_ports

type APIOutput interface {
	ShowUser(user []byte) error
	ShowUsers(users []byte) error
}
