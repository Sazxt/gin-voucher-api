package models

import "time"

type Transaction struct {
	ID        int                  `json:"id"`
	TotalCost int                  `json:"total_cost"`
	CreatedAt time.Time            `json:"created_at"`
	Details   []TransactionVoucher `json:"details,omitempty"`
}

type TransactionVoucher struct {
	ID            int `json:"id"`
	TransactionID int `json:"transaction_id"`
	VoucherID     int `json:"voucher_id"`
	Quantity      int `json:"quantity"`
}
