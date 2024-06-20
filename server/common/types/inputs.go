package types

type LOGIN_USER_INPUT struct {
	Email           string `db:"email" json:"email"`
	AccountPassword string `db:"account_password" json:"accountPassword"`
}

type GET_USER_BY_EMAIL_INPUT struct {
	Email string `db:"email" json:"email"`
}

type UPDATE_USER_INPUT struct {
	Email         string `db:"email" json:"email"`
	FirstName     string `db:"first_name" json:"firstname"`
	LastName      string `db:"last_name" json:"lastName"`
	PhoneNumber   string `db:"phone_number" json:"phoneNumber"`
	UserAddress   string `db:"user_address" json:"userAddress"`
	IsShopEnabled bool   `db:"is_shop_enabled" json:"isShopEnabled"`
}

type CREATE_SHOP_INPUT struct {
	ShopName        string `db:"shop_name" json:"shopName"`
	Email           string `db:"email" json:"email"`
	PhoneNumber     string `db:"phone_number" json:"phoneNumber"`
	MapLocation     string `db:"map_location" json:"mapLocation"`
	ShopType        string `db:"shop_type" json:"shopType"`
	ShopDescription string `db:"shop_description" json:"shopDescription"`
	City            string `db:"city" json:"city"`
	Country         string `db:"country" json:"country"`
}

type UPDATE_SHOP_INPUT struct {
	ShopName        string `db:"shop_name" json:"shopName"`
	Email           string `db:"email" json:"email"`
	PhoneNumber     string `db:"phone_number" json:"phoneNumber"`
	MapLocation     string `db:"map_location" json:"mapLocation"`
	ShopType        string `db:"shop_type" json:"shopType"`
	ShopDescription string `db:"shop_description" json:"shopDescription"`
}

type DELETE_SHOP_INPUT struct {
	Email string `db:"email" json:"email"`
}

type GET_SHOP_BY_ID_INPUT struct {
	Id int `db:"id" json:"id"`
}

type CREATE_PRODUCT_INPUT struct {
	ProductName                  string `db:"product_name" json:"productName"`
	ProductType                  string `db:"product_type" json:"productType"`
	ProductCondition             string `db:"product_condition" json:"productCondition"`
	Price                        string `db:"price" json:"price"`
	OriginalPurchasedDate        string `db:"original_purchased_date" json:"originalPurchasedDate"`
	OriginalPurchaisingRecieptNo string `db:"original_purchaising_reciept_no" json:"originalPurchaisingRecieptNo"`
	ProductDescription           string `db:"product_description" json:"productDescription"`
}

type DELETE_PRODUCT_INPUT struct {
	Id int `db:"id" json:"id"`
}

type GET_PRODUCT_BY_ID_INPUT struct {
	Id int `db:"id" json:"id"`
}

type SEARCH_PRODUCTS_INPUT struct {
	SearchTerm string `db:"search_term" json:"searchTerm"`
}

type PLACE_ORDER_INPUT struct {
	FromMapLocation     string `db:"from_map_location" json:"fromMapLocation"`
	ToMapLocation       string `db:"to_map_location" json:"toMapLocation"`
	LastStopMapLocation string `db:"last_stop_map_location" json:"lastStopMapLocation"`
	OrderStatus         string `db:"order_status" json:"orderStatus"`
	PaymentStatus       string `db:"payment_status" json:"paymentStatus"`
	ProductId           int    `db:"product_id" json:"productId"`
	BuyerId             int    `db:"buyer_id" json:"buyerId"`
	ShopId              int    `db:"shop_id" json:"shopId"`
	Quantity            int    `db:"quantity" json:"quantity"`
}
