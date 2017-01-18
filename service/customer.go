package service

import (
	"godb/model"
	"strconv"

	"net/http"

	"github.com/labstack/echo"
)

func Create(c echo.Context) error {

	var customer model.Customer
	id, _ := strconv.ParseInt(c.FormValue("Id"), 10, 64)
	customer.Id = id
	customer.Mobile = c.FormValue("Mobile")
	err := customer.Create()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: map[string]string{"status": "OK"}})
}

func Retrive(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cust, err := model.Customer{}.Get(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	result := make(map[string]interface{})
	if cust != nil {
		result["Mobile"] = cust.Mobile
		result["Id"] = cust.Id
	}
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
}
