import { FC } from "react";

import SearchBar from "./SearchBar/SearchBar";
import ChatItem from "./ChatItem/ChatItem";
import styles from './styles.module.scss';

const ChatBar: FC = () => {
  return (
    <div className="flex flex-col gap-4 bg-slate-500 h-full max-w-[400px] w-full p-2">
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
    </div>
  )
}

export default ChatBar