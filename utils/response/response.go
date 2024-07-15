package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NegativeResponse sends a negative response with a 400 status code
func NegativeResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": message,
	})
}

// PositiveResponse sends a positive response with a 200 status code
// It accepts an additional options map to include extra fields in the response
// In Go, the empty interface interface{} can hold any type. This makes it very 
// flexible for cases where you don't know the type of data at compile time or when 
// you want to handle multiple types generically
// A map allows you to store key-value pairs, where the keys are strings (string) 
// and the values can be of any type (interface{}). This is useful when you want to 
// pass a collection of data where each key (string) corresponds to a specific piece 
//of information, and the value can vary in type.
func PositiveResponse(c *gin.Context, message string, options ...map[string]interface{}) {
	response := gin.H{
		"success": true,
		"message": message,
	}

	// Add additional fields from the option map to the response

	
	if len(options)>0 {
		for key, value := range options[0] {
		response[key] = value
	   }
	}
	c.JSON(http.StatusOK, response)
}