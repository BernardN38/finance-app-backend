package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bernardn38/financefirst/db"
	"github.com/bernardn38/financefirst/models"
	"github.com/bernardn38/financefirst/routes"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routes.SetupRouter()
	db.InitDb()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetAllTransactions(t *testing.T) {
	router := routes.SetupRouter()
	db.InitDb()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/transactions", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	response := make([]models.Transactions, 1, 1000)
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(response), 1)
}

func TestGetSignleTransaction(t *testing.T) {
	router := routes.SetupRouter()
	db.InitDb()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/transactions/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	response := models.Transactions{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	// assert.Equal(t, len(response), 1)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.Description)
}

func TestGetMonthlySums(t *testing.T) {
	router := routes.SetupRouter()
	db.InitDb()

	types := []string{"deposit", "investments", "withdrawal", "retirement", "balance"}
	for _, v := range types {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/transactions/sum?type=%s", v), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		response := make(map[string]int)
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.Nil(t, err)
		fmt.Println(response)
		assert.NotEmpty(t, response)
		assert.Equal(t, len(response), 12)
	}

}
