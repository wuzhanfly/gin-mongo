package handler

import (
	"fmt"
	"gin-mongo-backend/logic"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetByAddress(c *gin.Context)  {	// 1. 获取参数（从URL中获取power）
	pidStr := c.Param("address")
	res,err := logic.MinerByPeerId(pidStr)
	if err != nil {
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c,res)
}

func GetPowerInByAddress(c *gin.Context)  {
	pidStr := c.Param("address")
	res,err := logic.GetPowerInByAddress(pidStr)
	if err != nil {
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c,res)

}

func GetInfoByAddress(c *gin.Context)  {
	pidStr := c.Param("address")
	res,err := logic.GetInfoByAddress(pidStr)
	if err != nil {
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c,res)
}
//MiningStats
func MiningStats(c *gin.Context)  {
	pidStr := c.Param("address")
	type_status := c.Param("day")
	res,err := logic.MiningStats(pidStr,type_status)
	if err != nil {
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c,res)
}

func BaseInfo(c *gin.Context)  {
	res,err :=  logic.BaseInfoFun()
	if err != nil {
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c,res)


}

func GetPowerIn(c *gin.Context)  {
	res,err :=  logic.GetPowerIn()
	if err != nil {
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c,res)

}

func GetGas(c *gin.Context)  {
	types := c.Param("type")
	res,err := logic.GetGas(types)
	if err != nil {
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c,res)
}

func Test(c *gin.Context)  {
	fmt.Println(4324)
	ResponseSuccess(c,"test")

}


// GetHTML 获取数据
func GetHTML(c *gin.Context) {

	Url := c.PostForm("url")
	Url1 := strings.Split(Url, "//")[1]
	UrlOrigin := "https://" + strings.Split(Url1, "/")[0]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", Url, nil)
	for k,v :=range c.Request.Header {
		//fmt.Printf(k)
		if k == "User-Agent" {
			req.Header.Set(k, v[0])
		}
		if k == "Origin" {
			//	使用origin v[0]值来限制请求来源
			fmt.Printf(v[0])
		}
	}
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Origin", UrlOrigin)
	res, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   string(body),
	})
}

