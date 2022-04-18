package controllers

import (
	"fmt"
	"time"

	"github.com/bernardn38/financefirst/db"
	"github.com/bernardn38/financefirst/models"
	"github.com/bernardn38/financefirst/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserCash(db *gorm.DB) string {
	var transaction models.Transactions
	db.Last(&transaction)

	cash := fmt.Sprint(transaction.Balance)
	return cash
}

func getMonthlyChange(db *gorm.DB) string {
	var transaction1 models.Transactions
	var transaction2 models.Transactions
	db.Select("date", "description", "balance").Where(fmt.Sprintf("date ilike '11-%%'")).Last(&transaction1)
	db.Select("date", "description", "balance").Where(fmt.Sprintf("date ilike '12-%%'")).Last(&transaction2)
	x, y, z := time.Now().Date()
	fmt.Println(x, "--", y, "--", z)
	monthlyChange := fmt.Sprintf("%.2f", (transaction2.Balance/transaction1.Balance*100 - 100))
	return monthlyChange
}

func getNetWorth(db *gorm.DB) string {
	var transaction models.Transactions
	db.Select("date", "description", "balance", "investment_total").Last(&transaction)
	return fmt.Sprintf("%.2f", (transaction.Balance + transaction.InvestmentTotal))
}
func GetUserSneakPeek(c *gin.Context) {
	db := db.DBConn
	cash := getUserCash(db)
	monthlyChange := getMonthlyChange(db)
	netWorth := getNetWorth(db)

	response := make(map[string]string)
	response["cash"] = cash
	response["monthly_change"] = monthlyChange
	response["net_worth"] = netWorth
	c.JSON(200, response)
}

func GetUserProfile(c *gin.Context) {
	var user models.User
	db := db.DBConn
	payload := c.MustGet(AuthorizationPayloadKey)
	data := payload.(*token.Payload)
	fmt.Println(data.Username)
	db.Where("username = ?", data.Username).First(&user)
	fmt.Println(user)
	c.JSON(200, user)
}
