import { FC } from "react";

import { ChatListProps } from "./models";
import ChatListItem from "./ChatListItem/ChatListItem";

const ChatList: FC<ChatListProps> = ({ }) => {
  return (
    <ul
      className="absolute pc:relative flex flex-col w-full"
    >
      <ChatListItem />
      <ChatListItem />
      <ChatListItem />
      <ChatListItem />
    </ul>
  )
}

export default ChatList