import { observer } from "mobx-react-lite";
import { FC } from "react";

import { chatBarState } from "@/store/ChatBarState";
import ChatItem from "./ChatItem/ChatItem";
import SearchBar from "./SearchBar/SearchBar";
import { ChatBarProps } from "./models";

const ChatBar: FC<ChatBarProps> = observer(({ chats }) => {
  return (
    <>
      <aside className={`${chatBarState.isActive
        ? 'translate-x-[0%]'
        : 'translate-x-[-100%]'} flex z-10 pc:translate-x-0 absolute pc:relative flex-col gap-4 
      bg-secondary h-full pc:max-w-[300px] max-w-[85%] w-full p-2 text-main transition-all duration-500`}>
        <SearchBar />
        <ul className={`flex flex-col gap-2 overflow-y-scroll pc:scrollBar`}>
          {chats.map(chat => (
            <ChatItem {...chat} key={chat.id} />
          ))}
        </ul>
      </aside>
      <div
        className={`${chatBarState.isActive
          ? 'z-[9] bg-[#000000af]'
          : 'z-[-1] bg-none'} h-full w-full duration-300 transition-all  absolute`}
        onClick={() => chatBarState.toggleIsActive()}
      />
    </>
  )
})

export default ChatBar;