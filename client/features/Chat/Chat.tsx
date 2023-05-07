import { FC } from "react";

import MessageInput from "./MessageInput/MessageInput";
import Messages from "./Messages/Messages";
import { ChatProps } from "./models";
import TopBar from "./TopBar/TopBar";

const Chat: FC<ChatProps> = ({ messages }) => {
  return (
    <div className="w-full h-full flex flex-col justify-end">
      <TopBar />
      <Messages messages={messages} />
      <MessageInput />
    </div>
  )
}

export default Chat;