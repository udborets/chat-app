import { FC } from "react";

import Message from "./Message/Message";
import { MessagesProps } from "./models";

const Messages: FC<MessagesProps> = ({ messages }) => {
  return (
    <div className="w-full h-full flex flex-col overflow-y">
      {messages.map((messageProps) => (
        <Message {...messageProps} key={messageProps.id} />
      ))}
    </div>
  )
}

export default Messages;