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

	var rows *sql.Rows
	var SQL_cmd string
	if c.QueryParam("month") == "all" {
		SQL_cmd := "SELECT amount, isPredict FROM carbonmap where year = ? and city = ?"
		log.Info(SQL_cmd)
		rows, err := db.Query(SQL_cmd, c.QueryParam("year"), c.QueryParam("city"))
		if err != nil {
			log.Error("查詢資料失敗:", err)
		}
		defer rows.Close()
	} else {
		// 查詢資料庫中的表格
		SQL_cmd := "SELECT amount, isPredict FROM carbonmap where year = ? and month = ? and city = ?"
		log.Info(SQL_cmd)
		rows, err := db.Query(SQL_cmd, c.QueryParam("year"), c.QueryParam("month"), c.QueryParam("city"))
		if err != nil {
			log.Error("查詢資料失敗:", err)
		}
		defer rows.Close()
	}

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

func DataRange(c echo.Context) error {

	log := model.InitLogger()

	// log.Error("DataRange")
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
	SQL_cmd := "SELECT * FROM carbonmap where year >= ? and year <= ? and month >= ? and month <= ? and city = ? order by year desc, month desc"
	log.Info(SQL_cmd)
	rows, err := db.Query(SQL_cmd, c.QueryParam("yearStart"), c.QueryParam("yearEnd"), c.QueryParam("monthStart"), c.QueryParam("monthEnd"), c.QueryParam("city"))
	if err != nil {
		log.Error("查詢資料失敗:", err)
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			log.Fatal("取得列名失敗:", err)
		}

		// 準備一個 slice 來保存查詢結果
		values := make([]interface{}, len(columns))
		for i := range columns {
			values[i] = new(interface{})
		}

		// 使用 Scan 函數讀取每一行的值
		if err := rows.Scan(values...); err != nil {
			log.Fatal("讀取行失敗:", err)
		}

		// 將每一行的結果映射為 map[string]interface{}
		rowData := make(map[string]interface{})
		for i, colName := range columns {
			val := *(values[i].(*interface{}))
			rowData[colName] = val
		}

		// 將該行的 map 加入結果 slice
		results = append(results, rowData)
	}

	// 回傳 JSON
	return c.JSON(200, results)

}
