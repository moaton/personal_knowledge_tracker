package interfaces

//go:generate mockgen -source=server.go -destination=../../mocks/server.go

type Server interface {
	Notify() <-chan error
	Shutdown() error
}
