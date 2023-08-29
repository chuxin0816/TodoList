package main

import (
	"github.com/chuxin0816/web-tutorial/dao"
	"github.com/chuxin0816/web-tutorial/models"
	"github.com/chuxin0816/web-tutorial/routers"
)

func main() {
	if err := dao.InitMysql(); err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":8080")
}
