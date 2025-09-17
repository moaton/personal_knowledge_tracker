package types

type State string

const (
	StateResourceCreate State = "resource_create"
	StateResourceList   State = "resource_list"
	StateResourceDelete State = "resource_delete"
)

func (s State) String() string {
	return string(s)
}
