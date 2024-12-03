package handlers

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Voucher struct {
    ID          int    `json:"id"`
    BrandID     int    `json:"brand_id"`
    Name        string `json:"name"`
    CostInPoint int    `json:"cost_in_point"`
}

func CreateVoucher(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var voucher Voucher
        if err := c.ShouldBindJSON(&voucher); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        query := "INSERT INTO vouchers (brand_id, name, cost_in_point) VALUES ($1, $2, $3) RETURNING id"
        err := db.QueryRow(query, voucher.BrandID, voucher.Name, voucher.CostInPoint).Scan(&voucher.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create voucher"})
            return
        }

        c.JSON(http.StatusCreated, voucher)
    }
}

func GetSingleVoucher(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Query("id")

        var voucher Voucher
        query := "SELECT id, brand_id, name, cost_in_point FROM vouchers WHERE id = $1"
        err := db.QueryRow(query, id).Scan(&voucher.ID, &voucher.BrandID, &voucher.Name, &voucher.CostInPoint)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Voucher not found"})
            return
        }

        c.JSON(http.StatusOK, voucher)
    }
}

func GetVouchersByBrand(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        brandID := c.Query("id")

        rows, err := db.Query("SELECT id, name, cost_in_point FROM vouchers WHERE brand_id = $1", brandID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vouchers"})
            return
        }
        defer rows.Close()

        var vouchers []Voucher
        for rows.Next() {
            var voucher Voucher
            rows.Scan(&voucher.ID, &voucher.Name, &voucher.CostInPoint)
            vouchers = append(vouchers, voucher)
        }

        c.JSON(http.StatusOK, vouchers)
    }
}
