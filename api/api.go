package api

import (
	"context"

	projinfo "github.com/NpoolPlatform/message/npool/project-info-manager"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	projinfo.UnimplementedProjectInfoManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	projinfo.RegisterProjectInfoManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return projinfo.RegisterProjectInfoManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
