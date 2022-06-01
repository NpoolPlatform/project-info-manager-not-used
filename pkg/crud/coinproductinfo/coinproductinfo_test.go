package coinproductinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/projectinfomgr"
	constant "github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coinproductinfo"
	testinit "github.com/NpoolPlatform/project-info-manager/pkg/test-init"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

var description = npool.CoinProductInfo{
	AppID:       uuid.New().String(),
	CoinTypeID:  uuid.New().String(),
	ProductPage: "test_ProductPage",
}

var description1 = npool.CoinProductInfo{
	AppID:       uuid.New().String(),
	CoinTypeID:  uuid.New().String(),
	ProductPage: "test_ProductPage1",
}

var description2 = npool.CoinProductInfo{
	AppID:       uuid.New().String(),
	CoinTypeID:  uuid.New().String(),
	ProductPage: "test_ProductPage2",
}

//nolint
func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err := schema.Create(context.Background(), &description)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			description.ID = info.ID
		}
		assert.Equal(t, info, &description)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	infos, err := schema.CreateBulk(context.Background(), []*npool.CoinProductInfo{&description1, &description2})
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}

	description.ID = info.ID
	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Update(context.Background(), &description)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &description)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Row(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &description)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	infos, total, err := schema.Rows(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
		0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0], &description)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.RowOnly(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &description)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	count, err := schema.Count(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Delete(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &description)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	count, err = schema.Count(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(0))
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	_, err = schema.Row(context.Background(), uuid.MustParse(info.ID))
	assert.NotNil(t, err)
}
