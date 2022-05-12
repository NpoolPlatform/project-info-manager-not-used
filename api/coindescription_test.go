package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/project-info-manager"
	testinit "github.com/NpoolPlatform/project-info-manager/pkg/test-init"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appid = uuid.NewString()

var description = npool.CoinDescriptionBase{
	AppID:      appid,
	CoinTypeID: uuid.NewString(),
	Message:    "test_message",
	Title:      "test_title",
	UsedFor:    "test_usedfor",
}

var description1 = npool.CoinDescriptionBase{
	AppID:      appid,
	CoinTypeID: uuid.NewString(),
	Message:    "test_message1",
	Title:      "test_title1",
	UsedFor:    "test_usedfor1",
}

var description2 = npool.CoinDescriptionBase{
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

func CreateOne(t *testing.T) npool.CoinDescription {
	ret := npool.CreateCoinDescriptionResponse{}
	resp := httpReq(npool.CreateCoinDescriptionRequest{
		Info: &description,
	}, "/v1/create/coin/description", t)
	err := json.Unmarshal(resp.Body(), &ret)
	assert.Nil(t, err)
	return *ret.GetInfo()
}

func CreateMany(t *testing.T, infos []*npool.CoinDescriptionBase) []*npool.CoinDescription {
	info := npool.CreateCoinDescriptionsResponse{}
	resp := httpReq(npool.CreateCoinDescriptionsRequest{
		Infos: infos,
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
		Infos: []*npool.CoinDescriptionBase{
			&description,
			&description1,
			&description2,
		},
	}, "/v1/create/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &info)
	assert.Nil(t, err)
	assert.Equal(t, len(info.GetInfos()), 3)
}

func TestUpdateCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	cInfo := CreateOne(t)
	upMsg := "yep"
	cInfo.Message = upMsg

	info := npool.UpdateCoinDescriptionResponse{}
	resp := httpReq(npool.UpdateCoinDescriptionRequest{
		AppID: cInfo.GetAppID(),
		Info:  &cInfo,
	}, "/v1/update/coin/description", t)
	err := json.Unmarshal(resp.Body(), &info)
	assert.Nil(t, err)
	assert.Equal(t, info.GetInfo().GetMessage(), upMsg)
}

func TestUpdateAppCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	info := CreateOne(t)
	upMsg := "yep"
	info.Message = upMsg

	upInfo := npool.UpdateAppCoinDescriptionResponse{}
	resp := httpReq(npool.UpdateAppCoinDescriptionRequest{
		TargetAppID: info.GetAppID(),
		Info:        &info,
	}, "/v1/update/app/coin/description", t)
	err := json.Unmarshal(resp.Body(), &upInfo)
	assert.Nil(t, err)
	assert.Equal(t, upInfo.GetInfo().GetMessage(), upMsg)
}

func TestGetCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	info := CreateOne(t)

	getInfo := npool.GetCoinDescriptionResponse{}
	resp := httpReq(npool.GetCoinDescriptionRequest{
		AppID: info.GetAppID(),
		ID:    info.GetID(),
	}, "/v1/get/coin/description", t)
	err := json.Unmarshal(resp.Body(), &getInfo)
	assert.Nil(t, err)
	assert.Equal(t, getInfo.GetInfo().GetID(), info.GetID())
}

func TestGetAppCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	info := CreateOne(t)

	getInfo := npool.GetAppCoinDescriptionResponse{}
	resp := httpReq(npool.GetAppCoinDescriptionRequest{
		TargetAppID: info.GetAppID(),
		ID:          info.GetID(),
	}, "/v1/get/app/coin/description", t)
	err := json.Unmarshal(resp.Body(), &getInfo)
	assert.Nil(t, err)
	assert.Equal(t, getInfo.GetInfo().GetID(), info.GetID())
}

func TestGetCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	appid := uuid.NewString()
	description.AppID = appid
	description1.AppID = appid
	description2.AppID = appid
	d1 := []*npool.CoinDescriptionBase{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, d1)

	getInfos := npool.GetCoinDescriptionsResponse{}
	resp := httpReq(npool.GetCoinDescriptionsRequest{
		AppID:  appid,
		Offset: 0,
		Limit:  5,
	}, "/v1/get/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &getInfos)
	assert.Nil(t, err)
	assert.Equal(t, getInfos.GetTotal(), int32(3))
}

func TestGetAppCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	appid := uuid.NewString()
	description.AppID = appid
	description1.AppID = appid
	description2.AppID = appid
	d1 := []*npool.CoinDescriptionBase{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, d1)

	getInfos := npool.GetAppCoinDescriptionsResponse{}
	resp := httpReq(npool.GetAppCoinDescriptionsRequest{
		TargetAppID: appid,
		Offset:      0,
		Limit:       5,
	}, "/v1/get/app/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &getInfos)
	assert.Nil(t, err)
	assert.Equal(t, getInfos.GetTotal(), int32(3))
}

func TestCountCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	appid := uuid.NewString()
	description.AppID = appid
	description1.AppID = appid
	description2.AppID = appid
	d1 := []*npool.CoinDescriptionBase{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, d1)

	getInfos := npool.CountCoinDescriptionsResponse{}
	resp := httpReq(npool.CountCoinDescriptionsRequest{
		AppID: appid,
	}, "/v1/count/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &getInfos)
	assert.Nil(t, err)
	assert.Equal(t, getInfos.GetResult(), uint32(3))
}

func TestCountAppCoinDescriptions(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	appid := uuid.NewString()
	description.AppID = appid
	description1.AppID = appid
	description2.AppID = appid
	d1 := []*npool.CoinDescriptionBase{
		&description,
		&description1,
		&description2,
	}
	CreateMany(t, d1)

	getInfos := npool.CountAppCoinDescriptionsResponse{}
	resp := httpReq(npool.CountAppCoinDescriptionsRequest{
		TargetAppID: appid,
	}, "/v1/count/app/coin/descriptions", t)
	err := json.Unmarshal(resp.Body(), &getInfos)
	assert.Nil(t, err)
	assert.Equal(t, getInfos.GetResult(), uint32(3))
}

func TestDeleteAppCoinDescription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	appid := uuid.NewString()
	description.AppID = appid

	cInfos := CreateMany(t, []*npool.CoinDescriptionBase{&description})

	delInfo := npool.DeleteAppCoinDescriptionResponse{}
	resp := httpReq(npool.DeleteAppCoinDescriptionRequest{
		TargetAppID: appid,
		ID:          cInfos[0].GetID(),
	}, "/v1/delete/app/coin/description", t)
	err := json.Unmarshal(resp.Body(), &delInfo)
	assert.Nil(t, err)

	ret := npool.CountAppCoinDescriptionsResponse{}
	resp = httpReq(npool.CountAppCoinDescriptionsRequest{
		TargetAppID: appid,
	}, "/v1/count/app/coin/descriptions", t)
	err = json.Unmarshal(resp.Body(), &ret)
	assert.Nil(t, err)
	assert.Equal(t, ret.GetResult(), uint32(0))
}
