export interface IUpdateUserInput {
    isShopEnabled: boolean,
    firstname: string,
    email: string,
    lastName: string,
    phoneNumber: string,
    userAddress: string,
    accountPassword: string,
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
    accountPassword: string
}