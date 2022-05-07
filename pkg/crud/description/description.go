package description

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/description"

	"github.com/NpoolPlatform/project-info-manager/pkg/db"

	npool "github.com/NpoolPlatform/message/npool/project-info-manager"

	"github.com/google/uuid"
)

type Description struct {
	*db.Entity
}

func New(ctx context.Context, tx *ent.Tx) (*Description, error) {
	e, err := db.NewEntity(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("fail create entity: %v", err)
	}

	return &Description{
		Entity: e,
	}, nil
}

func (s *Description) rowToObject(row *ent.Description) *npool.CoinDescriptionInfo {
	return &npool.CoinDescriptionInfo{
		ID:         row.ID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Title:      row.Title,
		Message:    row.Message,
		UsedFor:    row.UsedFor,
	}
}

func (s *Description) Create(ctx context.Context, in *npool.CoinDescriptionInfo) (*npool.CoinDescriptionInfo, error) {
	var info *ent.Description
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.Description.Create().
			SetCoinTypeID(uuid.MustParse(in.CoinTypeID)).
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

func (s *Description) CreateBulk(ctx context.Context, in []*npool.CoinDescriptionInfo) ([]*npool.CoinDescriptionInfo, error) {
	rows := []*ent.Description{}
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		bulk := make([]*ent.DescriptionCreate, len(in))
		for i, info := range in {
			bulk[i] = s.Tx.Description.Create().
				SetCoinTypeID(uuid.MustParse(info.CoinTypeID)).
				SetTitle(info.Title).
				SetMessage(info.Message).
				SetUsedFor(info.UsedFor)
		}
		rows, err = s.Tx.Description.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create Descriptions: %v", err)
	}

	infos := []*npool.CoinDescriptionInfo{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, nil
}

func (s *Description) Row(ctx context.Context, coinTypeID uuid.UUID) (*npool.CoinDescriptionInfo, error) {
	var info *ent.Description
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.Description.Query().Where(description.CoinTypeID(coinTypeID)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail get stock: %v", err)
	}

	return s.rowToObject(info), nil
}

// func (s *Description) Update(ctx context.Context, in *npool.Description) (*npool.Description, error) {
// 	var info *ent.Description
// 	var err error

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		info, err = s.Tx.Description.UpdateOneID(uuid.MustParse(in.GetID())).
// 			SetTotal(in.GetTotal()).
// 			Save(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("fail update Description: %v", err)
// 	}

// 	return s.rowToObject(info), nil
// }

// func (s *Description) UpdateFields(ctx context.Context, id uuid.UUID, fields cruder.Fields) (*npool.Description, error) {
// 	var info *ent.Description
// 	var err error

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		stm := s.Tx.Description.UpdateOneID(id)
// 		for k, v := range fields {
// 			total, err := cruder.AnyTypeUint32(v)
// 			if err != nil {
// 				return fmt.Errorf("invalid value type: %v", err)
// 			}

// 			switch k {
// 			case constant.DescriptionFieldTotal:
// 				stm = stm.SetTotal(total)
// 			default:
// 				return fmt.Errorf("invalid Description field")
// 			}
// 		}

// 		info, err = stm.Save(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail update Description fields: %v", err)
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("fail update Description: %v", err)
// 	}

// 	return s.rowToObject(info), nil
// }

// func (s *Description) AddFields(ctx context.Context, id uuid.UUID, fields cruder.Fields) (*npool.Description, error) { //nolint
// 	var info *ent.Description
// 	var err error

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		newSold := int32(0)

// 		for k, v := range fields {
// 			increment, err := cruder.AnyTypeInt32(v)
// 			if err != nil {
// 				return fmt.Errorf("invalid value type: %v", err)
// 			}

// 			switch k {
// 			case constant.DescriptionFieldLocked:
// 				fallthrough //nolint
// 			case constant.DescriptionFieldInService:
// 				newSold += increment
// 			}
// 		}

// 		info, err = s.Tx.Description.Query().Where(Description.ID(id)).ForUpdate().Only(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail query Description: %v", err)
// 		}

// 		if int32(info.InService+info.Locked)+newSold > int32(info.Total) {
// 			return fmt.Errorf("Description exhausted")
// 		}

// 		stm := info.Update()

// 		for k, v := range fields {
// 			increment, err := cruder.AnyTypeInt32(v)
// 			if err != nil {
// 				return fmt.Errorf("invalid value type: %v", err)
// 			}

// 			switch k {
// 			case constant.DescriptionFieldLocked:
// 				stm = stm.AddLocked(increment)
// 			case constant.DescriptionFieldInService:
// 				stm = stm.AddInService(increment)
// 				if increment > 0 {
// 					stm = stm.AddSold(increment)
// 				}
// 			}
// 		}

// 		info, err = stm.Save(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail to update Description: %v", err)
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("fail add Description fields: %v", err)
// 	}

// 	return s.rowToObject(info), nil
// }

// func (s *Description) Row(ctx context.Context, id uuid.UUID) (*npool.Description, error) {
// 	var info *ent.Description
// 	var err error

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		info, err = s.Tx.Description.Query().Where(Description.ID(id)).Only(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("fail get Description: %v", err)
// 	}

// 	return s.rowToObject(info), nil
// }

// func (s *Description) queryFromConds(conds cruder.Conds) (*ent.DescriptionQuery, error) { //nolint
// 	stm := s.Tx.Description.Query()
// 	for k, v := range conds {
// 		switch k {
// 		case constant.FieldID:
// 			id, err := cruder.AnyTypeUUID(v.Val)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid id: %v", err)
// 			}
// 			stm = stm.Where(Description.ID(id))
// 		case constant.DescriptionFieldGoodID:
// 			id, err := cruder.AnyTypeUUID(v.Val)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid good id: %v", err)
// 			}
// 			stm = stm.Where(Description.GoodID(id))
// 		case constant.DescriptionFieldTotal:
// 			value, err := cruder.AnyTypeUint32(v.Val)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid total value: %v", err)
// 			}
// 			switch v.Op {
// 			case cruder.EQ:
// 				stm = stm.Where(Description.TotalEQ(value))
// 			case cruder.GT:
// 				stm = stm.Where(Description.TotalGT(value))
// 			case cruder.LT:
// 				stm = stm.Where(Description.TotalLT(value))
// 			}
// 		case constant.DescriptionFieldLocked:
// 			value, err := cruder.AnyTypeUint32(v.Val)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid value type: %v", err)
// 			}
// 			switch v.Op {
// 			case cruder.EQ:
// 				stm = stm.Where(Description.LockedEQ(value))
// 			case cruder.GT:
// 				stm = stm.Where(Description.LockedGT(value))
// 			case cruder.LT:
// 				stm = stm.Where(Description.LockedLT(value))
// 			}
// 		case constant.DescriptionFieldInService:
// 			value, err := cruder.AnyTypeUint32(v.Val)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid value type: %v", err)
// 			}
// 			switch v.Op {
// 			case cruder.EQ:
// 				stm = stm.Where(Description.InServiceEQ(value))
// 			case cruder.GT:
// 				stm = stm.Where(Description.InServiceGT(value))
// 			case cruder.LT:
// 				stm = stm.Where(Description.InServiceLT(value))
// 			}
// 		case constant.DescriptionFieldSold:
// 			value, err := cruder.AnyTypeUint32(v.Val)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid value type: %v", err)
// 			}
// 			switch v.Op {
// 			case cruder.EQ:
// 				stm = stm.Where(Description.SoldEQ(value))
// 			case cruder.GT:
// 				stm = stm.Where(Description.SoldGT(value))
// 			case cruder.LT:
// 				stm = stm.Where(Description.SoldLT(value))
// 			}
// 		default:
// 			return nil, fmt.Errorf("invalid Description field")
// 		}
// 	}

// 	return stm, nil
// }

// func (s *Description) Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.Description, int, error) {
// 	rows := []*ent.Description{}
// 	var total int

// 	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		stm, err := s.queryFromConds(conds)
// 		if err != nil {
// 			return fmt.Errorf("fail construct stm: %v", err)
// 		}

// 		total, err = stm.Count(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail count Description: %v", err)
// 		}

// 		rows, err = stm.Order(ent.Desc(Description.FieldUpdatedAt)).Offset(offset).Limit(limit).All(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail query Description: %v", err)
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, 0, fmt.Errorf("fail get Description: %v", err)
// 	}

// 	infos := []*npool.Description{}
// 	for _, row := range rows {
// 		infos = append(infos, s.rowToObject(row))
// 	}

// 	return infos, total, nil
// }

// func (s *Description) RowOnly(ctx context.Context, conds cruder.Conds) (*npool.Description, error) {
// 	var info *ent.Description

// 	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		stm, err := s.queryFromConds(conds)
// 		if err != nil {
// 			return fmt.Errorf("fail construct stm: %v", err)
// 		}

// 		info, err = stm.Only(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail query Description: %v", err)
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("fail get Description: %v", err)
// 	}

// 	return s.rowToObject(info), nil
// }

// func (s *Description) Count(ctx context.Context, conds cruder.Conds) (uint32, error) {
// 	var err error
// 	var total int

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		stm, err := s.queryFromConds(conds)
// 		if err != nil {
// 			return fmt.Errorf("fail construct stm: %v", err)
// 		}

// 		total, err = stm.Count(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail check Descriptions: %v", err)
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return 0, fmt.Errorf("fail count Descriptions: %v", err)
// 	}

// 	return uint32(total), nil
// }

// func (s *Description) Exist(ctx context.Context, id uuid.UUID) (bool, error) {
// 	var err error
// 	exist := false

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		exist, err = s.Tx.Description.Query().Where(Description.ID(id)).Exist(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return false, fmt.Errorf("fail check Description: %v", err)
// 	}

// 	return exist, nil
// }

// func (s *Description) ExistConds(ctx context.Context, conds cruder.Conds) (bool, error) {
// 	var err error
// 	exist := false

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		stm, err := s.queryFromConds(conds)
// 		if err != nil {
// 			return fmt.Errorf("fail construct stm: %v", err)
// 		}

// 		exist, err = stm.Exist(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail check Descriptions: %v", err)
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return false, fmt.Errorf("fail check Descriptions: %v", err)
// 	}

// 	return exist, nil
// }

// func (s *Description) Delete(ctx context.Context, id uuid.UUID) (*npool.Description, error) {
// 	var info *ent.Description
// 	var err error

// 	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
// 		info, err = s.Tx.Description.UpdateOneID(id).
// 			SetDeletedAt(uint32(time.Now().Unix())).
// 			Save(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("fail delete Description: %v", err)
// 	}

// 	return s.rowToObject(info), nil
// }
