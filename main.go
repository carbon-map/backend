package main

import (
	"backend/model"   //引入model套件
	"backend/routing" //引入routing套件

	"github.com/labstack/echo/v4" //引入echo套件
)

func main() {
	log := model.InitLogger()
	e := echo.New()               //建立一個Echo的物件
	e.Static("/static", "static") //註冊靜態檔案路徑

	routing.Routing(e) //將建立的Echo物件傳入routing() function

	log.Fatal(e.Start(":8000"))
}
