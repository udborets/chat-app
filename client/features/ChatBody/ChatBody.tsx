import { FC } from "react";

import ChatBar from "./ChatBar/ChatBar";
import MessageInput from "./MessageInput/MessageInput";
import Messages from "./Messages/Messages";

const ChatBody: FC = ({ }) => {
  return (
    <div
      className="flex flex-col "
    >
      <ChatBar />
      <Messages />
      <MessageInput />
    </div>
  )
}

export default ChatBody;