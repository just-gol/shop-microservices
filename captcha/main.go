package main

import (
	"captcha/handler"
	pb "captcha/proto"
	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
	"go-micro.dev/v5/server"
)

var (
	serviceName = "captcha"
	version     = "0.0.1"
)

func main() {
	registry := consul.NewRegistry()
	// Create service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Registry(registry),
		micro.Server(server.NewServer(
			server.Name(serviceName),
			server.Registry(registry),
			server.Address("0.0.0.0:8088"),
			server.Advertise("127.0.0.1:8088"),
		)),
	)

	// Initialize service
	service.Init()

	// Register handler
	pb.RegisterCaptchaHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}
