package auth

import "github.com/tanveerprottoy/starter-go/stdlib/internal/apigateway/module/user/service"

type Module struct {
	Service *Service
}

func NewModule(userService *service.Service) *Module {
	m := new(Module)
	m.Service = NewService(userService)
	return m
}
