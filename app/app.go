package main

import (
	"com.huahuo/app/dao"
	"com.huahuo/app/dao/dal"
	"com.huahuo/app/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) { //全局中间件设置gorm操作的数据库实例
		dal.SetDefault(dao.DB)
		c.Next()
	})
	router.InitRouter(r)
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
