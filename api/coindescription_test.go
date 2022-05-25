//nolint
package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/projectinfomgr"
	constant "github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coindescription"
	testinit "github.com/NpoolPlatform/project-info-manager/pkg/test-init"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

var appid = uuid.NewString()

var description = npool.CoinDescription{
	AppID:      appid,
	CoinTypeID: uuid.NewString(),
	Message:    "test_message",
	Title:      "test_title",
	UsedFor:    "test_usedfor",
}

var description1 = npool.CoinDescription{
	AppID:      appid,
	CoinTypeID: uuid.NewString(),
	Message:    "test_message1",
	Title:      "test_title1",
	UsedFor:    "test_usedfor1",
}

var description2 = npool.CoinDescription{
	AppID:      appid,
	CoinTypeID: uuid.NewString(),
	Message:    "test_message2",
	Title:      "test_title2",
	UsedFor:    "test_usedfor2",
}

func httpReq(in interface{}, url string, t *testing.T) *resty.Response {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(in).
		Post(fmt.Sprintf("http://localhost:50310%v", url))
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		return resp
	}
	return nil
}

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func CreateOne(t *testing.T) npool.CoinDescription {
	ret := npool.CreateCoinDescriptionResponse{}
	resp := httpReq(npool.CreateCoinDescriptionRequest{
		Info: &description,
	}, "/v1/create/coin/description", t)
	err := json.Unmarshal(resp.Body(), &ret)
	assert.Nil(t, err)
	return *ret.GetInfo()
}

func CreateMany(t *testing.T, appid string, infos []*npool.CoinDescription) []*npool.CoinDescription {
	info := npool.CreateCoinDescriptionsResponse{}
	resp := httpReq(npool.CreateCoinDescriptionsRequest{
		Infos: infos,
		AppID: appid,
	}, "/v1/create/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &info)
	assert.Nil(t, err)
	return info.GetInfos()
}

func TestCreateCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	ret := npool.CreateCoinDescriptionResponse{}
	resp := httpReq(npool.CreateCoinDescriptionRequest{
		Info: &description,
	}, "/v1/create/coin/description", t)
	err := json.Unmarshal(resp.Body(), &ret)
	assert.Nil(t, err)
	assert.Equal(t, ret.GetInfo().GetCoinTypeID(), description.GetCoinTypeID())
}

func TestCreateCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	info := npool.CreateCoinDescriptionsResponse{}
	resp := httpReq(npool.CreateCoinDescriptionsRequest{
		AppID: uuid.NewString(),
		Infos: []*npool.CoinDescription{
			&description,
			&description1,
			&description2,
		},
	}, "/v1/create/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &info)
	assert.Nil(t, err)
	assert.Equal(t, len(info.GetInfos()), 3)
}

func TestCreateAppCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	ret := npool.CreateAppCoinDescriptionResponse{}
	resp := httpReq(npool.CreateAppCoinDescriptionRequest{
		Info:        &description,
		TargetAppID: description.AppID,
	}, "/v1/create/app/coin/description", t)
	err := json.Unmarshal(resp.Body(), &ret)
	assert.Nil(t, err)
	assert.Equal(t, ret.GetInfo().GetCoinTypeID(), description.GetCoinTypeID())
}

func TestCreateAppCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	info := npool.CreateAppCoinDescriptionsResponse{}
	resp := httpReq(npool.CreateAppCoinDescriptionsRequest{
		TargetAppID: uuid.NewString(),
		Infos: []*npool.CoinDescription{
			&description,
			&description1,
			&description2,
		},
	}, "/v1/create/app/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &info)
	assert.Nil(t, err)
	assert.Equal(t, len(info.GetInfos()), 3)
}

func TestUpdateCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	cInfo := CreateOne(t)
	upMsg := uuid.NewString()
	cInfo.Message = upMsg

	info := npool.UpdateCoinDescriptionResponse{}
	resp := httpReq(npool.UpdateCoinDescriptionRequest{
		Info: &cInfo,
	}, "/v1/update/coin/description", t)
	err := json.Unmarshal(resp.Body(), &info)
	assert.Nil(t, err)
	assert.Equal(t, info.GetInfo().GetMessage(), upMsg)
}

func TestGetCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	info := CreateOne(t)

	getInfo := npool.GetCoinDescriptionResponse{}
	resp := httpReq(npool.GetCoinDescriptionRequest{
		ID: info.GetID(),
	}, "/v1/get/coin/description", t)
	err := json.Unmarshal(resp.Body(), &getInfo)
	assert.Nil(t, err)
	assert.Equal(t, getInfo.GetInfo().GetID(), info.GetID())
}

func TestGetCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	rmsg := uuid.NewString()
	description.Message = rmsg
	description1.Message = rmsg
	description2.Message = rmsg
	d1 := []*npool.CoinDescription{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, appid, d1)

	getInfos := npool.GetCoinDescriptionsResponse{}
	conds := cruder.NewFilterConds()
	conds.WithCond(constant.FieldMessage, cruder.EQ, structpb.NewStringValue(rmsg))

	resp := httpReq(npool.GetCoinDescriptionsRequest{
		AppID:  description.AppID,
		Conds:  conds,
		Offset: 0,
		Limit:  5,
	}, "/v1/get/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &getInfos)
	assert.Nil(t, err)
	assert.Equal(t, getInfos.GetTotal(), int32(3))
}

func TestGetCoinDescriptionOnly(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	rmsg := uuid.NewString()
	description.Message = rmsg
	description1.Message = rmsg
	rmsg1 := uuid.NewString()
	description2.Message = rmsg1
	d1 := []*npool.CoinDescription{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, appid, d1)
	conds := cruder.NewFilterConds()
	conds.WithCond(constant.FieldMessage, cruder.EQ, structpb.NewStringValue(rmsg1))
	getInfo := npool.GetCoinDescriptionOnlyResponse{}
	resp := httpReq(npool.GetCoinDescriptionOnlyRequest{
		AppID: description.AppID,
		Conds: conds,
	}, "/v1/get/coin/description/only", t)
	err := json.Unmarshal(resp.Body(), &getInfo)
	assert.Nil(t, err)
	assert.Equal(t, getInfo.GetInfo().GetMessage(), description2.Message)
}

func TestGetAppCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	rmsg := uuid.NewString()
	description.Message = rmsg
	description1.Message = rmsg
	description2.Message = rmsg
	d1 := []*npool.CoinDescription{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, appid, d1)

	getInfos := npool.GetAppCoinDescriptionsResponse{}
	conds := cruder.NewFilterConds()
	conds.WithCond(constant.FieldMessage, cruder.EQ, structpb.NewStringValue(rmsg))

	resp := httpReq(npool.GetAppCoinDescriptionsRequest{
		TargetAppID: description.AppID,
		Conds:       conds,
		Offset:      0,
		Limit:       5,
	}, "/v1/get/app/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &getInfos)
	assert.Nil(t, err)
	assert.Equal(t, getInfos.GetTotal(), int32(3))
}

func TestGetAppCoinDescriptionOnly(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	rmsg := uuid.NewString()
	description.Message = rmsg
	description1.Message = rmsg
	rmsg1 := uuid.NewString()
	description2.Message = rmsg1
	d1 := []*npool.CoinDescription{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, appid, d1)
	conds := cruder.NewFilterConds()
	conds.WithCond(constant.FieldMessage, cruder.EQ, structpb.NewStringValue(rmsg1))
	getInfo := npool.GetCoinDescriptionOnlyResponse{}
	resp := httpReq(npool.GetAppCoinDescriptionOnlyRequest{
		TargetAppID: description.AppID,
		Conds:       conds,
	}, "/v1/get/app/coin/description/only", t)
	err := json.Unmarshal(resp.Body(), &getInfo)
	assert.Nil(t, err)
	assert.Equal(t, getInfo.GetInfo().GetMessage(), description2.Message)
}

func TestDeleteAppCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	appid := uuid.NewString()
	description.AppID = appid

	cInfos := CreateMany(t, appid, []*npool.CoinDescription{&description})

	delInfo := npool.DeleteCoinDescriptionResponse{}
	resp := httpReq(npool.DeleteCoinDescriptionRequest{
		ID: cInfos[0].GetID(),
	}, "/v1/delete/coin/description", t)
	err := json.Unmarshal(resp.Body(), &delInfo)
	assert.Nil(t, err)

	conds := cruder.NewFilterConds()
	conds.WithCond(constant.FieldAppID, cruder.EQ, structpb.NewStringValue(appid))

	getInfos := npool.GetCoinDescriptionsResponse{}
	resp = httpReq(npool.GetCoinDescriptionsRequest{
		AppID:  description.AppID,
		Conds:  conds,
		Offset: 0,
		Limit:  5,
	}, "/v1/get/coin/descriptions", t)
	err = json.Unmarshal(resp.Body(), &getInfos)
	assert.Nil(t, err)
	assert.Equal(t, getInfos.GetTotal(), int32(0))
}
