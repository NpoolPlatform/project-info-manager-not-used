// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coinproductinfo"
	"github.com/google/uuid"
)

// CoinProductInfo is the model entity for the CoinProductInfo schema.
type CoinProductInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// ProductPage holds the value of the "product_page" field.
	ProductPage string `json:"product_page,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinProductInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinproductinfo.FieldCreateAt, coinproductinfo.FieldUpdateAt, coinproductinfo.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case coinproductinfo.FieldProductPage:
			values[i] = new(sql.NullString)
		case coinproductinfo.FieldID, coinproductinfo.FieldAppID, coinproductinfo.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinProductInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinProductInfo fields.
func (cpi *CoinProductInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinproductinfo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cpi.ID = *value
			}
		case coinproductinfo.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				cpi.CreateAt = uint32(value.Int64)
			}
		case coinproductinfo.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				cpi.UpdateAt = uint32(value.Int64)
			}
		case coinproductinfo.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				cpi.DeleteAt = uint32(value.Int64)
			}
		case coinproductinfo.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				cpi.AppID = *value
			}
		case coinproductinfo.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cpi.CoinTypeID = *value
			}
		case coinproductinfo.FieldProductPage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_page", values[i])
			} else if value.Valid {
				cpi.ProductPage = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinProductInfo.
// Note that you need to call CoinProductInfo.Unwrap() before calling this method if this CoinProductInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (cpi *CoinProductInfo) Update() *CoinProductInfoUpdateOne {
	return (&CoinProductInfoClient{config: cpi.config}).UpdateOne(cpi)
}

// Unwrap unwraps the CoinProductInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cpi *CoinProductInfo) Unwrap() *CoinProductInfo {
	tx, ok := cpi.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinProductInfo is not a transactional entity")
	}
	cpi.config.driver = tx.drv
	return cpi
}

// String implements the fmt.Stringer.
func (cpi *CoinProductInfo) String() string {
	var builder strings.Builder
	builder.WriteString("CoinProductInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", cpi.ID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", cpi.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", cpi.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", cpi.DeleteAt))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", cpi.AppID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cpi.CoinTypeID))
	builder.WriteString(", product_page=")
	builder.WriteString(cpi.ProductPage)
	builder.WriteByte(')')
	return builder.String()
}

// CoinProductInfos is a parsable slice of CoinProductInfo.
type CoinProductInfos []*CoinProductInfo

func (cpi CoinProductInfos) config(cfg config) {
	for _i := range cpi {
		cpi[_i].config = cfg
	}
}
