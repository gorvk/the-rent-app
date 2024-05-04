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
    productName: string,
    shopId: number,
    productType: string,
    productCondition: string,
    price: string,
    originalPurchasedDate: string,
    originalPurchaisingRecieptNo: string,
    productDescription: string,
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
    shopId: number,
    productName: string,
    productType: string,
    productCondition: string,
    price: string,
    productDescription: string,
    shopName: string,
    city: string,
    country: string
}