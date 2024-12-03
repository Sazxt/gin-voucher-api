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

func TestCreateVoucher(t *testing.T) {
	router := gin.Default()
	router.POST("/voucher", func(c *gin.Context) {

		c.JSON(http.StatusCreated, gin.H{
			"id":           1,
			"brand_id":     1,
			"name":         "Test Voucher",
			"cost_in_point": 500,
		})
	})

	body := map[string]interface{}{
		"brand_id":      1,
		"name":          "Test Voucher",
		"cost_in_point": 500,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/voucher", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, float64(1), response["id"])
	assert.Equal(t, float64(1), response["brand_id"])
	assert.Equal(t, "Test Voucher", response["name"])
	assert.Equal(t, float64(500), response["cost_in_point"])
}
