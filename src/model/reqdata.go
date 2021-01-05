package model

import (
	"github.com/markcheno/go-quote"
	"time"
)

func ReqData(symbol string) (quote.Quote,int){

	before := time.Now().UTC().AddDate(0,0,-180).Format("2006-01-02")
	now := time.Now().UTC().Format("2006-01-02")
	var flag int
	flag = 0
	StockData, _ := quote.NewQuoteFromYahoo(symbol, before,now, quote.Daily, true)
	temp, _ := quote.NewQuoteFromCSV(symbol, StockData.CSV())
	if (temp.Date[0].Format("2006-01-02") == "0001-01-01"){
		flag =1
	}
	return temp ,flag
}
