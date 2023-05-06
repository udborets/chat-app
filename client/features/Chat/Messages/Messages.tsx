import { FC } from "react";

import Message from "./Message/Message";
import { MessagesProps } from "./models";

const Messages: FC<MessagesProps> = ({ messages }) => {
  return (
    <div className="w-full h-full flex text-main flex-col overflow-y-scroll gap-1 py-5">
      {messages.map((messageProps) => (
        <Message {...messageProps} key={messageProps.id} />
      ))}
    </div>
  )
}

export default Messages;