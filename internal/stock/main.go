package stock

import (
	"context"
	"github.com/spf13/viper"
	"github.com/yanlong-l/gorder/common/genproto/stockpb"
	"github.com/yanlong-l/gorder/common/server"
	"github.com/yanlong-l/gorder/stock/ports"
	"github.com/yanlong-l/gorder/stock/service"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	app := service.NewApplication(ctx)
	defer cancel()
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(app)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// 暂时不用
	default:
		panic("unexpected server type")
	}
}
