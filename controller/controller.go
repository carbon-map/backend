package controller

import (
	"net/http"

	"github.com/labstack/echo/v4" //引入echo框架
)

// echo.Context中存放了request跟response
// Echo的Handler function 預設會回傳一個error變數，存放錯誤訊息
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!") //回傳字串形式的response跟status code給客戶端
}
