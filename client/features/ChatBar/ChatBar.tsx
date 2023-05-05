import { FC } from "react";

import SearchBar from "./SearchBar/SearchBar";
import ChatItem from "./ChatItem/ChatItem";
import styles from './styles.module.scss';

const ChatBar: FC = () => {
  return (
    <aside className="flex flex-col gap-4 chats-bg h-full max-w-[400px] w-full p-2 text-main">
      <SearchBar />
      <ul className={`flex flex-col gap-2 overflow-y-scroll ${styles.ChatScrollBar}`}>
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
        <ChatItem />
      </ul>
    </aside>
  )
}

export default ChatBar