package server

type Server interface {
	StartServer(serviceVersion string) error
}
