package handlers

import (
	"service"
	"github.com/labstack/echo"
	"net/http"
	"model"
)

func GetData(e echo.Context)error{
	symbol := e.Param("id")
	data ,err :=model.ReqData(symbol)
	if err != 0{
		return e.JSON(http.StatusNotFound,"symbol code가 잘못되었습니다")
	}
	senddata := model.MyInformation{}
	buy,sell,diff := service.FindDate(data)
	senddata.Buy =buy.Format("2006-01-02")
	senddata.Sell =sell.Format("2006-01-02")
	senddata.Profit = float32(diff)
	senddata.Symbol = symbol
	return e.JSON(http.StatusOK,senddata)
}

