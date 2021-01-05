package model

type MyInformation struct {
	Buy			string 	`json:"buy_date"`
	Sell		string	`json:"sell_date"`
	Profit		float32	`json:"profit"`
	Symbol		string	`json:"symbol"`
}
