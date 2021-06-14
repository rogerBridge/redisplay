package controllers2

import (
	"encoding/json"
	"go-seckill/internal/easyjsonprocess"
	"go-seckill/internal/mysql"
	"go-seckill/internal/mysql/shop_orm"
	"go-seckill/internal/redisconf"
	"go-seckill/internal/utils"

	"github.com/valyala/fasthttp"
)

// 一次创建一个PurchaseLimit实例, 并将其添加到purchase_limits table中
func CreatePurchaseLimit(ctx *fasthttp.RequestCtx) {
	p := new(shop_orm.PurchaseLimit)
	err := json.Unmarshal(ctx.Request.Body(), p)
	if err != nil {
		logger.Warnf("Unmarshal PurchaseLimit error happen: %v", err)
		utils.ResponseWithJson(ctx, 400, easyjsonprocess.CommonResponse{
			Code: 8400,
			Msg:  "解析PurchaseLimit时出现错误",
			Data: nil,
		})
		return
	}
	logger.Infof("解析后的PurchaseLimit是: %+v", p)

	// 首先查看, PurchaseLimit的product_id是否存在于purchase_limits表格中
	if p.IfPurchaseLimitExist() {
		logger.Warnf("PurchaseLimit已有相同ID的在表格中")
		utils.ResponseWithJson(ctx, 404, easyjsonprocess.CommonResponse{
			Code: 8404,
			Msg:  "欲添加的PurchaseLimit已经存在于数据库中",
			Data: nil,
		})
		return
	}
	tx := mysql.Conn2.Begin()
	err = p.CreatePurchaseLimit(tx)
	if err != nil {
		logger.Warnf("当添加PurchaseLimit时, 错误: %v", err)
		tx.Rollback()
		utils.ResponseWithJson(ctx, 500, easyjsonprocess.CommonResponse{
			Code: 8500,
			Msg:  "当添加PurchaseLimit时, 出错",
			Data: nil,
		})
		return
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Infof("CreatePurchaseLimit tx commit error: %v", err)
		utils.ResponseWithJson(ctx, 200, easyjsonprocess.CommonResponse{
			Code: 8200,
			Msg:  "添加PurchaseLimit成功",
			Data: nil,
		})
		return
	}
	logger.Infof("添加PurchaseLimit成功")
	utils.ResponseWithJson(ctx, 200, easyjsonprocess.CommonResponse{
		Code: 8200,
		Msg:  "添加PurchaseLimit成功",
		Data: nil,
	})
}

// 根据product_id获取商品的purchase_limit
func QueryPurchaseLimit(ctx *fasthttp.RequestCtx) {
	// query purchaseLimit
	p := new(shop_orm.PurchaseLimit)
	err := json.Unmarshal(ctx.Request.Body(), p)
	if err != nil {
		logger.Warnf("While unmarshal request.body(), error: %v", err)
		utils.ResponseWithJson(ctx, 400, easyjsonprocess.CommonResponse{
			Code: 8400,
			Msg:  "While unmarshal request.body(), error",
			Data: nil,
		})
		return
	}
	logger.Infof("unmarshal request.body() success")

	// 首先查看, request.body()是否存在于purchase_limits table中
	if !p.IfPurchaseLimitExist() {
		logger.Warnf("query PurchaseLimit is not exist in purchase_limits table")
		utils.ResponseWithJson(ctx, 404, easyjsonprocess.CommonResponse{
			Code: 8404,
			Msg:  "query PurchaseLimit is not exist in purchase_limits table",
			Data: nil,
		})
		return
	}

	// 如果存在的话, 返回Purchase_limit这个对象
	var purchaseLimit *shop_orm.PurchaseLimit
	purchaseLimit, err = p.QueryPurchaseLimit()
	if err != nil {
		logger.Warnf("purchaseLimit query error: %v", err)
		utils.ResponseWithJson(ctx, 404, easyjsonprocess.CommonResponse{
			Code: 8404,
			Msg:  "query purchaseLimit error",
			Data: nil,
		})
		return
	}
	logger.Infof("purchaseLimit query succuesful")
	utils.ResponseWithJson(ctx, 200, easyjsonprocess.CommonResponse{
		Code: 8200,
		Msg:  "query PurchaseLimit successful",
		Data: purchaseLimit,
	})
}

func UpdatePurchaseLimit(ctx *fasthttp.RequestCtx) {
	p := new(shop_orm.PurchaseLimit)
	err := json.Unmarshal(ctx.Request.Body(), p)
	if err != nil {
		logger.Warnf("While unmarshal request.body(), error: %v", err)
		utils.ResponseWithJson(ctx, 400, easyjsonprocess.CommonResponse{
			Code: 8400,
			Msg:  "While unmarshal request.body(), error",
			Data: nil,
		})
		return
	}
	logger.Infof("unmarshal request.body() success")

	// 首先查看, request.body()是否存在于purchase_limits table中
	if !p.IfPurchaseLimitExist() {
		logger.Warnf("query PurchaseLimit is not exist in purchase_limits table")
		utils.ResponseWithJson(ctx, 404, easyjsonprocess.CommonResponse{
			Code: 8404,
			Msg:  "query PurchaseLimit is not exist in purchase_limits table",
			Data: nil,
		})
		return
	}

	// 如果存在的话, 返回Purchase_limit这个对象
	tx := mysql.Conn2.Begin()
	err = p.UpdatePurchaseLimit(tx)
	if err != nil {
		logger.Warnf("UpdatePurchaseLimit transaction error: %v", err)
		tx.Rollback()
		utils.ResponseWithJson(ctx, 500, easyjsonprocess.CommonResponse{
			Code: 8500,
			Msg:  "UpdatePurchaseLimit error",
			Data: nil,
		})
		return
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Warnf("UpdatePurchaseLimit transaction commit error: %v", err)
		utils.ResponseWithJson(ctx, 500, easyjsonprocess.CommonResponse{
			Code: 8500,
			Msg:  "UpdatePurchaseLimit transaction commit error",
			Data: nil,
		})
		return
	}
	logger.Infof("UpdatePurchaseLimit transaction commit successful")
	utils.ResponseWithJson(ctx, 200, easyjsonprocess.CommonResponse{
		Code: 8200,
		Msg:  "UpdatePurchaseLimit transaction commit successful",
		Data: nil,
	})
}

func DeletePurchaseLimit(ctx *fasthttp.RequestCtx) {
	// delete purchaseLimit
	p := new(shop_orm.PurchaseLimit)
	err := json.Unmarshal(ctx.Request.Body(), p)
	if err != nil {
		logger.Warnf("While unmarshal request.body(), error: %v", err)
		utils.ResponseWithJson(ctx, 400, easyjsonprocess.CommonResponse{
			Code: 8400,
			Msg:  "While unmarshal request.body(), error",
			Data: nil,
		})
		return
	}
	logger.Infof("DeletePurchaseLimit success")

	// 首先查看, request.body()是否存在于purchase_limits table中
	if !p.IfPurchaseLimitExist() {
		logger.Warnf("query PurchaseLimit is not exist in purchase_limits table")
		utils.ResponseWithJson(ctx, 404, easyjsonprocess.CommonResponse{
			Code: 8404,
			Msg:  "query PurchaseLimit is not exist in purchase_limits table",
			Data: nil,
		})
		return
	}

	// 如果存在的话, 返回Purchase_limit这个对象
	tx := mysql.Conn2.Begin()
	err = p.DeletePurchaseLimit(tx)
	if err != nil {
		logger.Warnf("DeletePurchaseLimit transaction error: %v", err)
		tx.Rollback()
		utils.ResponseWithJson(ctx, 500, easyjsonprocess.CommonResponse{
			Code: 8500,
			Msg:  "DeletePurchaseLimit error",
			Data: nil,
		})
		return
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Warnf("DeletePurchaseLimit transaction commit error: %v", err)
		utils.ResponseWithJson(ctx, 500, easyjsonprocess.CommonResponse{
			Code: 8500,
			Msg:  "DeletePurchaseLimit transaction commit error",
			Data: nil,
		})
		return
	}
	logger.Infof("DeletePurchaseLimit transaction commit successful")
	utils.ResponseWithJson(ctx, 200, easyjsonprocess.CommonResponse{
		Code: 8200,
		Msg:  "DeletePurchaseLimit transaction commit successful",
		Data: nil,
	})
}

// SyncGoodsLimit ...
// 更新商品限制计划
// 例如, 在更新MySQL的限制购买条件后, 若要将商品购买限制同步到app中, 只需要调用goodsLimit这个接口就可以
func LoadGoodPurchaseLimit(ctx *fasthttp.RequestCtx) {
	// 加载limit限制计划
	err := redisconf.LoadLimits()
	if err != nil {
		logger.Warnf("SyncGoodsLimit: 加载limit变量到全局变量purchaseLimit时出现错误 %v", err)
		utils.ResponseWithJson(ctx, 500, easyjsonprocess.CommonResponse{
			Code: 8500,
			Msg:  "加载mysql中限制购买的数据到全局变量purchaseLimit时出现错误",
			Data: nil,
		})
		return
		//ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	logger.Infof("SyncGoodsLimit: 加载limit变量到全局变量purchaseLimit成功")
	utils.ResponseWithJson(ctx, 500, easyjsonprocess.CommonResponse{
		Code: 8001,
		Msg:  "加载mysql中限制购买的数据到全局变量purchaseLimit",
		Data: nil,
	})
}
