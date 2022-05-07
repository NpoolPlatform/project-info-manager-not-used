package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/project-info-manager"
	crud "github.com/NpoolPlatform/project-info-manager/pkg/crud/description"
	ccoin "github.com/NpoolPlatform/project-info-manager/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCoinDescription ..
func (s *Server) GetCoinDescription(ctx context.Context, in *npool.GetCoinDescriptionRequest) (*npool.GetCoinDescriptionResponse, error) {
	if in.GetCoinTypeID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check CoinTypeID is empty")
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse CoinTypeID: %s invalid", in.GetCoinTypeID())
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, coinID)
	if err != nil {
		logger.Sugar().Errorf("fail get stock: %v", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	des := &npool.CoinDescriptionInfo{
		ID:         info.GetID(),
		CoinTypeID: info.GetCoinTypeID(),
		Title:      info.GetTitle(),
		Message:    info.GetMessage(),
		UsedFor:    info.GetUsedFor(),
		CreatedAt:  info.GetCreatedAt(),
		UpdatedAt:  info.GetUpdatedAt(),
	}
	return &npool.GetCoinDescriptionResponse{
		Total: 1,
		Infos: []*npool.CoinDescriptionInfo{
			des,
		},
	}, nil
}
