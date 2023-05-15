import { configureStore } from "@reduxjs/toolkit";

import { userState } from "./reducers/userState";

export const store = configureStore({
  reducer: {
    user: userState.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
