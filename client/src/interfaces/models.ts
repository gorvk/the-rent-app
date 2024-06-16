export interface IUser {
    id: number,
    isShopEnabled: boolean,
    firstname: string,
    email: string,
    lastName: string,
    phoneNumber: string,
    userAddress: string,
    accountPassword: string,
}

export interface IProduct {
    productId: number,
    shopId: number,
    productName: string,
    productType: string,
    productCondition: string,
    productDescription: string,
    price: string,
    shopName: string,
    city: string,
    country: string,
    originalPurchasedDate: string,
    originalPurchaisingRecieptNo: string,
}

export interface IShop {
    ownerId: number,
    shopName: string,
    email: string,
    phoneNumber: string,
    mapLocation: string,
    shopType: string,
    shopDescription: string
}

export interface IProductCard {
    productId: number,
    productName: string,
    productType: string,
    productCondition: string,
    price: string,
    productDescription: string,
    shopName: string,
    city: string,
    country: string
}