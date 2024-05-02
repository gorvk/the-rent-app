import { IGetAllProductsOutput } from "../interfaces/outputs";
import { getMethodGetHeader } from "../utils/https";

const baseUrl: string = "http://localhost:9090/api/";

export const getAllProductsApi = async (): Promise<IGetAllProductsOutput> => {
    const url: string = "product/get-all-products";
    const requestInit: RequestInit | undefined = getMethodGetHeader();
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: IGetAllProductsOutput = await response.json()
    return json;
}