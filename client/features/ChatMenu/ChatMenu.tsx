import { FC } from "react";
import ChatList from "./ChatList/ChatList";

const ChatMenu: FC = () => {
  return (
    <aside
      className="absolute pc:relative max-w-[90%] pc:max-w-[270px] w-full"
    >
      <ChatList />
    </aside>
  )
}

export default ChatMenu;