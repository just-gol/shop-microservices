package main

import (
	"rbac/handler"
	pb "rbac/proto/rbacLogin"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
	log "go-micro.dev/v5/logger"
	"go-micro.dev/v5/server"
	"gopkg.in/ini.v1"
)

var (
	serviceName = "rbac"
	version     = "0.0.1"
)

func main() {
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	address := cfg.Section("consul").Key("address").String()
	advertise := cfg.Section("consul").Key("advertise").String()
	registry := consul.NewRegistry()
	// Create service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Registry(registry),
		micro.Server(server.NewServer(
			server.Name(serviceName),
			server.Registry(registry),
			server.Address(address),
			server.Advertise(advertise),
		)),
	)

	// Initialize service
	service.Init()

	// Register handler
	pb.RegisterRbacLoginHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}
