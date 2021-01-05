package test

import (
	"api/handlers"
	"fmt"
)

func test(){
	testList := []string{"AAPL",
		"GOOG",
		"APPLE",
		"SAMSUNG",
		"amd",
		"intc",
		"MSFT",
		"nflx",
		"orly",
	}
	for i :=0;i < len(testList);i++{
		fmt.Println(handlers.getData(testList[i]))
	}
}
