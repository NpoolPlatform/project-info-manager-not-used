package api

// import (
// 	"context"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
// 	npool "github.com/NpoolPlatform/message/npool/coininfo"
// 	"github.com/NpoolPlatform/project-info-manager/pkg/crud/description"
// 	ccoin "github.com/NpoolPlatform/project-info-manager/pkg/message/const"
// 	"github.com/google/uuid"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func (s *Server) UpdateCoinDescription(ctx context.Context, in *npool.UpdateCoinDescriptionRequest) (*npool.UpdateCoinDescriptionResponse, error) {
// 	if in.GetTitle() == "" {
// 		logger.Sugar().Error("UpdateCoinDescription check Title is empty")
// 		return nil, status.Error(codes.InvalidArgument, "Title empty")
// 	}

// 	if in.GetMessage() == "" {
// 		logger.Sugar().Error("UpdateCoinDescription check Message is empty")
// 		return nil, status.Error(codes.InvalidArgument, "Message empty")
// 	}

// 	if in.GetUsedFor() == "" {
// 		logger.Sugar().Error("UpdateCoinDescription check UseFor is empty")
// 		return nil, status.Error(codes.InvalidArgument, "UseFor empty")
// 	}

// 	if in.GetID() == "" {
// 		logger.Sugar().Error("UpdateCoinDescription check ID is empty")
// 		return nil, status.Error(codes.InvalidArgument, "ID empty")
// 	}

// 	id, err := uuid.Parse(in.GetID())
// 	if err != nil {
// 		logger.Sugar().Errorf("UpdateCoinDescription parse ID: %s invalid", in.GetID())
// 		return nil, status.Error(codes.InvalidArgument, "ID invalid")
// 	}

// 	if in.GetCoinTypeID() == "" {
// 		logger.Sugar().Error("UpdateCoinDescription check CoinTypeID is empty")
// 		return nil, status.Error(codes.InvalidArgument, "CoinTypeID empty")
// 	}

// 	coinID, err := uuid.Parse(in.GetCoinTypeID())
// 	if err != nil {
// 		logger.Sugar().Errorf("UpdateCoinDescription parse CoinTypeID: %s invalid", in.GetCoinTypeID())
// 		return nil, status.Error(codes.InvalidArgument, "CoinTypeID invalid")
// 	}

// 	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
// 	defer cancel()

// 	exist, err := description.ExistCoinDescriptionByID(ctx, id)
// 	if err != nil {
// 		logger.Sugar().Errorf("UpdateCoinDescription call ExistCoinDescriptionByID error %v", err)
// 		return nil, status.Error(codes.Internal, "internal server error")
// 	}

// 	if !exist {
// 		logger.Sugar().Errorf("UpdateCoinDescription call ExistCoinDescriptionByID ID: %v not found", in.GetID())
// 		return nil, status.Errorf(codes.NotFound, "ID: %v not found", in.GetID())
// 	}

// 	coinDesc, err := description.UpdateCoinDescriptionByID(ctx, id, coinID, in.GetUsedFor(), in.GetTitle(), in.GetMessage())
// 	if err != nil {
// 		logger.Sugar().Errorf("UpdateCoinDescription call UpdateCoinDescriptionByID error %v", err)
// 		return nil, status.Error(codes.Internal, "internal server error")
// 	}

// 	return &npool.UpdateCoinDescriptionResponse{
// 		Info: &npool.CoinDescriptionInfo{
// 			ID:         coinDesc.ID.String(),
// 			CoinTypeID: coinDesc.CoinTypeID.String(),
// 			Title:      coinDesc.Title,
// 			Message:    coinDesc.Message,
// 			UsedFor:    coinDesc.UsedFor,
// 			CreatedAt:  coinDesc.CreatedAt,
// 			UpdatedAt:  coinDesc.UpdatedAt,
// 		},
// 	}, nil
// }
