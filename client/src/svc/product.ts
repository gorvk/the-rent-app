import { ICreateProductInput, ISearchProductsInput } from "../interfaces/inputs";
import { ICommonOutput, IGetAllProductsOutput, IGetCurrentShopProductOutput, IGetSearchedProductsOutput } from "../interfaces/outputs";
import { getMethodGetHeader, getMethodPostHeader } from "../utils/https";
import { API_URL } from "../utils/constants";

const baseUrl = API_URL + "/product";

export const getAllProductsApi = async (): Promise<IGetAllProductsOutput> => {
    const url: string = "/get-all-products";
    const requestInit: RequestInit | undefined = getMethodGetHeader();
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: IGetAllProductsOutput = await response.json()
    return json;
}

export const getSearchedProductsApi = async (payload: ISearchProductsInput): Promise<IGetSearchedProductsOutput> => {
    const url: string = "/search";
    const requestInit: RequestInit | undefined = getMethodPostHeader(payload);
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: IGetSearchedProductsOutput = await response.json()
    return json;
}

export const createProductApi = async (payload: ICreateProductInput): Promise<ICommonOutput> => {
    const url = '/create'
    const requestInit: RequestInit | undefined = getMethodPostHeader(payload);
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: ICommonOutput = await response.json()
    return json;
}

export const GetCurrentShopProductApi = async (): Promise<IGetCurrentShopProductOutput> => {
    const url: string = "/get-current-shop-product";
    const requestInit: RequestInit | undefined = getMethodGetHeader();
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: IGetCurrentShopProductOutput = await response.json()
    return json;
}