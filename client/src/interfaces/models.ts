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
    quantity: number,
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

export interface IOrder {
    id: number,
    fromMapLocation: string,
    toMapLocation: string,
    lastStopMapLocation: string,
    orderStatus: string,
    paymentStatus: string,
    productId: number,
    buyerId: number,
    quantity: number
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

export interface IOrderList {
    id: number,
    productId: number,
    orderStatus: string,
    productName: string,
    price: string,
}