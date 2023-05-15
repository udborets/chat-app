import { FC } from "react";
import ChatList from "./ChatList/ChatList";

const ChatMenu: FC = () => {
  return (
    <aside
      className="max-w-4/5 pc:max-w-[270px] w-full "
    >
      <ChatList />
    </aside>
  )
}

export default ChatMenu;