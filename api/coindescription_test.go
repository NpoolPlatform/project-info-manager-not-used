package api

import (
	"encoding/json"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/project-info-manager"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

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

var s = Server{}

func httpReq(in interface{}, url string, t *testing.T, t_func func(*resty.Response)) {
	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(in).
		Post("http://localhost:50310" + url)
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		t_func(resp)

	}
}

func CreateOne(t *testing.T) npool.CoinDescription {
	in := npool.CreateCoinDescriptionRequest{
		Info: &description,
	}
	ret := npool.CreateCoinDescriptionResponse{}
	httpReq(in, "/v1/create/coin/description", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &ret)
		assert.Nil(t, err)
	})
	return *ret.GetInfo()
}

func CreateMany(t *testing.T, infos []*npool.CoinDescriptionBase) []*npool.CoinDescription {
	in := npool.CreateCoinDescriptionsRequest{
		Infos: infos,
	}
	info := npool.CreateCoinDescriptionsResponse{}
	httpReq(in, "/v1/create/coin/descriptions", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &info)
		assert.Nil(t, err)
	})
	return info.GetInfos()
}

func TestCreateCoinDescription(t *testing.T) {
	in := npool.CreateCoinDescriptionRequest{
		Info: &description,
	}
	ret := npool.CreateCoinDescriptionResponse{}
	httpReq(in, "/v1/create/coin/description", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &ret)
		assert.Nil(t, err)
	})
	assert.Equal(t, ret.GetInfo().GetCoinTypeID(), in.Info.GetCoinTypeID())
}

func TestCreateCoinDescriptions(t *testing.T) {
	in := npool.CreateCoinDescriptionsRequest{
		Infos: []*npool.CoinDescriptionBase{
			&description,
			&description1,
			&description2,
		},
	}
	info := npool.CreateCoinDescriptionsResponse{}
	httpReq(in, "/v1/create/coin/descriptions", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &info)
		assert.Nil(t, err)
	})
	assert.Equal(t, len(info.GetInfos()), 3)
}

func TestUpdateCoinDescription(t *testing.T) {
	cInfo := CreateOne(t)
	upMsg := "yep"
	cInfo.Message = upMsg
	in := npool.UpdateCoinDescriptionRequest{
		AppID: cInfo.GetAppID(),
		Info:  &cInfo,
	}
	info := npool.UpdateCoinDescriptionResponse{}
	httpReq(in, "/v1/update/coin/description", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &info)
		assert.Nil(t, err)
	})
	assert.Equal(t, info.GetInfo().GetMessage(), upMsg)
}

func TestUpdateAppCoinDescription(t *testing.T) {
	info := CreateOne(t)
	upMsg := "yep"
	info.Message = upMsg
	in := npool.UpdateAppCoinDescriptionRequest{
		TargetAppID: info.GetAppID(),
		Info:        &info,
	}
	upInfo := npool.UpdateAppCoinDescriptionResponse{}
	httpReq(in, "/v1/update/app/coin/description", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &upInfo)
		assert.Nil(t, err)
	})
	assert.Equal(t, upInfo.GetInfo().GetMessage(), upMsg)
}

func TestGetCoinDescription(t *testing.T) {
	info := CreateOne(t)

	in := npool.GetCoinDescriptionRequest{
		AppID: info.GetAppID(),
		ID:    info.GetID(),
	}

	getInfo := npool.GetCoinDescriptionResponse{}
	httpReq(in, "/v1/get/coin/description", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &getInfo)
		assert.Nil(t, err)
	})
	assert.Equal(t, getInfo.GetInfo().GetID(), info.GetID())
}

func TestGetAppCoinDescription(t *testing.T) {
	info := CreateOne(t)

	in := npool.GetAppCoinDescriptionRequest{
		TargetAppID: info.GetAppID(),
		ID:          info.GetID(),
	}

	getInfo := npool.GetAppCoinDescriptionResponse{}
	httpReq(in, "/v1/get/app/coin/description", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &getInfo)
		assert.Nil(t, err)
	})
	assert.Equal(t, getInfo.GetInfo().GetID(), info.GetID())
}

func TestGetCoinDescriptions(t *testing.T) {
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
	in := npool.GetCoinDescriptionsRequest{
		AppID:  appid,
		Offset: 0,
		Limit:  5,
	}

	getInfos := npool.GetCoinDescriptionsResponse{}
	httpReq(in, "/v1/get/coin/descriptions", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &getInfos)
		assert.Nil(t, err)
	})
	assert.Equal(t, getInfos.GetTotal(), int32(3))
}

func TestGetAppCoinDescriptions(t *testing.T) {
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
	in := npool.GetAppCoinDescriptionsRequest{
		TargetAppID: appid,
		Offset:      0,
		Limit:       5,
	}

	getInfos := npool.GetAppCoinDescriptionsResponse{}
	httpReq(in, "/v1/get/app/coin/descriptions", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &getInfos)
		assert.Nil(t, err)
	})
	assert.Equal(t, getInfos.GetTotal(), int32(3))
}

func TestCountCoinDescriptions(t *testing.T) {
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

	in := npool.CountCoinDescriptionsRequest{
		AppID: appid,
	}

	getInfos := npool.CountCoinDescriptionsResponse{}
	httpReq(in, "/v1/count/coin/descriptions", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &getInfos)
		assert.Nil(t, err)
	})
	assert.Equal(t, getInfos.GetResult(), uint32(3))
}

func TestCountAppCoinDescriptions(t *testing.T) {
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

	in := npool.CountAppCoinDescriptionsRequest{
		TargetAppID: appid,
	}

	getInfos := npool.CountAppCoinDescriptionsResponse{}
	httpReq(in, "/v1/count/app/coin/descriptions", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &getInfos)
		assert.Nil(t, err)
	})
	assert.Equal(t, getInfos.GetResult(), uint32(3))
}

func TestDeleteAppCoinDescription(t *testing.T) {
	appid := uuid.NewString()
	description.AppID = appid

	cInfos := CreateMany(t, []*npool.CoinDescriptionBase{&description})
	in := npool.DeleteAppCoinDescriptionRequest{
		TargetAppID: appid,
		ID:          cInfos[0].GetID(),
	}

	delInfo := npool.DeleteAppCoinDescriptionResponse{}
	httpReq(in, "/v1/delete/app/coin/description", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &delInfo)
		assert.Nil(t, err)
	})

	cIn := npool.CountAppCoinDescriptionsRequest{
		TargetAppID: appid,
	}

	ret := npool.CountAppCoinDescriptionsResponse{}
	httpReq(cIn, "/v1/count/app/coin/descriptions", t, func(resp *resty.Response) {
		err := json.Unmarshal(resp.Body(), &ret)
		assert.Nil(t, err)
	})
	assert.Equal(t, ret.GetResult(), uint32(0))
}
