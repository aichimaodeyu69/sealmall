package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	serviceAccount "server/service/account/account"
	serviceOrder "server/service/account/order"
	"server/service/server"
	"server/source/server/logic/orm/dal"
)

func main() {
	g := gin.New()

	const dsn = "root:123456@tcp(localhost:3306)/sealmall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	dal.SetDefault(db)
	acc, err := dal.Account.Where(dal.Account.ID.Eq("a")).First()
	if err != nil {
		return
	}
	fmt.Println(acc)
	//登录 注册 忘记密码 注销
	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout", server.Logout)

	account := g.Group("/account")
	order := account.Group("/order")
	order.POST("/list", serviceOrder.List)

	account.POST("Profile", serviceAccount.Profile)

	g.Run(":8080")
}
