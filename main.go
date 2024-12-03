package main

import (
    "github.com/gin-gonic/gin"
    "gin-voucher-api/database"
    "gin-voucher-api/handlers"
)

func main() {
    r := gin.Default()

    // Setup koneksi database
    db := database.Connect()

    // Routes
    r.POST("/brand", handlers.CreateBrand(db))
    r.POST("/voucher", handlers.CreateVoucher(db))
    r.GET("/voucher", handlers.GetSingleVoucher(db))
    r.GET("/voucher/brand", handlers.GetVouchersByBrand(db))
    r.POST("/transaction/redemption", handlers.MakeRedemption(db))
    r.GET("/transaction/redemption", handlers.GetTransactionDetail(db))

    // Jalankan server
    r.Run(":8080")
}
