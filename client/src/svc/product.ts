import { ISearchProductsInput } from "../interfaces/inputs";
import { IGetAllProductsOutput, IGetSearchedProductsOutput } from "../interfaces/outputs";
import { getMethodGetHeader, getMethodPostHeader } from "../utils/https";

const baseUrl: string = "http://localhost:9090/api/product";

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