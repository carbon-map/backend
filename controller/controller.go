package controller

import (
	"backend/model" //引入model套件
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4" //引入echo框架
)

func Home(c echo.Context) error {
	return c.File("view/home.html")
}

func DataServe(c echo.Context) error {
	log := model.InitLogger()

	// 讀取 .env 檔案
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("無法載入 .env 檔案")
	}

	// 連線至資料庫
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("user"), os.Getenv("password"), os.Getenv("host"), os.Getenv("port"), os.Getenv("db")))
	if err != nil {
		log.Error("無法連線至資料庫:", err)
	}
	defer db.Close()

	// 查詢資料庫中的表格
	SQL_cmd := "SELECT amount, isPredict FROM carbonmap where year = ? and month = ? and city = '?'"
	log.Info(SQL_cmd)
	rows, err := db.Query(SQL_cmd, c.QueryParam("year"), c.QueryParam("month"), c.QueryParam("city"))
	if err != nil {
		log.Error("查詢資料失敗:", err)
	}
	defer rows.Close()

	var amount string = ""
	var isPredict string = ""
	for rows.Next() { //逐 row 讀取回傳的資料
		// var ( //定義一系列的變數，對應至回傳資料中的 column
		// 	_amount int64
		// 	_isPredict bool
		// )
		if err := rows.Scan(&amount, &isPredict); err != nil { //將讀取到的資料存入變數中
			log.Error("讀取資料失敗:", err)
		}
		if isPredict == "0" {
			isPredict = "false"
		} else if isPredict == "1" {
			isPredict = "true"
		}
	}

	// 準備回傳給客戶端的資料
	returnValue := map[string]interface{}{
		"SQL_cmd":   SQL_cmd,
		"amount":    amount,
		"isPredict": isPredict,
	}

	// 將回傳資料轉換成JSON格式，jsonData 是 []byte 型態，已打包好的 JSON 資料
	jsonData, err := json.Marshal(returnValue)
	if err != nil {
		log.Error("轉換 JSON 時發生錯誤:", err)
	}

	// 回傳字串格式的 json 資料給客戶端
	return c.String(http.StatusOK, string(jsonData))
}
