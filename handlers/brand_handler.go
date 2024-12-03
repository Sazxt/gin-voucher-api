package handlers

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Brand struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func CreateBrand(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var brand Brand
        if err := c.ShouldBindJSON(&brand); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        query := "INSERT INTO brands (name) VALUES ($1) RETURNING id"
        err := db.QueryRow(query, brand.Name).Scan(&brand.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brand"})
            return
        }

        c.JSON(http.StatusCreated, brand)
    }
}
