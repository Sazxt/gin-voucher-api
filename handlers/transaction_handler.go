package handlers

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
)

type RedemptionRequest struct {
    Vouchers []struct {
        VoucherID int `json:"voucher_id"`
        Quantity  int `json:"quantity"`
    } `json:"vouchers"`
}

func MakeRedemption(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req RedemptionRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        tx, err := db.Begin()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
            return
        }

        var totalCost int
        for _, voucher := range req.Vouchers {
            var costPerUnit int
            query := "SELECT cost_in_point FROM vouchers WHERE id = $1"
            err := db.QueryRow(query, voucher.VoucherID).Scan(&costPerUnit)
            if err != nil {
                tx.Rollback()
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid voucher ID"})
                return
            }
            totalCost += costPerUnit * voucher.Quantity
        }

        var transactionID int
        query := "INSERT INTO transactions (total_cost) VALUES ($1) RETURNING id"
        err = tx.QueryRow(query, totalCost).Scan(&transactionID)
        if err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
            return
        }

        for _, voucher := range req.Vouchers {
            query = "INSERT INTO transaction_vouchers (transaction_id, voucher_id, quantity) VALUES ($1, $2, $3)"
            _, err := tx.Exec(query, transactionID, voucher.VoucherID, voucher.Quantity)
            if err != nil {
                tx.Rollback()
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link vouchers to transaction"})
                return
            }
        }

        tx.Commit()
        c.JSON(http.StatusCreated, gin.H{"transaction_id": transactionID, "total_cost": totalCost})
    }
}

func GetTransactionDetail(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        transactionID := c.Query("transactionId")

        var totalCost int
        query := "SELECT total_cost FROM transactions WHERE id = $1"
        err := db.QueryRow(query, transactionID).Scan(&totalCost)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
            return
        }

        rows, err := db.Query("SELECT voucher_id, quantity FROM transaction_vouchers WHERE transaction_id = $1", transactionID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transaction details"})
            return
        }
        defer rows.Close()

        var details []map[string]interface{}
        for rows.Next() {
            var voucherID, quantity int
            rows.Scan(&voucherID, &quantity)
            details = append(details, map[string]interface{}{
                "voucher_id": voucherID,
                "quantity":   quantity,
            })
        }

        c.JSON(http.StatusOK, gin.H{
            "transaction_id": transactionID,
            "total_cost":     totalCost,
            "details":        details,
        })
    }
}
