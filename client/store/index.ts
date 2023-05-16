import { configureStore } from "@reduxjs/toolkit";

import { userReducer } from "./reducers/userReducer";

export const store = configureStore({
  reducer: {
    user: userReducer.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
