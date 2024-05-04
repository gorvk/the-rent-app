import { configureStore } from "@reduxjs/toolkit";
import userSlice from "./slices/userSlice"
import productsSlice from "./slices/productsSlice"

export const store = configureStore({
    reducer: { currentUser: userSlice.reducer, productsList: productsSlice.reducer }
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;