//nolint
package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/projectinfomgr"
	crud "github.com/NpoolPlatform/project-info-manager/pkg/crud/coinproductinfo"
	constant "github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coinproductinfo"
	ccoin "github.com/NpoolPlatform/project-info-manager/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkFeildsInCoinProdInfo(info *npool.CoinProductInfo) error {
	if info.GetProductPage() == "" {
		logger.Sugar().Error("check ProductPage is empty")
		return status.Error(codes.InvalidArgument, "ProductPage empty")
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

func (s *Server) CreateCoinProductInfo(ctx context.Context, in *npool.CreateCoinProductInfoRequest) (*npool.CreateCoinProductInfoResponse, error) {
	info := in.GetInfo()
	err := checkFeildsInCoinProdInfo(info)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateCoinProductInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	desc, err := schema.Create(ctx, &npool.CoinProductInfo{
		CoinTypeID:  info.GetCoinTypeID(),
		AppID:       info.GetAppID(),
		ProductPage: info.GetProductPage(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail create CoinProductInfo error %v", err)
		return &npool.CreateCoinProductInfoResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinProductInfoResponse{
		Info: desc,
	}, nil
}

func (s *Server) CreateCoinProductInfos(ctx context.Context, in *npool.CreateCoinProductInfosRequest) (*npool.CreateCoinProductInfosResponse, error) {
	_, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse AppID: %s invalid", in.GetAppID())
		return nil, status.Error(codes.InvalidArgument, "AppID invalid")
	}

	for _, info := range in.Infos {
		err := checkFeildsInCoinProdInfo(info)
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

	infos := []*npool.CoinProductInfo{}
	for _, info := range in.Infos {
		infos = append(infos, &npool.CoinProductInfo{
			CoinTypeID:  info.GetCoinTypeID(),
			AppID:       in.GetAppID(),
			ProductPage: info.GetProductPage(),
		})
	}
	descs, err := schema.CreateBulk(ctx, infos)
	if err != nil {
		logger.Sugar().Errorf("fail create CoinProductInfos error %v", err)
		return &npool.CreateCoinProductInfosResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinProductInfosResponse{
		Infos: descs,
	}, nil
}

func (s *Server) CreateAppCoinProductInfo(ctx context.Context, in *npool.CreateAppCoinProductInfoRequest) (*npool.CreateAppCoinProductInfoResponse, error) {
	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	info := in.GetInfo()
	err = checkFeildsInCoinProdInfo(info)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateAppCoinProductInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	info.AppID = in.GetTargetAppID()

	desc, err := schema.Create(ctx, info)
	if err != nil {
		logger.Sugar().Errorf("fail create CoinProductInfo error %v", err)
		return &npool.CreateAppCoinProductInfoResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateAppCoinProductInfoResponse{
		Info: desc,
	}, nil
}

func (s *Server) CreateAppCoinProductInfos(ctx context.Context, in *npool.CreateAppCoinProductInfosRequest) (*npool.CreateAppCoinProductInfosResponse, error) {
	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	for _, info := range in.Infos {
		err := checkFeildsInCoinProdInfo(info)
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

	infos := []*npool.CoinProductInfo{}
	for _, info := range in.Infos {
		infos = append(infos, &npool.CoinProductInfo{
			CoinTypeID:  info.GetCoinTypeID(),
			AppID:       in.GetTargetAppID(),
			ProductPage: info.GetProductPage(),
		})
	}
	descs, err := schema.CreateBulk(ctx, infos)
	if err != nil {
		logger.Sugar().Errorf("fail create CoinProductInfos error %v", err)
		return &npool.CreateAppCoinProductInfosResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateAppCoinProductInfosResponse{Infos: descs}, nil
}

func (s *Server) UpdateCoinProductInfo(ctx context.Context, in *npool.UpdateCoinProductInfoRequest) (*npool.UpdateCoinProductInfoResponse, error) {
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
		return &npool.UpdateCoinProductInfoResponse{}, status.Error(codes.Internal, err.Error())
	}
	updateInfo, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update CoinProductInfo: %v", err)
		return &npool.UpdateCoinProductInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinProductInfoResponse{
		Info: updateInfo,
	}, nil
}

func coinProductInfoCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
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
		case constant.FieldProductPage:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		default:
			return nil, fmt.Errorf("invalid CoinProductInfo field")
		}
	}

	return newConds, nil
}

func (s *Server) GetCoinProductInfo(ctx context.Context, in *npool.GetCoinProductInfoRequest) (*npool.GetCoinProductInfoResponse, error) {
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
		return &npool.GetCoinProductInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, uuid.MustParse(in.GetID()))
	if err != nil {
		logger.Sugar().Errorf("fail get CoinProductInfo: %v", err)
		return &npool.GetCoinProductInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinProductInfoResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinProductInfos(ctx context.Context, in *npool.GetCoinProductInfosRequest) (*npool.GetCoinProductInfosResponse, error) {
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

	newConds, err := coinProductInfoCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds: %v", err)
		return &npool.GetCoinProductInfosResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID())

	infos, total, err := schema.Rows(ctx, newConds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get CoinProductInfos: %v", err)
		return &npool.GetCoinProductInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinProductInfosResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) GetCoinProductInfoOnly(ctx context.Context, in *npool.GetCoinProductInfoOnlyRequest) (*npool.GetCoinProductInfoOnlyResponse, error) {
	_, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse AppID: %s invalid", in.GetAppID())
		return nil, status.Error(codes.InvalidArgument, "AppID invalid")
	}

	newConds, err := coinProductInfoCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds fields: %v", err)
		return &npool.GetCoinProductInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetAppID())

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCoinProductInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get CoinProductInfo: %v", err)
		return &npool.GetCoinProductInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinProductInfoOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppCoinProductInfos(ctx context.Context, in *npool.GetAppCoinProductInfosRequest) (*npool.GetAppCoinProductInfosResponse, error) {
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

	newConds, err := coinProductInfoCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds: %v", err)
		return &npool.GetAppCoinProductInfosResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetTargetAppID())

	infos, total, err := schema.Rows(ctx, newConds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get CoinProductInfos: %v", err)
		return &npool.GetAppCoinProductInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinProductInfosResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) GetAppCoinProductInfoOnly(ctx context.Context, in *npool.GetAppCoinProductInfoOnlyRequest) (*npool.GetAppCoinProductInfoOnlyResponse, error) {
	_, err := uuid.Parse(in.GetTargetAppID())
	if err != nil {
		logger.Sugar().Errorf("parse TargetAppID: %s invalid", in.GetTargetAppID())
		return nil, status.Error(codes.InvalidArgument, "TargetAppID invalid")
	}

	newConds, err := coinProductInfoCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds : %v", err)
		return &npool.GetAppCoinProductInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	newConds.WithCond(constant.FieldAppID, cruder.EQ, in.GetTargetAppID())

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetAppCoinProductInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get CoinProductInfo: %v", err)
		return &npool.GetAppCoinProductInfoOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppCoinProductInfoOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) DeleteCoinProductInfo(ctx context.Context, in *npool.DeleteCoinProductInfoRequest) (*npool.DeleteCoinProductInfoResponse, error) {
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
		logger.Sugar().Errorf("delete CoinProductInfo: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinProductInfoResponse{
		Info: deletedInfo,
	}, nil
}
