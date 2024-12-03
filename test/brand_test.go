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

func TestCreateBrand(t *testing.T) {
	router := gin.Default()
	router.POST("/brand", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"id":   1,
			"name": "Test Brand",
		})
	})

	body := map[string]string{"name": "Test Brand"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/brand", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, "Test Brand", response["name"])
	assert.Equal(t, float64(1), response["id"])
}