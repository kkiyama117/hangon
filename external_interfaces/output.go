package external_interfaces

type Output interface {
	Push(interface{}) error
}
