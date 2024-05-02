import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { IUser } from "../../interfaces/models";


const initialState: IUser | null = null as IUser | null;

const setCurrentUser = (state: IUser | null, action: PayloadAction<IUser | null>) => {
    return action.payload;
}

const removeCurrentUser = (state: IUser | null) => {
    state = {} as IUser;
}

const slice = createSlice({
    name: "currentUser",
    initialState,
    reducers: { setCurrentUser, removeCurrentUser }
});

export const actions = slice.actions;
export default slice.reducer;