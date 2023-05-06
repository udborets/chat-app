import { FC } from "react";

import ChatItem from "./ChatItem/ChatItem";
import SearchBar from "./SearchBar/SearchBar";
import { ChatBarProps } from "./models";
import styles from './styles.module.scss';

const ChatBar: FC<ChatBarProps> = ({ chats }) => {
  return (
    <aside className="flex flex-col gap-4 chats-bg h-full max-w-[400px] w-full p-2 text-main">
      <SearchBar />
      <ul className={`flex flex-col gap-2 overflow-y-scroll ${styles.ChatScrollBar}`}>
        {chats.map(chat => (
          <ChatItem {...chat} key={chat.id} />
        ))}
      </ul>
    </aside>
  )
}

export default ChatBar