package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/yanlong-l/gorder/common/config"
	"github.com/yanlong-l/gorder/common/genproto/orderpb"
	"github.com/yanlong-l/gorder/common/server"
	"github.com/yanlong-l/gorder/order/ports"
	"github.com/yanlong-l/gorder/order/service"
	"google.golang.org/grpc"
)

func init() {
	err := config.NewViperConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app := service.NewApplication(ctx)
	go server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, NewHTTPServer(app), ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

	server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(app)
		orderpb.RegisterOrderServiceServer(server, svc)
	})
}
