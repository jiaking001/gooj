package router

import (
	"github.com/gin-gonic/gin"
	"gooj/service"
)

func Router() {
	r := gin.Default()

	// 首页
	r.GET("/", service.Home)
	// 登录
	//r.GET("/login", func(c *gin.Context) {})
	// 注册
	//r.GET("/register", func(c *gin.Context) {})

	// 用户
	//userGroup := r.Group("/user")
	{
		//所有以/user开头

		// 排行榜
		//userGroup.GET("/home", func(c *gin.Context) {})
		// 提交记录
		//userGroup.GET("/home", func(c *gin.Context) {})
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
