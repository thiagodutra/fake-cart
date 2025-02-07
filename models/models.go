package models

type Cart struct {
	ID                  string  `json:"id" bson:"_id,omitempty"`
	CustomerID          string  `json:"customer_id" bson:"customer_id"`
	TransactionID       string  `json:"transaction_id" bson:"transaction_id"`
	CheckoutType        string  `json:"checkout_type" bson:"checkout_type"`
	Submitted           bool    `json:"submitted" bson:"submitted"`
	PaymentSplitPercent float64 `json:"payment_split_percent" bson:"payment_split_percent"`
	Channel             string  `json:"channel" bson:"channel"`
	Item                Item    `json:"item" bson:"item"`
}

type Item struct {
	ID        string `json:"id" bson:"id"`
	Quantity  int    `json:"quantity" bson:"quantity"`
	SkuID     string `json:"sku_id" bson:"sku_id"`
	ProductID string `json:"product_id" bson:"product_id"`
	CatalogID string `json:"catalog_id" bson:"catalog_id"`
	Price     Price  `json:"price" bson:"price"`
}

type Price struct {
	SelectedPriceList string  `json:"selected_price_list" bson:"selected_price_list"`
	PointsSalePrice   int32   `json:"points_sale_price" bson:"points_sale_price"`
	CashSalePrice     float64 `json:"cash_sale_price" bson:"cash_sale_price"`
}
