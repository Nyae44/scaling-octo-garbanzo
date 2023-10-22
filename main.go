package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Payment represents data about payments made
type Payment struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	Balance int    `json:"balance"`
}

// Payments slice to seed record payment data
var payments = []Payment{
	{"1", "Collins Ruto", 1500, 500},
	{"2", "Daley Nyae", 1800, 200},
	{"3", "Timothy Arege", 1900, 100},
	{"4", "Ken Walibora", 1000, 1000},
	{"5", "Miles Davis", 1950, 50},
	{"6", "John Coltrane", 2000, 0},
}

// A handler to return all payments
func getPayments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, payments)
}

// A handler to return payments by ID
func getPaymentByID(c *gin.Context) {
	payment := &payments[0] // Get the first element of the array
	if c.Param("id") == payment.ID {
		c.IndentedJSON(http.StatusOK, payment)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No payment found with that id."})
	}
}
func main() {
	router := gin.Default()
	router.GET("/payments", getPayments)
	router.GET("/payments/:id", getPaymentByID)
	router.Run("localhost:8080")
}
