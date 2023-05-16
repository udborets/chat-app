import { useDispatch, useSelector } from "react-redux";

import { User } from "@/models/user";
import { selectUser, userActions } from "@/store/reducers/userReducer";

export const useUser = () => {
  const dispatch = useDispatch();
  const userState = useSelector(selectUser);
  const setUserState = (userParams: Partial<User>) => {
    dispatch(userActions.setState(userParams));
  };
  return {
    userState,
    setUserState,
  };
};
