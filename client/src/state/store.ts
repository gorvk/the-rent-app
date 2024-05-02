import { configureStore } from "@reduxjs/toolkit";
import currentUser from "./user/userSlice"

export const store = configureStore({
    reducer: { currentUser: currentUser }
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;