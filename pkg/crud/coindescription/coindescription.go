package coindescription

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/projectinfomgr"
	"github.com/NpoolPlatform/project-info-manager/pkg/db"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coindescription"

	"github.com/google/uuid"
)

type CoinDescription struct {
	*db.Entity
}

func New(ctx context.Context, tx *ent.Tx) (*CoinDescription, error) {
	e, err := db.NewEntity(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("fail create entity: %v", err)
	}

	return &CoinDescription{
		Entity: e,
	}, nil
}

func (s *CoinDescription) rowToObject(row *ent.CoinDescription) *npool.CoinDescription {
	return &npool.CoinDescription{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Title:      row.Title,
		Message:    row.Message,
		UsedFor:    row.UsedFor,
	}
}

func (s *CoinDescription) Create(ctx context.Context, in *npool.CoinDescription) (*npool.CoinDescription, error) {
	var info *ent.CoinDescription
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinDescription.Create().
			SetCoinTypeID(uuid.MustParse(in.CoinTypeID)).
			SetAppID(uuid.MustParse(in.AppID)).
			SetTitle(in.Title).
			SetMessage(in.Message).
			SetUsedFor(in.UsedFor).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create Description: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinDescription) CreateBulk(ctx context.Context, in []*npool.CoinDescription) ([]*npool.CoinDescription, error) {
	rows := []*ent.CoinDescription{}
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		bulk := make([]*ent.CoinDescriptionCreate, len(in))
		for i, info := range in {
			bulk[i] = s.Tx.CoinDescription.Create().
				SetCoinTypeID(uuid.MustParse(info.CoinTypeID)).
				SetAppID(uuid.MustParse(info.AppID)).
				SetTitle(info.Title).
				SetMessage(info.Message).
				SetUsedFor(info.UsedFor)
		}
		rows, err = s.Tx.CoinDescription.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create Descriptions: %v", err)
	}

	infos := []*npool.CoinDescription{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, nil
}

func (s *CoinDescription) Row(ctx context.Context, id uuid.UUID) (*npool.CoinDescription, error) {
	var info *ent.CoinDescription
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinDescription.Query().Where(coindescription.ID(id)).Only(_ctx)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("fail get Description: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinDescription) Update(ctx context.Context, in *npool.CoinDescription) (*npool.CoinDescription, error) {
	var info *ent.CoinDescription
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinDescription.UpdateOneID(uuid.MustParse(in.GetID())).
			SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID())).
			SetTitle(in.GetTitle()).
			SetMessage(in.GetMessage()).
			SetUsedFor(in.GetUsedFor()).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail update Description: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinDescription) Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.CoinDescription, int, error) {
	rows := []*ent.CoinDescription{}
	var total int

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail count Description: %v", err)
		}

		rows, err = stm.Order(ent.Desc(coindescription.FieldUpdateAt)).Offset(offset).Limit(limit).All(_ctx)
		if err != nil {
			return fmt.Errorf("fail query Description: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get Description: %v", err)
	}

	infos := []*npool.CoinDescription{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, total, nil
}

func (s *CoinDescription) queryFromConds(conds cruder.Conds) (*ent.CoinDescriptionQuery, error) {
	stm := s.Tx.CoinDescription.Query()
	for k, v := range conds {
		switch k {
		case coindescription.FieldID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(coindescription.ID(id))
		case coindescription.FieldAppID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid AppID: %v", err)
			}
			stm = stm.Where(coindescription.AppID(id))
		case coindescription.FieldCoinTypeID:
			cointypeid, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid cointypeid: %v", err)
			}
			stm = stm.Where(coindescription.CoinTypeID(cointypeid))
		case coindescription.FieldUsedFor:
			usedfor, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid UsedFor: %v", err)
			}
			stm = stm.Where(coindescription.UsedFor(usedfor))
		case coindescription.FieldMessage:
			message, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Message: %v", err)
			}
			stm = stm.Where(coindescription.Message(message))
		case coindescription.FieldTitle:
			title, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(coindescription.Title(title))
		default:
			return nil, fmt.Errorf("invalid CoinDescription field")
		}
	}

	return stm, nil
}

func (s *CoinDescription) RowOnly(ctx context.Context, conds cruder.Conds) (*npool.CoinDescription, error) {
	var info *ent.CoinDescription

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query Description: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get Description: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinDescription) Count(ctx context.Context, conds cruder.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail check Descriptions: %v", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count Descriptions: %v", err)
	}

	return uint32(total), nil
}

func (s *CoinDescription) Delete(ctx context.Context, id uuid.UUID) (*npool.CoinDescription, error) {
	var info *ent.CoinDescription
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinDescription.UpdateOneID(id).
			SetDeleteAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete Description: %v", err)
	}

	return s.rowToObject(info), nil
}
