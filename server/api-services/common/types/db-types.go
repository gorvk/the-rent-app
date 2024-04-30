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
}
