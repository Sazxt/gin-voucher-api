package models

type Voucher struct {
	ID          int    `json:"id"`
	BrandID     int    `json:"brand_id"`
	Name        string `json:"name"`
	CostInPoint int    `json:"cost_in_point"`
}
