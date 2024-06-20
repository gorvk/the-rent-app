import { IPlaceOrderInput } from "../interfaces/inputs";
import { ICommonOutput } from "../interfaces/outputs";
import { getMethodPostHeader } from "../utils/https";
import { API_URL } from "../utils/constants";

const baseUrl = API_URL + "/order";

export const createOrderApi = async (payload: IPlaceOrderInput): Promise<ICommonOutput> => {
    const url = '/create'
    const requestInit: RequestInit | undefined = getMethodPostHeader(payload);
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: ICommonOutput = await response.json()
    return json;
}