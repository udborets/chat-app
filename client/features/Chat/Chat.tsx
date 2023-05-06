import { FC } from "react";

import MessageInput from "./MessageInput/MessageInput";
import Messages from "./Messages/Messages";
import { ChatProps } from "./models";

const Chat: FC<ChatProps> = ({ messages }) => {
  return (
    <div className="w-full h-full flex flex-col justify-end">
      <Messages messages={messages} />
      <MessageInput />
    </div>
  )
}

export default Chat;