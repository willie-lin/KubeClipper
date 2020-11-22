package handler

import (
	"KubeClipper/model"
	"github.com/gin-gonic/gin"
	"net/http"
)


// @Summary 通过 ip 获取ip 是否为CDNIP
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 {object} model.Result  成功后返回值
// @Router /cdn/{ip} [get]
func getCndIp(ctx gin.Context) {
	ip := ctx.Param("ip")

	cndip := model.Cdnip{
		Ip:   ip,
	}

	i := cndip.QueryByIp()

	ctx.JSON(http.StatusOK, gin.H{
		"result": model.Result{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    i,
		},

	})
}
