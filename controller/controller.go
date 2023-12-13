package controller

import (
	"backend/model" //引入model套件
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4" //引入echo框架
)

// echo.Context中存放了request跟response
// Echo的Handler function 預設會回傳一個error變數，存放錯誤訊息
func Home(c echo.Context) error {
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
	rows, err := db.Query("SELECT * FROM carbonmap")
	if err != nil {
		log.Fatal("查詢資料失敗:", err)
	}
	defer rows.Close()

	var ans = ""
	for rows.Next() {
		var (
			year   int
			month  int
			city   string
			amount int64
		)
		if err := rows.Scan(&year, &month, &city, &amount); err != nil {
			log.Fatal("讀取資料失敗:", err)
		}
		ans += fmt.Sprintln(year, month, city, amount)
	}

	return c.String(http.StatusOK, ans) //回傳字串形式的response跟status code給客戶端
}
