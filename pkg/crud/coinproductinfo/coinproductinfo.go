package coinproductinfo

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/projectinfomgr"
	"github.com/NpoolPlatform/project-info-manager/pkg/db"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coinproductinfo"

	"github.com/google/uuid"
)

type CoinProductInfo struct {
	*db.Entity
}

func New(ctx context.Context, tx *ent.Tx) (*CoinProductInfo, error) {
	e, err := db.NewEntity(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("fail create entity: %v", err)
	}

	return &CoinProductInfo{
		Entity: e,
	}, nil
}

func (s *CoinProductInfo) rowToObject(row *ent.CoinProductInfo) *npool.CoinProductInfo {
	return &npool.CoinProductInfo{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		CoinTypeID:  row.CoinTypeID.String(),
		ProductPage: row.ProductPage,
	}
}

func (s *CoinProductInfo) Create(ctx context.Context, in *npool.CoinProductInfo) (*npool.CoinProductInfo, error) {
	var info *ent.CoinProductInfo
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinProductInfo.Create().
			SetCoinTypeID(uuid.MustParse(in.CoinTypeID)).
			SetAppID(uuid.MustParse(in.AppID)).
			SetProductPage(in.ProductPage).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create ProductInfo: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinProductInfo) CreateBulk(ctx context.Context, in []*npool.CoinProductInfo) ([]*npool.CoinProductInfo, error) {
	rows := []*ent.CoinProductInfo{}
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		bulk := make([]*ent.CoinProductInfoCreate, len(in))
		for i, info := range in {
			bulk[i] = s.Tx.CoinProductInfo.Create().
				SetCoinTypeID(uuid.MustParse(info.CoinTypeID)).
				SetAppID(uuid.MustParse(info.AppID)).
				SetProductPage(info.ProductPage)
		}
		rows, err = s.Tx.CoinProductInfo.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create ProductInfos: %v", err)
	}

	infos := []*npool.CoinProductInfo{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, nil
}

func (s *CoinProductInfo) Row(ctx context.Context, id uuid.UUID) (*npool.CoinProductInfo, error) {
	var info *ent.CoinProductInfo
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinProductInfo.Query().Where(coinproductinfo.ID(id)).Only(_ctx)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("fail get ProductInfo: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinProductInfo) Update(ctx context.Context, in *npool.CoinProductInfo) (*npool.CoinProductInfo, error) {
	var info *ent.CoinProductInfo
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinProductInfo.UpdateOneID(uuid.MustParse(in.GetID())).
			SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID())).
			SetProductPage(in.GetProductPage()).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail update ProductInfo: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinProductInfo) Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.CoinProductInfo, int, error) {
	rows := []*ent.CoinProductInfo{}
	var total int

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail count ProductInfo: %v", err)
		}

		rows, err = stm.Order(ent.Desc(coinproductinfo.FieldUpdateAt)).Offset(offset).Limit(limit).All(_ctx)
		if err != nil {
			return fmt.Errorf("fail query ProductInfo: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get ProductInfo: %v", err)
	}

	infos := []*npool.CoinProductInfo{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, total, nil
}

func (s *CoinProductInfo) queryFromConds(conds cruder.Conds) (*ent.CoinProductInfoQuery, error) {
	stm := s.Tx.CoinProductInfo.Query()
	for k, v := range conds {
		switch k {
		case coinproductinfo.FieldID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(coinproductinfo.ID(id))
		case coinproductinfo.FieldAppID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid AppID: %v", err)
			}
			stm = stm.Where(coinproductinfo.AppID(id))
		case coinproductinfo.FieldCoinTypeID:
			cointypeid, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid CoinTypeID: %v", err)
			}
			stm = stm.Where(coinproductinfo.CoinTypeID(cointypeid))
		case coinproductinfo.FieldProductPage:
			ProductPage, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ProductPage: %v", err)
			}
			stm = stm.Where(coinproductinfo.ProductPage(ProductPage))
		default:
			return nil, fmt.Errorf("invalid CoinProductInfo field")
		}
	}

	return stm, nil
}

func (s *CoinProductInfo) RowOnly(ctx context.Context, conds cruder.Conds) (*npool.CoinProductInfo, error) {
	var info *ent.CoinProductInfo

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query ProductInfo: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ProductInfo: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinProductInfo) Count(ctx context.Context, conds cruder.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail check ProductInfos: %v", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count ProductInfos: %v", err)
	}

	return uint32(total), nil
}

func (s *CoinProductInfo) Delete(ctx context.Context, id uuid.UUID) (*npool.CoinProductInfo, error) {
	var info *ent.CoinProductInfo
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinProductInfo.UpdateOneID(id).
			SetDeleteAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete ProductInfo: %v", err)
	}

	return s.rowToObject(info), nil
}
