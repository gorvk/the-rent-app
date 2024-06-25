import { IUpdateUserInput } from "../interfaces/inputs";
import { ICommonOutput } from "../interfaces/outputs";
import { getMethodPostHeader } from "../utils/https";
import { API_URL } from "../utils/constants";

const baseUrl = API_URL + "/user/";

export const updateUserApi = async (payload: IUpdateUserInput): Promise<ICommonOutput> => {
    const url = 'update'
    const requestInit: RequestInit | undefined = getMethodPostHeader(payload);
    const response: Response = await fetch(baseUrl + url, requestInit)
    const json: ICommonOutput = await response.json()
    return json;
}