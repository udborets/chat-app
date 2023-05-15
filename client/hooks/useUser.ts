import { useDispatch, useSelector } from "react-redux";

import { User } from "@/models/user";
import { selectUser, userActions } from "@/store/reducers/userState";

export const useUser = () => {
  const dispatch = useDispatch();
  const user = useSelector(selectUser);
  const setState = (userParams: Partial<User>) => {
    dispatch(userActions.setState(userParams));
  };
  return {
    user,
    setState,
  };
};
