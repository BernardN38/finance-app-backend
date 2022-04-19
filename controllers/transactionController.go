package controllers

import (
	"fmt"
	"strconv"

	"github.com/bernardn38/financefirst/db"
	"github.com/bernardn38/financefirst/models"
	"github.com/bernardn38/financefirst/token"
	"github.com/gin-gonic/gin"
)

var months = [12]string{"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec"}

var daysInMonths = [12]string{
	"31",
	"28",
	"31",
	"30",
	"31",
	"30",
	"31",
	"31",
	"30",
	"31",
	"30",
	"31",
}

func GetAllTransactions(c *gin.Context) {
	// Read from DB get all transactions
	db := db.DBConn
	var transactions []models.Transactions
	db.Find(&transactions)
	c.JSON(200, transactions)
}
func GetTransaction(c *gin.Context) {
	// Read from DB get single transaction by id
	db := db.DBConn
	id := c.Param("id")
	var transaction models.Transactions
	db.Find(&transaction, id)
	c.JSON(200, transaction)
}
func GetDashboard(c *gin.Context) {
	monthlySums := make(map[string]map[string]float64)
	db := db.DBConn
	// Read from DB
	var transactions []models.Transactions
	db.Find(&transactions) // get all transactions
	for index, element := range months {
		db.Where(fmt.Sprintf("date ilike '%s-%%' ", strconv.Itoa(index+1))).Find(&transactions)
		monthlySums[element] = make(map[string]float64)
		for _, el := range transactions {
			fmt.Println(el)

			monthlySums[element]["withdrawal"] = monthlySums[element]["withdrawal"] + el.Withdrawal
			monthlySums[element]["deposit"] = monthlySums[element]["deposit"] + el.Deposit
			monthlySums[element]["balance"] = el.Balance
		}
	}
	c.JSON(200, monthlySums)
}

func GetMonthSumsByType(c *gin.Context) {
	monthlySums := make(map[string]float64)
	// Read from DB
	db := db.DBConn
	queryType := c.Query("type")

	var transactions []models.Transactions

	switch queryType {
	case "balance":
		for index, element := range months {
			db.Select("date", "description", queryType).Where(fmt.Sprintf("date ilike '%s-%s%%' ", strconv.Itoa(index+1), daysInMonths[index])).Find(&transactions)
			for _, el := range transactions {
				monthlySums[element] = el.Balance
			}

		}
		c.JSON(200, monthlySums)
		return
	case "deposit":
		for index, element := range months {
			db.Select("date", "description", queryType).Where(fmt.Sprintf("date ilike '%s-%%' and description='pay_day'", strconv.Itoa(index+1))).Find(&transactions)
			for _, el := range transactions {
				monthlySums[element] = monthlySums[element] + el.Deposit
			}

		}
		c.JSON(200, monthlySums)
		return
	case "withdrawal":
		for index, element := range months {
			db.Select("id", "date", "description", queryType).Where(fmt.Sprintf("date ilike '%s-%%' ", strconv.Itoa(index+1))).Find(&transactions)
			for _, el := range transactions {
				monthlySums[element] = monthlySums[element] + el.Withdrawal
			}

		}
		c.JSON(200, monthlySums)
		return
	case "retirement":
		for index, element := range months {
			db.Select("date", "description", "withdrawal").Where(fmt.Sprintf("date ilike '%s-%%' and description = 'retirement_contribution' ", strconv.Itoa(index+1))).Find(&transactions)
			for _, el := range transactions {
				monthlySums[element] = monthlySums[element] + el.Withdrawal
			}
		}
		c.JSON(200, monthlySums)
		return
	case "investments":
		for index, element := range months {
			db.Select("date", "description", "investment_total").Where(fmt.Sprintf("date ilike '%s-%s-%%' ", strconv.Itoa(index+1), daysInMonths[index])).Last(&transactions)
			for _, el := range transactions {
				monthlySums[element] = monthlySums[element] + el.InvestmentTotal
			}
		}
		c.JSON(200, monthlySums)
		return
	case "net_worth":
		for index, element := range months {
			db.Select("date", "description", "balance", "investment_total").Where(fmt.Sprintf("date ilike '%s-%s-%%' ", strconv.Itoa(index+1), daysInMonths[index])).Last(&transactions)

			for _, el := range transactions {
				monthlySums[element] = el.Balance + el.InvestmentTotal
			}
			db.Select("date", "description", "withdrawal").Where("description = 'retirement_contribution'").Where(fmt.Sprintf("date ilike '%s-%%' ", strconv.Itoa(index+1))).Find(&transactions)
			for _, el := range transactions {
				monthlySums[element] = monthlySums[element] + el.Withdrawal
			}

		}
		c.JSON(200, monthlySums)
		return
	}
	c.JSON(200, monthlySums)
	return
}

func GetTransactionsLimit(c *gin.Context) {
	db := db.DBConn
	payload := c.MustGet(AuthorizationPayloadKey)
	user := payload.(*token.Payload)

	queryType := c.Query("type")
	queryLimit := c.Query("limit")

	userId := c.Param("userId")
	if fmt.Sprintf("%v", user.UserId) != userId {
		c.AbortWithStatusJSON(401, "mismatched user")
	}

	// Read from DB
	var transactions []models.Transactions
	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		panic("no limit value found")
	}
	switch queryType {
	case "retirement":
		db.Select("id", "date", "description", "withdrawal").Where("description = 'retirement_contribution' and user_id = ? ", user.UserId).Limit(limit).Order("id desc").Find(&transactions)
		c.JSON(200, transactions)
		return
	case "investments":
		db.Select("id", "date", "description", "withdrawal", "deposit", "investment_total").Where("description ilike 'investment%' and user_id = ? ", user.UserId).Limit(limit).Order("id desc").Find(&transactions)
		c.JSON(200, transactions)
		return
	case "net_worth":
		db.Select("id", "date", "description", "balance").Where("balance > 0 and user_id = ? ", user.UserId).Limit(limit).Order("id desc").Find(&transactions)
		c.JSON(200, transactions)
		return
	case "withdrawal":
		db.Select("id", "date", "description", "withdrawal").Where("withdrawal > 0 and user_id = ? ", user.UserId).Limit(limit).Order("id desc").Find(&transactions)
		c.JSON(200, transactions)
		return
	case "deposit":
		db.Select("id", "date", "description", "deposit").Where(" deposit > 0 and user_id = ? ", user.UserId).Limit(limit).Order("id desc").Find(&transactions)
		c.JSON(200, transactions)
		return
	default:
		db.Select("*").Limit(limit).Order("id desc").Find(&transactions)
		c.JSON(200, transactions)
		return
	}

}
