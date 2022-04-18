package models

type Transactions struct {
	// gorm.Model
	Id              int     `json:"id"`
	Date            string  `json:"date"`
	Balance         float64 `json:"balance"`
	Deposit         float64 `json:"deposit"`
	Withdrawal      float64 `json:"withdrawal"`
	Description     string  `json:"description"`
	InvestmentTotal float64 `json:"investment_total"`
}
