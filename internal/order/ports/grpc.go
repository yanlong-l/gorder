package ports

import (
	"context"
	"github.com/yanlong-l/gorder/common/genproto/orderpb"
	"github.com/yanlong-l/gorder/order/app"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	app app.Application
}

func (s *GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) UpdateOrder(ctx context.Context, order *orderpb.Order) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{
		app: app,
	}
}
