package types

type User struct {
	Id              int    `db:"id" json:"id"`
	IsShopEnabled   bool   `db:"is_shop_enabled" json:"isShopEnabled"`
	FirstName       string `db:"first_name" json:"firstname"`
	Email           string `db:"email" json:"email"`
	LastName        string `db:"last_name" json:"lastName"`
	PhoneNumber     string `db:"phone_number" json:"phoneNumber"`
	UserAddress     string `db:"user_address" json:"userAddress"`
	AccountPassword string `db:"account_password" json:"accountPassword"`
}

type Shop struct {
	Id              int    `db:"id" json:"id"`
	OwnerId         int    `db:"owner_id" json:"ownerId"`
	ShopName        string `db:"shop_name" json:"shopName"`
	Email           string `db:"email" json:"email"`
	PhoneNumber     string `db:"phone_number" json:"phoneNumber"`
	MapLocation     string `db:"map_location" json:"mapLocation"`
	ShopType        string `db:"shop_type" json:"shopType"`
	ShopDescription string `db:"shop_description" json:"shopDescription"`
	City            string `db:"city" json:"city"`
	Country         string `db:"country" json:"country"`
}

type Product struct {
	Id                           int    `db:"id" json:"id"`
	ProductName                  string `db:"product_name" json:"productName"`
	ShopId                       int    `db:"shop_id" json:"shopId"`
	ProductType                  string `db:"product_type" json:"productType"`
	ProductCondition             string `db:"product_condition" json:"productCondition"`
	Price                        string `db:"price" json:"price"`
	OriginalPurchasedDate        string `db:"original_purchased_date" json:"originalPurchasedDate"`
	OriginalPurchaisingRecieptNo string `db:"original_purchaising_reciept_no" json:"originalPurchaisingRecieptNo"`
	ProductDescription           string `db:"product_description" json:"productDescription"`
	Quantity                     int    `db:"quantity" json:"quantity"`
}

type Order struct {
	Id                  int    `db:"id" json:"id"`
	FromMapLocation     string `db:"from_map_location" json:"fromMapLocation"`
	ToMapLocation       string `db:"to_map_location" json:"toMapLocation"`
	LastStopMapLocation string `db:"last_stop_map_location" json:"lastStopMapLocation"`
	OrderStatus         string `db:"order_status" json:"orderStatus"`
	PaymentStatus       string `db:"payment_status" json:"paymentStatus"`
	ProductId           int    `db:"product_id" json:"productId"`
	BuyerId             int    `db:"buyer_id" json:"buyerId"`
	Quantity            int    `db:"quantity" json:"quantity"`
}

type SearchProductView struct {
	ProductId          int    `db:"product_id" json:"productId"`
	ShopId             int    `db:"shop_id" json:"shopId"`
	ProductName        string `db:"product_name" json:"productName"`
	ProductType        string `db:"product_type" json:"productType"`
	ProductCondition   string `db:"product_condition" json:"productCondition"`
	Price              string `db:"price" json:"price"`
	ProductDescription string `db:"product_description" json:"productDescription"`
	ShopName           string `db:"shop_name" json:"shopName"`
	City               string `db:"city" json:"city"`
	Country            string `db:"country" json:"country"`
}

type OrderListView struct {
	Id          int    `db:"id" json:"id"`
	ProductId   int    `db:"product_id" json:"productId"`
	BuyerId     int    `db:"buyer_id" json:"-"`
	OrderStatus string `db:"order_status" json:"orderStatus"`
	ProductName string `db:"product_name" json:"productName"`
	Price       string `db:"price" json:"price"`
}
