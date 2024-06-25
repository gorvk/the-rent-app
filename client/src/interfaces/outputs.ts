import { IOrderList, IProduct, IProductCard, IUser } from "./models";

export interface IGetAllProductsOutput {
    isSuccess: boolean,
    result: IProductCard[]
}

export interface IGetUserOutput {
    isSuccess: boolean,
    result: IUser
}

export interface ICommonOutput {
    isSuccess: boolean,
    result: boolean
}

export interface IGetSearchedProductsOutput {
    isSuccess: boolean,
    result: IProductCard[]
}

export interface IGetCurrentShopProductOutput {
    isSuccess: boolean,
    result: IProductCard[]
}

export interface IGetProductOutput {
    isSuccess: boolean,
    result: IProduct
}

export interface IGetCurrentUserOrdersOutput {
    isSuccess: boolean,
    result: IOrderList[]
}