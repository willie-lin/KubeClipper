package handler

import (
	"KubeClipper/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册新用户

func regester(ctx *gin.Context)  {

	// 获取用户
	user := &model.User{}

	result :=&model.Result{
		Code:	200,
		Message:	"注册成功！",
		Data:	nil,
	}

	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败！"
		result.Code = http.StatusUnauthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}
	u := user.QueryUsername()
	if u.Passwoed == user.Passwoed {

	}


}
