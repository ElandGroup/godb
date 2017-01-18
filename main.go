package main

import (
	"godb/model"

	"github.com/labstack/echo"
)

var (
	e = echo.New()
)

func main() {
	RegisterApi()
	model.InitDB("mysql", "root:Eland123@tcp(120.27.157.9:3306)/hellodb?charset=utf8&parseTime=True&loc=UTC")
	e.Start(":1114")
}
