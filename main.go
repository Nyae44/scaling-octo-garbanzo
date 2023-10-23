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
	c.IndentedJSON(http.StatusOK, payments) // IndentedJson serializes the struct into JSON and adds it  to the response
}

// A handler to add a new payment from JSON received in the request body
// Th
func postPayments(c *gin.Context) {
	var newPayment Payment
	// Call BindJson to bind the received JSON to newPayment
	if err := c.BindJSON(&newPayment); err != nil {
		return
	}
	// Add the new payment to slice
	payments = append(payments, newPayment)
	c.IndentedJSON(http.StatusCreated, newPayment)
}

// A handler to return a specific payment by unique ID
// getPaymentByID locates the payment whose ID value matches the id
// parameter sent by the client, then returns that album as the response
func getPaymentById(c *gin.Context) {
	id := c.Param("id")
	// loop over the list of payments, looking for
	// a payment whose id value matches the parameter
	for _, p := range payments {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "payment not found"})
}

func main() {
	router := gin.Default()
	router.GET("/payments", getPayments)
	router.GET("/payments/:id", getPaymentById)
	router.POST("/payments", postPayments)
	router.Run("localhost:8080")
}
