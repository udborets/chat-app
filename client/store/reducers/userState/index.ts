import { PayloadAction, createSlice } from "@reduxjs/toolkit";

import { User } from "@/models/user";
import { RootState } from "@/store";
import { UserState } from "./models";

const initialState: UserState = {
  value: {
    avatar_url: "",
    created_at: 0,
    email: "",
    hash_password: "",
    id: 0,
    last_seen: 0,
    name: "",
    phone: "",
    updated_at: 0,
  },
};

export const userState = createSlice({
  initialState,
  name: "user",
  reducers: {
    setState(state, action: PayloadAction<Partial<User>>) {
      state.value = { ...state.value, ...action.payload };
    },
  },
});

export const userActions = userState.actions;
export const selectUser = (state: RootState) => state.user;
