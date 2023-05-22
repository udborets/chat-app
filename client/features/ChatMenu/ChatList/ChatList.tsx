import { FC } from "react";

import { ChatListProps } from "./models";
import ChatListItem from "./ChatListItem/ChatListItem";
import styles from './styles.module.scss';

const ChatList: FC<ChatListProps> = ({ }) => {
  return (
    <ul
      className={`flex flex-col w-full ${styles.chatListItemBorder}`}
    >
      <ChatListItem />
      <ChatListItem />
      <ChatListItem />
      <ChatListItem />
    </ul>
  )
}

export default ChatList