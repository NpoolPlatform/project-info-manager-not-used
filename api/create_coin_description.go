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

// CreateCoinDescription ..
func (s *Server) CreateCoinDescription(ctx context.Context, in *npool.CreateCoinDescriptionRequest) (*npool.CreateCoinDescriptionResponse, error) {
	if in.GetTitle() == "" {
		logger.Sugar().Error("CreateCoinDescription check Title is empty")
		return nil, status.Error(codes.InvalidArgument, "Title empty")
	}

	if in.GetMessage() == "" {
		logger.Sugar().Error("CreateCoinDescription check Message is empty")
		return nil, status.Error(codes.InvalidArgument, "Message empty")
	}

	if in.GetUsedFor() == "" {
		logger.Sugar().Error("CreateCoinDescription check UseFor is empty")
		return nil, status.Error(codes.InvalidArgument, "UseFor empty")
	}

	if in.GetCoinTypeID() == "" {
		logger.Sugar().Error("CreateCoinDescription check CoinTypeID is empty")
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID empty")
	}

	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription parse CoinTypeID: %s invalid", in.GetCoinTypeID())
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Create(ctx, &npool.CoinDescriptionInfo{
		CoinTypeID: coinID.String(),
		Title:      in.GetTitle(),
		Message:    in.GetMessage(),
		UsedFor:    in.GetUsedFor(),
	})
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription call CreateCoinDescription error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinDescriptionResponse{
		Info: &npool.CoinDescriptionInfo{
			ID:         info.GetID(),
			CoinTypeID: info.GetCoinTypeID(),
			Title:      info.Title,
			Message:    info.Message,
			UsedFor:    info.UsedFor,
			CreatedAt:  info.CreatedAt,
			UpdatedAt:  info.UpdatedAt,
		},
	}, nil
}
