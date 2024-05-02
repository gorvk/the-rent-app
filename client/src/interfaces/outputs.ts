import { IProduct, IUser } from "./models";

export interface IGetAllProductsOutput {
    isSuccess: boolean,
    result: IProduct[]
}

export interface IGetUserOutput {
    isSuccess: boolean,
    result: IUser
}

export interface ICommonOutput {
    isSuccess: boolean,
    result: boolean
}