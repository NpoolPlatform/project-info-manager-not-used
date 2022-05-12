package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/project-info-manager"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedProjectInfoManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterProjectInfoManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterProjectInfoManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
