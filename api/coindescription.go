package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/project-info-manager"
	crud "github.com/NpoolPlatform/project-info-manager/pkg/crud/coindescription"
	constant "github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coindescription"
	ccoin "github.com/NpoolPlatform/project-info-manager/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkFeildsInCoinDesc(info *npool.CoinDescriptionBase) error {
	if info.GetTitle() == "" {
		logger.Sugar().Error("CreateCoinDescription check Title is empty")
		return status.Error(codes.InvalidArgument, "Title empty")
	}

	if info.GetMessage() == "" {
		logger.Sugar().Error("CreateCoinDescription check Message is empty")
		return status.Error(codes.InvalidArgument, "Message empty")
	}

	if info.GetUsedFor() == "" {
		logger.Sugar().Error("CreateCoinDescription check UseFor is empty")
		return status.Error(codes.InvalidArgument, "UseFor empty")
	}

	if info.GetCoinTypeID() == "" {
		logger.Sugar().Error("CreateCoinDescription check CoinTypeID is empty")
		return status.Error(codes.InvalidArgument, "CoinTypeID empty")
	}

	if info.GetAppID() == "" {
		logger.Sugar().Error("CreateCoinDescription check CoinTypeID is empty")
		return status.Error(codes.InvalidArgument, "CoinTypeID empty")
	}

	_, err := uuid.Parse(info.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription parse CoinTypeID: %s invalid", info.GetCoinTypeID())
		return status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	_, err = uuid.Parse(info.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription parse CoinTypeID: %s invalid", info.GetCoinTypeID())
		return status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}
	return nil
}

func (s *Server) CreateCoinDescription(ctx context.Context, in *npool.CreateCoinDescriptionRequest) (*npool.CreateCoinDescriptionResponse, error) {
	info := in.GetInfo()
	err := checkFeildsInCoinDesc(info)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	desc, err := schema.Create(ctx, &npool.CoinDescription{
		CoinTypeID: info.CoinTypeID,
		AppID:      info.AppID,
		Title:      info.GetTitle(),
		Message:    info.GetMessage(),
		UsedFor:    info.GetUsedFor(),
	})
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription call CreateCoinDescription error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinDescriptionResponse{
		Info: desc,
	}, nil
}

func (s *Server) CreateCoinDescriptions(ctx context.Context, in *npool.CreateCoinDescriptionsRequest) (*npool.CreateCoinDescriptionsResponse, error) {
	for _, info := range in.Infos {
		err := checkFeildsInCoinDesc(info)
		if err != nil {
			return nil, err
		}
	}
	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	infos := []*npool.CoinDescription{}
	for _, info := range in.Infos {
		infos = append(infos, &npool.CoinDescription{
			CoinTypeID: info.CoinTypeID,
			AppID:      info.AppID,
			Title:      info.GetTitle(),
			Message:    info.GetMessage(),
			UsedFor:    info.GetUsedFor(),
		})
	}
	desc, err := schema.CreateBulk(ctx, infos)
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription call CreateCoinDescription error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinDescriptionsResponse{Infos: desc}, nil
}

func (s *Server) UpdateCoinDescription(ctx context.Context, in *npool.UpdateCoinDescriptionRequest) (*npool.UpdateCoinDescriptionResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorf("invalid request app id: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorf("invalid request app id: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("invalid coin_description id: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	updateInfo, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update coin description: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinDescriptionResponse{
		Info: updateInfo,
	}, nil
}

func (s *Server) UpdateAppCoinDescription(ctx context.Context, in *npool.UpdateAppCoinDescriptionRequest) (*npool.UpdateAppCoinDescriptionResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorf("invalid request app id: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorf("invalid request target app id: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("invalid coin_description id: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	updateInfo, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update stock: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppCoinDescriptionResponse{
		Info: updateInfo,
	}, nil
}

func (s *Server) GetCoinDescription(ctx context.Context, in *npool.GetCoinDescriptionRequest) (*npool.GetCoinDescriptionResponse, error) {
	if in.GetAppID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check AppID is empty")
		return nil, status.Error(codes.InvalidArgument, "AppID is empty")
	}

	_, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse GetAppID: %s invalid", in.GetAppID())
		return nil, status.Error(codes.InvalidArgument, "GetAppID invalid")
	}

	_, err = uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse GetID: %s invalid", in.GetID())
		return nil, status.Error(codes.InvalidArgument, "GetID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, uuid.MustParse(in.GetID()))
	if err != nil {
		logger.Sugar().Errorf("fail get description: %v", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppCoinDescription(ctx context.Context, in *npool.GetAppCoinDescriptionRequest) (*npool.GetAppCoinDescriptionResponse, error) {
	if in.GetTargetAppID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check AppID is empty")
		return nil, status.Error(codes.InvalidArgument, "AppID is empty")
	}

	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	_, err = uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse GetID: %s invalid", in.GetID())
		return nil, status.Error(codes.InvalidArgument, "GetID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetAppCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := schema.Row(ctx, uuid.MustParse(in.GetID()))
	if err != nil {
		logger.Sugar().Errorf("fail get description: %v", err)
		return &npool.GetAppCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinDescriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinDescriptions(ctx context.Context, in *npool.GetCoinDescriptionsRequest) (*npool.GetCoinDescriptionsResponse, error) {
	if in.GetAppID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check AppID is empty")
		return nil, status.Error(codes.InvalidArgument, "AppID is empty")
	}

	_, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse AppID: %s invalid", in.GetAppID())
		return nil, status.Error(codes.InvalidArgument, "AppID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	newConds := cruder.NewConds()
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID())
	infos, total, err := schema.Rows(ctx, newConds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get description: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionsResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) GetAppCoinDescriptions(ctx context.Context, in *npool.GetAppCoinDescriptionsRequest) (*npool.GetAppCoinDescriptionsResponse, error) {
	if in.GetTargetAppID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check TargetAppID is empty")
		return nil, status.Error(codes.InvalidArgument, "TargetAppID is empty")
	}

	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	newConds := cruder.NewConds()
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetTargetAppID())
	infos, total, err := schema.Rows(ctx, newConds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get description: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinDescriptionsResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) CountCoinDescriptions(ctx context.Context, in *npool.CountCoinDescriptionsRequest) (*npool.CountCoinDescriptionsResponse, error) {
	if in.GetAppID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check AppID is empty")
		return nil, status.Error(codes.InvalidArgument, "AppID is empty")
	}

	_, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse AppID: %s invalid", in.GetAppID())
		return nil, status.Error(codes.InvalidArgument, "AppID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	newConds := cruder.NewConds()
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID())
	total, err := schema.Count(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get description: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCoinDescriptionsResponse{
		Result: total,
	}, nil
}

func (s *Server) CountAppCoinDescriptions(ctx context.Context, in *npool.CountAppCoinDescriptionsRequest) (*npool.CountAppCoinDescriptionsResponse, error) {
	if in.GetTargetAppID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check TargetAppID is empty")
		return nil, status.Error(codes.InvalidArgument, "TargetAppID is empty")
	}

	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	newConds := cruder.NewConds()
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetTargetAppID())
	total, err := schema.Count(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get description: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAppCoinDescriptionsResponse{
		Result: total,
	}, nil
}

func (s *Server) DeleteAppCoinDescription(ctx context.Context, in *npool.DeleteAppCoinDescriptionRequest) (*npool.DeleteAppCoinDescriptionResponse, error) {
	if in.GetTargetAppID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check TargetAppID is empty")
		return nil, status.Error(codes.InvalidArgument, "TargetAppID is empty")
	}

	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	if in.GetID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check ID is empty")
		return nil, status.Error(codes.InvalidArgument, "ID is empty")
	}

	_, err = uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse ID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	deletedInfo, err := schema.Delete(ctx, uuid.MustParse(in.GetID()))
	if err != nil {
		logger.Sugar().Errorf("delete description: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppCoinDescriptionResponse{
		Info: deletedInfo,
	}, nil
}
