import { IPlaceOrderInput } from "../interfaces/inputs";
import { ICommonOutput, IGetCurrentUserOrdersOutput } from "../interfaces/outputs";
import { getMethodGetHeader, getMethodPostHeader } from "../utils/https";
import { API_URL } from "../utils/constants";

const baseUrl = API_URL + "/order/";

export const createOrderApi = async (payload: IPlaceOrderInput): Promise<ICommonOutput> => {
    const url = 'create'
    const requestInit: RequestInit | undefined = getMethodPostHeader(payload);
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: ICommonOutput = await response.json()
    return json;
}

export const getCurrentUserOrdersApi = async (): Promise<IGetCurrentUserOrdersOutput> => {
    const url = 'get-current-user-orders'
    const requestInit: RequestInit | undefined = getMethodGetHeader();
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: IGetCurrentUserOrdersOutput = await response.json()
    return json;
}