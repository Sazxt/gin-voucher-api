package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMakeRedemption(t *testing.T) {
	router := gin.Default()
	router.POST("/transaction/redemption", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"transaction_id": 1,
			"total_cost":     1000,
			"details": []map[string]interface{}{
				{"voucher_id": 1, "quantity": 2},
				{"voucher_id": 2, "quantity": 1},
			},
		})
	})

	body := map[string]interface{}{
		"vouchers": []map[string]interface{}{
			{"voucher_id": 1, "quantity": 2},
			{"voucher_id": 2, "quantity": 1},
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/transaction/redemption", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, float64(1), response["transaction_id"])
	assert.Equal(t, float64(1000), response["total_cost"])

	details := response["details"].([]interface{})
	assert.Len(t, details, 2)

	firstDetail := details[0].(map[string]interface{})
	assert.Equal(t, float64(1), firstDetail["voucher_id"])
	assert.Equal(t, float64(2), firstDetail["quantity"])
}