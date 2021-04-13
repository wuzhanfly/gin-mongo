package api

import (
	"gin-mongo-backend/app/handler"
	"github.com/gin-gonic/gin"
)

func ToolsAPI(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	v1.POST("/getHTML", handler.GetHTML)
	v1.GET("/getByAddress/:address",handler.GetByAddress)
	v1.GET("/getPower",handler.GetPowerIn)
	v1.GET("/getPowerIn/:address",handler.GetPowerInByAddress)
	v1.GET("/getInfoByAddress/:address",handler.GetInfoByAddress)
	v1.GET("/MiningStats/:address/:day",handler.MiningStats)
	v1.GET("/baseInfo",handler.BaseInfo)
	v1.GET("/test",handler.Test)
	v1.GET("/gatGas/:type",handler.GetGas)

}