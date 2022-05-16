//nolint
package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/projectinfomgr"
	crud "github.com/NpoolPlatform/project-info-manager/pkg/crud/coindescription"
	constant "github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coindescription"
	ccoin "github.com/NpoolPlatform/project-info-manager/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkFeildsInCoinDesc(info *npool.CoinDescription) error {
	if info.GetTitle() == "" {
		logger.Sugar().Error("check Title is empty")
		return status.Error(codes.InvalidArgument, "Title empty")
	}

	if info.GetMessage() == "" {
		logger.Sugar().Error("check Message is empty")
		return status.Error(codes.InvalidArgument, "Message empty")
	}

	if info.GetUsedFor() == "" {
		logger.Sugar().Error("check UsedFor is empty")
		return status.Error(codes.InvalidArgument, "UsedFor empty")
	}

	if info.GetCoinTypeID() == "" {
		logger.Sugar().Error("check CoinTypeID is empty")
		return status.Error(codes.InvalidArgument, "CoinTypeID empty")
	}

	if info.GetAppID() == "" {
		logger.Sugar().Error("check AppID is empty")
		return status.Error(codes.InvalidArgument, "AppID empty")
	}

	_, err := uuid.Parse(info.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("parse CoinTypeID: %s invalid", info.GetCoinTypeID())
		return status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	_, err = uuid.Parse(info.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse AppID: %s invalid", info.GetCoinTypeID())
		return status.Error(codes.InvalidArgument, "AppID invalid")
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
		CoinTypeID: info.GetCoinTypeID(),
		AppID:      info.GetAppID(),
		Title:      info.GetTitle(),
		Message:    info.GetMessage(),
		UsedFor:    info.GetUsedFor(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail create CoinDescription error %v", err)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.Internal, "internal server error")
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
			CoinTypeID: info.GetCoinTypeID(),
			AppID:      in.GetAppID(),
			Title:      info.GetTitle(),
			Message:    info.GetMessage(),
			UsedFor:    info.GetUsedFor(),
		})
	}
	descs, err := schema.CreateBulk(ctx, infos)
	if err != nil {
		logger.Sugar().Errorf("fail create CoinDescriptions error %v", err)
		return &npool.CreateCoinDescriptionsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinDescriptionsResponse{
		Infos: descs,
	}, nil
}

func (s *Server) CreateAppCoinDescription(ctx context.Context, in *npool.CreateAppCoinDescriptionRequest) (*npool.CreateAppCoinDescriptionResponse, error) {
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
		return &npool.CreateAppCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	desc, err := schema.Create(ctx, &npool.CoinDescription{
		CoinTypeID: info.GetCoinTypeID(),
		AppID:      info.GetAppID(),
		Title:      info.GetTitle(),
		Message:    info.GetMessage(),
		UsedFor:    info.GetUsedFor(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail create CoinDescription error %v", err)
		return &npool.CreateAppCoinDescriptionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateAppCoinDescriptionResponse{
		Info: desc,
	}, nil
}

func (s *Server) CreateAppCoinDescriptions(ctx context.Context, in *npool.CreateAppCoinDescriptionsRequest) (*npool.CreateAppCoinDescriptionsResponse, error) {
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
			CoinTypeID: info.GetCoinTypeID(),
			AppID:      in.GetTargetAppID(),
			Title:      info.GetTitle(),
			Message:    info.GetMessage(),
			UsedFor:    info.GetUsedFor(),
		})
	}
	descs, err := schema.CreateBulk(ctx, infos)
	if err != nil {
		logger.Sugar().Errorf("fail create CoinDescriptions error %v", err)
		return &npool.CreateAppCoinDescriptionsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateAppCoinDescriptionsResponse{Infos: descs}, nil
}

func (s *Server) UpdateCoinDescription(ctx context.Context, in *npool.UpdateCoinDescriptionRequest) (*npool.UpdateCoinDescriptionResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorf("parse request AppID: %v invalid", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("parse request ID: %v invalid", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		logger.Sugar().Errorf("parse request CoinTypeID: %v invalid", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}
	updateInfo, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update CoinDescription: %v", err)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinDescriptionResponse{
		Info: updateInfo,
	}, nil
}

func coinDescriptionCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
	newConds := cruder.NewConds()

	for k, v := range conds {
		switch v.Op {
		case cruder.EQ:
		case cruder.GT:
		case cruder.LT:
		case cruder.LIKE:
		default:
			return nil, fmt.Errorf("invalid filter condition op")
		}

		switch k {
		case constant.FieldID:
			fallthrough //nolint
		case constant.FieldAppID:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.FieldCoinTypeID:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.FieldTitle:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.FieldMessage:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.FieldUsedFor:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		default:
			return nil, fmt.Errorf("invalid CoinDescription field")
		}
	}

	return newConds, nil
}

func (s *Server) GetCoinDescription(ctx context.Context, in *npool.GetCoinDescriptionRequest) (*npool.GetCoinDescriptionResponse, error) {
	_, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("parse ID: %s invalid", in.GetID())
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
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
		logger.Sugar().Errorf("fail get CoinDescription: %v", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinDescriptions(ctx context.Context, in *npool.GetCoinDescriptionsRequest) (*npool.GetCoinDescriptionsResponse, error) {
	_, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse AppID: %s invalid", in.GetAppID())
		return nil, status.Error(codes.InvalidArgument, "AppID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	newConds, err := coinDescriptionCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds: %v", err)
		return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID())

	infos, total, err := schema.Rows(ctx, newConds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get CoinDescriptions: %v", err)
		return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionsResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) GetCoinDescriptionOnly(ctx context.Context, in *npool.GetCoinDescriptionOnlyRequest) (*npool.GetCoinDescriptionOnlyResponse, error) {
	_, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse AppID: %s invalid", in.GetAppID())
		return nil, status.Error(codes.InvalidArgument, "AppID invalid")
	}

	newConds, err := coinDescriptionCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds fields: %v", err)
		return &npool.GetCoinDescriptionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID())

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCoinDescriptionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get CoinDescription: %v", err)
		return &npool.GetCoinDescriptionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppCoinDescription(ctx context.Context, in *npool.GetAppCoinDescriptionRequest) (*npool.GetAppCoinDescriptionResponse, error) {
	_, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("parse ID: %s invalid", in.GetID())
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
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
		logger.Sugar().Errorf("fail get CoinDescription: %v", err)
		return &npool.GetAppCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinDescriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppCoinDescriptions(ctx context.Context, in *npool.GetAppCoinDescriptionsRequest) (*npool.GetAppCoinDescriptionsResponse, error) {
	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	newConds, err := coinDescriptionCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds: %v", err)
		return &npool.GetAppCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetTargetAppID())

	infos, total, err := schema.Rows(ctx, newConds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get CoinDescriptions: %v", err)
		return &npool.GetAppCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinDescriptionsResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) GetAppCoinDescriptionOnly(ctx context.Context, in *npool.GetAppCoinDescriptionOnlyRequest) (*npool.GetAppCoinDescriptionOnlyResponse, error) {
	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	newConds, err := coinDescriptionCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds : %v", err)
		return &npool.GetAppCoinDescriptionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetTargetAppID())

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetAppCoinDescriptionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get CoinDescription: %v", err)
		return &npool.GetAppCoinDescriptionOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinDescriptionOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) DeleteCoinDescription(ctx context.Context, in *npool.DeleteCoinDescriptionRequest) (*npool.DeleteCoinDescriptionResponse, error) {
	_, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("parse ID: %s invalid", in.GetID())
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
		logger.Sugar().Errorf("delete CoinDescription: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinDescriptionResponse{
		Info: deletedInfo,
	}, nil
}
