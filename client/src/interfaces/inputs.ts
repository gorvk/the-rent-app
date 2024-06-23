export interface IUpdateUserInput {
    isShopEnabled: boolean;
    firstname: string;
    email: string;
    lastName: string;
    phoneNumber: string;
    userAddress: string;
    accountPassword: string;
}

export interface IRegisterInput {
    email: string;
    firstName: string;
    lastName: string;
    phoneNumber: string;
    userAddress: string;
    accountPassword: string;
}

export interface ILoginInput {
    email: string;
    accountPassword: string;
}

export interface ISearchProductsInput {
    searchTerm: string;
}

export interface ICreateShopInput {
    shopName: string;
    email: string;
    phoneNumber: string;
    mapLocation: string;
    shopType: string;
    shopDescription: string;
    city: string;
    country: string;
}

export interface ICreateProductInput {
    productName: string,
    productType: string,
    productCondition: string,
    price: string,
    originalPurchasedDate: string,
    originalPurchaisingRecieptNo: string,
    productDescription: string,
    quantity: number,
}

export interface IGetProductInput {
    id: number
}

export interface IPlaceOrderInput {
    toMapLocation: string,
    productId: number,
    quantity: number
}