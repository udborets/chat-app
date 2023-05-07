import { observer } from "mobx-react-lite";
import { FC } from "react";

import { chatBarState } from "@/store/ChatBarState";
import ChatItem from "./ChatItem/ChatItem";
import SearchBar from "./SearchBar/SearchBar";
import { ChatBarProps } from "./models";
import styles from './styles.module.scss';

const ChatBar: FC<ChatBarProps> = observer(({ chats }) => {
  return (
    <aside className={`${chatBarState.isActive
      ? 'flex'
      : 'hidden pc:flex'} z-10 absolute pc:relative flex-col gap-4 chats-bg h-full pc:max-w-[300px] max-w-[85%] w-full p-2 text-main`}>
      <SearchBar />
      <ul className={`flex flex-col gap-2 overflow-y-scroll ${styles.chatsScrollBar}`}>
        {chats.map(chat => (
          <ChatItem {...chat} key={chat.id} />
        ))}
      </ul>
    </aside>
  )
})

export default ChatBar;