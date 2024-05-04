import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { IProductCard } from "../../interfaces/models";

const initialState: IProductCard[] = [];

const setProductsList = (_: IProductCard[], action: PayloadAction<IProductCard[]>) => {
    return action.payload;
}

const slice = createSlice({
    name: "productsList",
    initialState,
    reducers: { setProductsList }
})

const reducerAction = { actions: slice.actions, reducer: slice.reducer }
export default reducerAction;