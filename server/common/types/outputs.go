package types

type RESPONSE_PARAMETERS struct {
	IsSuccess bool `json:"isSuccess"`
	Result    any  `json:"result"`
}

type GET_SEARCHED_PRODUCTS_OUTPUT struct {
	SearchProductView
	Document        string `db:"document" json:"-"`
	SsQuery         string `db:"ss_query" json:"-"`
	RankProductName string `db:"rank_product_name" json:"-"`
	Similarity      string `db:"similarity" json:"-"`
}

type GET_PRODUCT_BY_ID_OUTPUT struct {
	Id                           int    `db:"products_id" json:"productId"`
	ProductName                  string `db:"product_name" json:"productName"`
	ProductType                  string `db:"product_type" json:"productType"`
	ProductCondition             string `db:"product_condition" json:"productCondition"`
	Price                        string `db:"price" json:"price"`
	OriginalPurchasedDate        string `db:"original_purchased_date" json:"originalPurchasedDate"`
	OriginalPurchaisingRecieptNo string `db:"original_purchaising_reciept_no" json:"originalPurchaisingRecieptNo"`
	ProductDescription           string `db:"product_description" json:"productDescription"`
	Quantity                     int    `db:"quantity" json:"quantity"`
	ShopId                       int    `db:"shop_id" json:"shopId"`
	ShopName                     string `db:"shop_name" json:"shopName"`
	City                         string `db:"city" json:"city"`
	Country                      string `db:"country" json:"country"`
	Email                        string `db:"email" json:"email"`
	PhoneNumber                  string `db:"phone_number" json:"phoneNumber"`
	MapLocation                  string `db:"map_location" json:"mapLocation"`
	ShopType                     string `db:"shop_type" json:"shopType"`
	ShopDescription              string `db:"shop_description" json:"shopDescription"`
}
