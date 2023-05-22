import { FC } from "react";

import ChatList from "./ChatList/ChatList";
import SearchInput from "./SearchInput/SearchInput";

const ChatMenu: FC = () => {
  return (
    <aside
      className="absolute bg-white border-r-[1px] h-full z-[100] pc:relative max-w-[90%] pc:max-w-[270px] w-full"
    >
      <SearchInput />
      <ChatList />
    </aside>
  )
}

export default ChatMenu;