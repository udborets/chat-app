import ChatBody from "@/features/ChatBody/ChatBody";
import ChatMenu from "@/features/ChatMenu/ChatMenu";
import { FC } from "react";

const ChatPage: FC = () => {
  return (
    <main
      className="w-full h-full flex"
    >
      <ChatMenu />
      <ChatBody />
    </main>
  )
}

export default ChatPage;